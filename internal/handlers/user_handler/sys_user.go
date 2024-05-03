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
			utils.Error("System user fields required", err.Error()),
		)

		return
	}

	sysUser.UID = uuid.New().String()

	if ok := sysUser.CheckEmailAndPhoneNotExist(); !ok {

		c.JSON(
			http.StatusInternalServerError,
			utils.Error("System user creation failed", "email or phone already exist"),
		)

		return
	}

	if err := sysUser.Create(); err != nil {
		c.JSON(
			http.StatusInternalServerError,
			utils.Error("System user creation failed", err.Error()),
		)

		return
	}

	c.JSON(200, utils.Success("System user created", struct{}{}))

}

func GetSysUser(c *gin.Context) {

	var err error
	var sysUser user.SysUser
	id, ok := c.GetQuery("id")

	if !ok {
		c.JSON(
			http.StatusBadRequest,
			utils.Error("System user id required", "id must be required"),
		)
		return
	}

	sysUser.ID, err = strconv.ParseInt(id, 10, 64)

	if err != nil {
		c.JSON(
			http.StatusBadRequest,
			utils.Error("System user id cannot be parsed", err.Error()),
		)
		return
	}

	if err := sysUser.Get(); err != nil {
		c.JSON(
			http.StatusInternalServerError,
			utils.Error("System user not found", err.Error()),
		)
		return
	}

	if sysUser.ID == 0 {
		c.JSON(
			http.StatusNoContent,
			utils.Success("No content", nil),
		)
		return
	}

	sysUser.Password = ""

	c.JSON(200, utils.Success("System user", sysUser))
}

func GetSysUserList(c *gin.Context) {

	// var err error
	var sysUser user.SysUser

	sysUsers, err := sysUser.GetList()

	if err != nil {
		c.JSON(
			http.StatusInternalServerError,
			utils.Error("System user list not found", err.Error()),
		)
		return
	}

	var i int
	for i = 0; i < len(sysUsers); i++ {
		sysUsers[i].Password = ""
	}

	c.JSON(200, utils.Success("System user list", sysUsers))

}

func LoginSysUser(c *gin.Context) {

	var input user.SysUserLoginInput
	var sysUser user.SysUser

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(
			http.StatusBadRequest,
			utils.Error("System user login fields required", err.Error()),
		)
		return
	}

	sysUser, err := sysUser.Login(input.Phone, input.Password)

	// add redis token generate
	if err != nil {

		c.JSON(
			http.StatusInternalServerError,
			utils.Error("System user login failed", err.Error()),
		)
		return
	}

	sysUser.Password = ""

	if sysUser.IsActive == 0 {

		c.JSON(
			http.StatusForbidden,
			utils.Error("System user login failed", "inactive user"),
		)
		return
	}

	token := uuid.New().String()

	jsonSysUser, err := sysUser.MarshalJSON()
	if err != nil {

		c.JSON(
			http.StatusInternalServerError,
			utils.Error("System user login failed in marshal json", err.Error()),
		)
		return
	}

	if err := config.RS.Set(token, string(jsonSysUser), 12*time.Hour).Err(); err != nil {

		c.JSON(
			http.StatusInternalServerError,
			utils.Error("System user login failed in set redis", err.Error()),
		)

		return
	}

	sysUser.Token = token

	c.JSON(200, utils.Success("System user login", sysUser))

}

func UpdateSysUser(c *gin.Context) {
	var sysUser user.SysUser

	if err := c.ShouldBindJSON(&sysUser); err != nil {
		c.JSON(
			http.StatusBadRequest,
			utils.Error("System user fields required", err.Error()),
		)
		return
	}

	if err := sysUser.Update(); err != nil {

		c.JSON(
			http.StatusInternalServerError,
			utils.Error("System user update failed", err.Error()),
		)
		return
	}

	c.JSON(200, utils.Success("System user updated", sysUser))

}

func DeleteSysUser(c *gin.Context) {
	var err error
	var sysUser user.SysUser

	id, ok := c.GetQuery("id")

	if !ok {
		c.JSON(
			http.StatusBadRequest,
			utils.Error("System user id required", "id must be required"),
		)
		return
	}

	sysUser.ID, err = strconv.ParseInt(id, 10, 64)

	if err != nil {
		c.JSON(
			http.StatusBadRequest,
			utils.Error("System user id cannot be parsed", err.Error()),
		)
		return
	}
	if err := sysUser.Delete(); err != nil {

		c.JSON(
			http.StatusInternalServerError,
			utils.Error("System user delete failed", err.Error()),
		)
		return
	}

	c.JSON(200, utils.Success("System user deleted", struct{}{}))
}
