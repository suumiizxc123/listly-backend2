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
			utils.Error("User fields required", err.Error()),
		)
		return
	}

	if ok := user.CheckEmailAndPhoneNotExist(); !ok {

		c.JSON(
			http.StatusInternalServerError,
			utils.Error("User creation failed", "email or phone already exist"),
		)
		return
	}

	user.UID = uuid.New().String()

	if err := user.Create(); err != nil {
		c.JSON(
			http.StatusInternalServerError,
			utils.Error("User creation failed", err.Error()),
		)
		return
	}

	user.Password = ""

	c.JSON(200, utils.Success("User created", user))
}

func LoginUser(c *gin.Context) {
	var input user.UserLoginInput
	var user user.User

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(
			http.StatusBadRequest,
			utils.Error("Login fields required", err.Error()),
		)
		return
	}

	user, err := user.Login(input.Phone, input.Password)

	if err != nil {
		c.JSON(
			http.StatusInternalServerError,
			utils.Error("Login failed", err.Error()),
		)
		return
	}

	user.Password = ""

	if user.IsActive == 0 {

		c.JSON(
			http.StatusForbidden,
			utils.Error("Login failed", "user not active"),
		)
		return
	}

	token := uuid.New().String()

	jsonUser, err := user.MarshalJSON()

	if err != nil {

		c.JSON(
			http.StatusInternalServerError,
			utils.Error("Login failed", err.Error()),
		)
		return
	}

	if err := config.RS.Set(token, jsonUser, 12*time.Hour).Err(); err != nil {

		c.JSON(
			http.StatusInternalServerError,
			utils.Error("Login failed", err.Error()),
		)
		return
	}

	user.Token = token

	c.JSON(200, utils.Success("Login successful", user))
}

func GetUser(c *gin.Context) {
	var err error
	var user user.User
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

	user.Password = ""

	c.JSON(200, utils.Success("User found", user))
}

func GetUserList(c *gin.Context) {
	var err error
	var user user.User

	users, err := user.GetList()

	if err != nil {
		c.JSON(
			http.StatusInternalServerError,
			utils.Error("User list not found", err.Error()),
		)
		return
	}

	var i int
	for i = 0; i < len(users); i++ {
		users[i].Password = ""
	}

	c.JSON(200, utils.Success("User list", users))
}

func UpdateUser(c *gin.Context) {
	var user user.User

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

	c.JSON(200, utils.Success("User updated", struct{}{}))
}

func DeleteUser(c *gin.Context) {
	var user user.User
	var err error
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
			utils.Error("User id cannot be parsed", err.Error()),
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

	c.JSON(200, utils.Success("User deleted", struct{}{}))
}
