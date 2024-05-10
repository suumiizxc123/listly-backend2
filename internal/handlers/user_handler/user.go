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

func CreateUser(c *gin.Context) {
	var user user.User

	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(
			http.StatusBadRequest,
			utils.Error([]string{"User fields required", "Хэрэглэгчийн мэдээлэл дутуу байна"}, err.Error()),
		)
		return
	}

	if ok := user.CheckEmailAndPhoneNotExist(); !ok {

		c.JSON(
			http.StatusInternalServerError,
			utils.Error([]string{"User creation failed", "Хэрэглэгчийн мэдээлэл үүссэн хэрэглэгчтэй давхардаж байна. Имэйл эсвэл утасны дугаар давхардсан байна"}, "email or phone already exist"),
		)
		return
	}

	user.UID = uuid.New().String()

	if err := user.Create(); err != nil {
		c.JSON(
			http.StatusInternalServerError,
			utils.Error([]string{"User creation failed", "Хэрэглэгчийг хадгалахад алдаа гарлаа"}, err.Error()),
		)
		return
	}

	user.Password = ""

	c.JSON(200, utils.Success([]string{"User creation successful", "Хэрэглэгчийг хадгаллаа"}, user))
}

func LoginUser(c *gin.Context) {
	var input user.UserLoginInput
	var user user.User

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(
			http.StatusBadRequest,
			utils.Error([]string{"User fields required", "Хэрэглэгчийн мэдээлэл дутуу байна"}, err.Error()),
		)
		return
	}

	user, err := user.Login(input.Phone, input.Password)

	if err != nil {
		c.JSON(
			http.StatusInternalServerError,
			utils.Error([]string{"Login failed", "Хэрэглэгч нэвтрэхэд алдаа гарлаа"}, err.Error()),
		)
		return
	}

	user.Password = ""

	if user.IsActive == 0 {

		c.JSON(
			http.StatusForbidden,
			utils.Error([]string{"Login failed", "Хэрэглэгч идэвхгүй байна"}, "user not active"),
		)
		return
	}

	token := uuid.New().String()

	jsonUser, err := user.MarshalJSON()

	if err != nil {

		c.JSON(
			http.StatusInternalServerError,
			utils.Error([]string{"Login failed", "Хэрэглэгчийн мэдээлэл ашиглахад алдаа гарлаа. "}, err.Error()),
		)
		return
	}

	if err := config.RS.Set(token, jsonUser, 12*time.Hour).Err(); err != nil {

		c.JSON(
			http.StatusInternalServerError,
			utils.Error([]string{"Login failed", "Хэрэглэгчийн мэдээллийг санахад алдаа гарлаа. RS"}, err.Error()),
		)
		return
	}

	user.Token = token

	c.JSON(200, utils.Success([]string{"Login successful", "Амжилттай нэвтэрлээ"}, user))
}

func GetUser(c *gin.Context) {
	var err error
	var user user.User
	id, ok := c.GetQuery("id")

	if !ok {
		c.JSON(
			http.StatusBadRequest,
			utils.Error([]string{"User id required", "Хэрэглэгчийн ID дутуу байна"}, "id must be required"),
		)
		return
	}

	user.ID, err = strconv.ParseInt(id, 10, 64)

	if err != nil {
		c.JSON(
			http.StatusBadRequest,
			utils.Error([]string{"User id required", "Хэрэглэгчийн ID буруу байна"}, err.Error()),
		)
		return
	}

	if user.ID == 0 {
		c.JSON(
			http.StatusBadRequest,
			utils.Error([]string{"User id required", "Хэрэглэгчийн ID дутуу байна"}, "id must be required"),
		)
		return
	}

	if err := user.Get(); err != nil {
		c.JSON(
			http.StatusInternalServerError,
			utils.Error([]string{"User not found", "Хэрэглэгчийг олдсонгүй алдаа гарлаа"}, err.Error()),
		)
		return
	}

	user.Password = ""

	c.JSON(200, utils.Success([]string{"User found", "Хэрэглэгч олдлоо"}, user))
}

