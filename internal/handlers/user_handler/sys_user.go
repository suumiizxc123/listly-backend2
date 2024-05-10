package user_handler

import (
	"kcloudb1/internal/config"
	"kcloudb1/internal/models/user"
	"kcloudb1/internal/utils"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func CreateSysUser(c *gin.Context) {
	var sysUser user.SysUser

	if err := c.ShouldBindJSON(&sysUser); err != nil {
		c.JSON(
			http.StatusBadRequest,
			utils.Error([]string{"System user fields required", "Системийн хэрэглэгчийн мэдээлэл дутуу байна"}, err.Error()),
		)

		return
	}

	sysUser.UID = uuid.New().String()

	if ok := sysUser.CheckEmailAndPhoneNotExist(); !ok {

		c.JSON(
			http.StatusInternalServerError,
			utils.Error([]string{"System user creation failed", "Системийн хэрэглэгчийг хадгалахад боломжгүй байна. Имэйл эсвэл утасны дугаар давхардсан байна"}, "email or phone already exist"),
		)

		return
	}

	if err := sysUser.Create(); err != nil {
		c.JSON(
			http.StatusInternalServerError,
			utils.Error([]string{"System user creation failed", "Системийн хэрэглэгчийг хадгалахад алдаа гарлаа"}, err.Error()),
		)

		return
	}

	c.JSON(200, utils.Success([]string{"System user created", "Системийн хэрэглэгчийг хадгаллаа"}, struct{}{}))

}

func GetSysUser(c *gin.Context) {

	var err error
	var sysUser user.SysUser
	id, ok := c.GetQuery("id")

	if !ok {
		c.JSON(
			http.StatusBadRequest,
			utils.Error([]string{"System user id required", "Системийн хэрэглэгчийн id дутуу байна"}, "id must be required"),
		)
		return
	}

	sysUser.ID, err = strconv.ParseInt(id, 10, 64)

	if err != nil {
		c.JSON(
			http.StatusBadRequest,
			utils.Error([]string{"System user id cannot be parsed", "Системийн хэрэглэгчийн id давхардсан байна"}, err.Error()),
		)
		return
	}

	if err := sysUser.Get(); err != nil {
		c.JSON(
			http.StatusInternalServerError,
			utils.Error([]string{"System user not found", "Системийн хэрэглэгчийг олдсонгүй"}, err.Error()),
		)
		return
	}

	if sysUser.ID == 0 {
		c.JSON(
			http.StatusNoContent,
			utils.Success([]string{"System user not found", "Системийн хэрэглэгчийг олдсонгүй"}, nil),
		)
		return
	}

	sysUser.Password = ""

	c.JSON(200, utils.Success([]string{"User found", "Хэрэглэгч олдсон"}, sysUser))
}

func GetSysUserList(c *gin.Context) {

	// var err error
	var sysUser user.SysUser

	sysUsers, err := sysUser.GetList()

	if err != nil {
		c.JSON(
			http.StatusInternalServerError,
			utils.Error([]string{"System user list not found", "Системийн хэрэглэгчийн жагсаалт олдсонгүй"}, err.Error()),
		)
		return
	}

	var i int
	for i = 0; i < len(sysUsers); i++ {
		sysUsers[i].Password = ""
	}

	c.JSON(200, utils.Success([]string{"System user list", "Системийн хэрэглэгчийн жагсаалт"}, sysUsers))

}

func LoginSysUser(c *gin.Context) {

	var input user.SysUserLoginInput
	var sysUser user.SysUser

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(
			http.StatusBadRequest,
			utils.Error([]string{"System user login fields required", "Системийн хэрэглэгчийн мэдээлэл дутуу байна"}, err.Error()),
		)
		return
	}

	sysUser, err := sysUser.Login(input.Phone, input.Password)

	// add redis token generate
	if err != nil {

		c.JSON(
			http.StatusInternalServerError,
			utils.Error([]string{"System user login failed", "Хэрэглэгч нэвтрэхэд алдаа гарлаа"}, err.Error()),
		)
		return
	}

	sysUser.Password = ""

	if sysUser.IsActive == 0 {

		c.JSON(
			http.StatusForbidden,
			utils.Error([]string{"System user login failed", "Системийн хэрэглэгчийн мэдээлэл давхардсан байна"}, "inactive user"),
		)
		return
	}

	token := uuid.New().String()

	jsonSysUser, err := sysUser.MarshalJSON()
	if err != nil {

		c.JSON(
			http.StatusInternalServerError,
			utils.Error([]string{"System user login failed in marshal json", "Системийн хэрэглэгч нэвтрэх json задлахад алдаа гарлаа"}, err.Error()),
		)
		return
	}

	if err := config.RS.Set(token, string(jsonSysUser), 12*time.Hour).Err(); err != nil {

		c.JSON(
			http.StatusInternalServerError,
			utils.Error([]string{"System user login failed in redis set", "Системийн хэрэглэгч нэвтрэх redis холбоход алдаа гарлаа"}, err.Error()),
		)

		return
	}

	sysUser.Token = token

	c.JSON(200, utils.Success([]string{"System user login success", "Системийн хэрэглэгч нэвтрэх амжилттай боллоо"}, sysUser))

}

