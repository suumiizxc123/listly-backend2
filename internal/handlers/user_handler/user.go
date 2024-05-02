package user_handler

import (
	"kcloudb1/internal/models/user"
	"kcloudb1/internal/utils"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func CreateUser(c *gin.Context) {
	var user user.SysUser

	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(
			http.StatusBadRequest,
			utils.Error("User fields required", err.Error()),
		)
		return
	}

	if err := user.Create(); err != nil {
		c.JSON(
			http.StatusInternalServerError,
			utils.Error("User creation failed", err.Error()),
		)
		return
	}

	c.JSON(200, utils.Success("User created", user))
}

func GetUser(c *gin.Context) {
	var err error
	var user user.SysUser
	id, ok := c.GetQuery("id")

	if !ok {
		c.JSON(
			http.StatusBadRequest,
			utils.Error("User id required", "id must be required"),
		)
		return
	}

	user.ID, err = strconv.ParseInt(id, 10, 64)

	if err != nil {
		c.JSON(
			http.StatusBadRequest,
			utils.Error("User id required", err.Error()),
		)
		return
	}

	if user.ID == 0 {
		c.JSON(
			http.StatusBadRequest,
			utils.Error("User id required", "id must be required"),
		)
		return
	}

	if err := user.Get(); err != nil {
		c.JSON(
			http.StatusInternalServerError,
			utils.Error("User not found", err.Error()),
		)
		return
	}

	c.JSON(200, utils.Success("User found", user))
}

func GetUserList(c *gin.Context) {
	var err error
	var user user.SysUser

	users, err := user.GetList()

	if err != nil {
		c.JSON(
			http.StatusInternalServerError,
			utils.Error("User list not found", err.Error()),
		)
		return
	}

	c.JSON(200, utils.Success("User list", users))
}

func UpdateUser(c *gin.Context) {
	var user user.SysUser

	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(
			http.StatusBadRequest,
			utils.Error("User fields required", err.Error()),
		)
		return
	}

	if err := user.Update(); err != nil {
		c.JSON(
			http.StatusInternalServerError,
			utils.Error("User update failed", err.Error()),
		)
		return
	}

	c.JSON(200, utils.Success("User updated", user))
}

func DeleteUser(c *gin.Context) {
	var user user.SysUser

	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(
			http.StatusBadRequest,
			utils.Error("User fields required", err.Error()),
		)
		return
	}

	if err := user.Delete(); err != nil {
		c.JSON(
			http.StatusInternalServerError,
			utils.Error("User delete failed", err.Error()),
		)
		return
	}

	c.JSON(200, utils.Success("User deleted", user))
}