func GetUserList(c *gin.Context) {
	var err error
	var user user.User

	users, err := user.GetList()

	if err != nil {
		c.JSON(
			http.StatusInternalServerError,
			utils.Error([]string{"User list not found", "Хэрэглэгчийн жагсаалт олдсонгүй алдаа гарлаа"}, err.Error()),
		)
		return
	}

	var i int
	for i = 0; i < len(users); i++ {
		users[i].Password = ""
	}

	c.JSON(200, utils.Success([]string{"User list", "Хэрэглэгчийн жагсаалт"}, users))
}

func UpdateUser(c *gin.Context) {
	var user user.User

	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(
			http.StatusBadRequest,
			utils.Error([]string{"User fields required", "Хэрэглэгчийн мэдээлэл дутуу байна"}, err.Error()),
		)
		return
	}

	if err := user.Update(); err != nil {
		c.JSON(
			http.StatusInternalServerError,
			utils.Error([]string{"User update failed", "Хэрэглэгчийг хадгалахад алдаа гарлаа"}, err.Error()),
		)
		return
	}

	c.JSON(200, utils.Success([]string{"User update success", "Хэрэглэгчийг шинэчилэх амжилттай боллоо"}, struct{}{}))
}

func UpdateUserPassword(c *gin.Context) {
	var user user.User
	var input struct {
		Phone       string `json:"phone"`
		Password    string `json:"password"`
		NewPassword string `json:"new_password"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(
			http.StatusBadRequest,
			utils.Error([]string{"User fields required", "Хэрэглэгчийн мэдээлэл дутуу байна"}, err.Error()),
		)
		return
	}

	if err := config.DB.Where("phone = ? ", input.Phone).First(&user).Error; err != nil {
		c.JSON(
			http.StatusInternalServerError,
			utils.Error([]string{"User not found", "Хэрэглэгчийг олдсонгүй алдаа гарлаа"}, err.Error()),
		)
		return
	}

	if user.ID == 0 {
		c.JSON(
			http.StatusBadRequest,
			utils.Error([]string{"User not found", "Хэрэглэгчийг олдсонгүй алдаа гарлаа"}, "user not found"),
		)
		return
	}

	if input.Password != user.Password {
		c.JSON(
			http.StatusBadRequest,
			utils.Error([]string{"User password incorrect", "Хэрэглэгчийн нууц үг буруу байна"}, "password incorrect"),
		)
		return
	}

	user.Password = input.NewPassword

	if err := user.Update(); err != nil {
		c.JSON(
			http.StatusInternalServerError,
			utils.Error([]string{"User update failed", "Хэрэглэгчийг шинэчилэхэд алдаа гарлаа"}, err.Error()),
		)
		return
	}

	c.JSON(200, utils.Success([]string{"User update success", "Хэрэглэгчийг шинэчилэх амжилттай боллоо"}, struct{}{}))
}

func DeleteUser(c *gin.Context) {
	var user user.User
	var err error
	id, ok := c.GetQuery("id")

	if !ok {

		c.JSON(
			http.StatusBadRequest,
			utils.Error([]string{"User id required", "Хэрэглэгчийн ID дутуу байна"}, "id must be required"),
		)
		return
	}

	user.ID, err = strconv.ParseInt(id, 10, 64)

	if err != nil {

		c.JSON(
			http.StatusBadRequest,
			utils.Error([]string{"User id cannot be parsed", "Хэрэглэгчийн ID буруу байна"}, err.Error()),
		)
		return
	}

	if err := user.Delete(); err != nil {
		c.JSON(
			http.StatusInternalServerError,
			utils.Error([]string{"User delete failed", "Хэрэглэгчийг устгахад алдаа гарлаа"}, err.Error()),
		)
		return
	}

	c.JSON(200, utils.Success([]string{"User delete success", "Хэрэглэгчийг устгах амжилттай боллоо"}, struct{}{}))
}