func UpdateSysUser(c *gin.Context) {
	var sysUser user.SysUser

	if err := c.ShouldBindJSON(&sysUser); err != nil {
		c.JSON(
			http.StatusBadRequest,
			utils.Error([]string{"System user update fields required", "Системийн хэрэглэгчийн мэдээлэл дутуу байна"}, err.Error()),
		)
		return
	}

	if err := sysUser.Update(); err != nil {

		c.JSON(
			http.StatusInternalServerError,
			utils.Error([]string{"System user update failed", "Системийн хэрэглэгчийн шинэчилэхэд алдаа гарлаа байна"}, err.Error()),
		)
		return
	}

	c.JSON(200, utils.Success([]string{"System user update success", "Системийн хэрэглэгчийн шинэчилэх амжилттай боллоо"}, sysUser))

}

func UpdateSysUserPassword(c *gin.Context) {
	var input struct {
		Phone       string `json:"phone"`
		Password    string `json:"password"`
		NewPassword string `json:"new_password"`
	}
	var sysUser user.SysUser

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(
			http.StatusBadRequest,
			utils.Error([]string{"System user update fields required", "Системийн хэрэглэгчийн мэдээлэл дутуу байна"}, err.Error()),
		)
		return
	}

	if err := config.DB.Where("phone = ?", input.Phone).First(&sysUser).Error; err != nil {

		c.JSON(
			http.StatusInternalServerError,
			utils.Error([]string{"System user not found", "Системийн хэрэглэгчийг олдсонгүй"}, err.Error()),
		)
		return
	}

	if input.Password != sysUser.Password {
		c.JSON(
			http.StatusForbidden,
			utils.Error([]string{"System user password did not match", "Системийн хэрэглэгчийн нууц үг буруу боллоо"}, "incorrect password"),
		)
		return
	}

	sysUser.Password = input.NewPassword

	if err := sysUser.Update(); err != nil {

		c.JSON(
			http.StatusInternalServerError,
			utils.Error([]string{"System user update failed", "Системийн хэрэглэгчийн шинэчилэхэд алдаа гарлаа байна"}, err.Error()),
		)
		return
	}

	c.JSON(200, utils.Success([]string{"System user update password success", "Системийн хэрэглэгчийн шинэчилэх амжилттай боллоо"}, struct{}{}))
}

func DeleteSysUser(c *gin.Context) {
	var err error
	var sysUser user.SysUser

	id, ok := c.GetQuery("id")

	if !ok {
		c.JSON(
			http.StatusBadRequest,
			utils.Error([]string{"System user id required", "Системийн хэрэглэгчийн id дутуу байна"}, "id must be required"),
		)
		return
	}

	sysUser.ID, err = strconv.ParseInt(id, 10, 64)

	if err != nil {
		c.JSON(
			http.StatusBadRequest,
			utils.Error([]string{"System user id cannot be parsed", "Системийн хэрэглэгчийн id давхардсан байна"}, err.Error()),
		)
		return
	}
	if err := sysUser.Delete(); err != nil {

		c.JSON(
			http.StatusInternalServerError,
			utils.Error([]string{"System user delete failed", "Системийн хэрэглэгчийг устгахад алдаа гарлаа"}, err.Error()),
		)
		return
	}

	c.JSON(200, utils.Success([]string{"System user delete success", "Системийн хэрэглэгчийг амжилттай устгалаа"}, struct{}{}))
}
