package org_handler

import (
	"kcloudb1/internal/models/org"
	"kcloudb1/internal/models/user"
	"kcloudb1/internal/utils"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

func CreateOrg(c *gin.Context) {
	var org org.Org
	var err error

	if err := c.ShouldBindJSON(&org); err != nil {
		c.JSON(
			http.StatusBadRequest,
			utils.Error([]string{"Org fields required", "Байгууллагын мэдээлэл дутуу байна"}, err),
		)
		return
	}

	org.CreatedAt = time.Now()
	org.ExpireDate = time.Now().AddDate(1, 0, 0)

	if err = org.Create(); err != nil {
		c.JSON(
			http.StatusInternalServerError,
			utils.Error([]string{"Org creation failed", "Байгууллага бүртгэхэд алдаа гарлаа"}, err),
		)
		return
	}

	c.JSON(
		http.StatusCreated,
		utils.Success([]string{"Org created", "Байгууллага бүртгэгдлээ"}, org),
	)
}

func UpdateOrg(c *gin.Context) {
	var org org.Org
	var err error

	if err := c.ShouldBindJSON(&org); err != nil {
		c.JSON(
			http.StatusBadRequest,
			utils.Error([]string{"Org fields required", "Байгууллагын мэдээлэл дутуу байна"}, err),
		)
		return
	}

	if err = org.Update(); err != nil {

		c.JSON(
			http.StatusInternalServerError,
			utils.Error([]string{"Org update failed", "Байгууллага бүртгэхэд алдаа гарлаа"}, err),
		)
		return
	}

	c.JSON(
		200,
		utils.Success([]string{"Org updated", "Байгууллага бүртгэгдлээ"}, org),
	)
}

func DeleteOrg(c *gin.Context) {
	var org org.Org
	var err error

	id, ok := c.GetQuery("id")

	if !ok {
		c.JSON(
			http.StatusBadRequest,
			utils.Error([]string{"Org id required", "Байгууллагын id дутуу байна"}, err),
		)
		return
	}

	org.ID, err = strconv.ParseInt(id, 10, 64)

	if err != nil {
		c.JSON(
			http.StatusBadRequest,
			utils.Error([]string{"Org id cannot be parsed", "Байгууллагын id буруу байна"}, err),
		)
		return
	}

	if err = org.Delete(); err != nil {
		c.JSON(
			http.StatusInternalServerError,
			utils.Error([]string{"Org deletion failed", "Байгууллага устгахад алдаа гарлаа"}, err),
		)
		return
	}

	c.JSON(
		200,
		utils.Success([]string{"Org deleted", "Байгууллага устгагдлаа"}, struct{}{}),
	)

}

func GetOrg(c *gin.Context) {
	var org org.OrgExtend
	var err error

	id, ok := c.GetQuery("id")

	if !ok {
		c.JSON(
			http.StatusBadRequest,
			utils.Error([]string{"Org id required", "Байгууллагын id дутуу байна"}, err),
		)
		return
	}

	org.ID, err = strconv.ParseInt(id, 10, 64)

	if err != nil {
		c.JSON(
			http.StatusBadRequest,
			utils.Error([]string{"Org id cannot be parsed", "Байгууллагын id буруу байна"}, err),
		)
		return
	}

	if err = org.Get(); err != nil {
		c.JSON(
			http.StatusInternalServerError,
			utils.Error([]string{"Org not found", "Байгууллага буруу байна"}, err),
		)
		return
	}

	c.JSON(
		200,
		utils.Success([]string{"Org found", "Байгууллага бүртгэгдлээ"}, org),
	)

}

func GetOrgList(c *gin.Context) {
	var org org.OrgExtend
	var err error

	offset, _ := c.Get("offset")
	limit, _ := c.Get("limit")
	sort, _ := c.Get("sort")
	order, _ := c.Get("order")
	offsetInt, _ := strconv.Atoi(offset.(string))
	limitInt, _ := strconv.Atoi(limit.(string))

	orgs, err := org.GetList(offsetInt, limitInt, order.(string), sort.(string))

	if err != nil {
		c.JSON(
			http.StatusInternalServerError,
			utils.Error([]string{"Org list not found", "Байгууллагын жагсаалт буруу байна"}, err),
		)
		return
	}

	c.JSON(
		200,
		utils.Success([]string{"Org list", "Байгууллагын жагсаалт"}, orgs),
	)
}

func GetOrgListUser(c *gin.Context) {
	var org org.Org
	var usr user.User
	var err error
	userID, ok := c.Get("user_id")

	if !ok {
		c.JSON(
			http.StatusBadRequest,
			utils.Error([]string{"User id required", "Хэрэглэгч id дутуу байна"}, nil),
		)
		return
	}

	usr.ID = userID.(int64)

	if err != nil {

		c.JSON(
			http.StatusBadRequest,
			utils.Error([]string{"User id cannot be parsed", "Хэрэглэгч id буруу байна"}, err),
		)
		return
	}

	err = usr.Get()

	if err != nil {

		c.JSON(
			http.StatusInternalServerError,
			utils.Error([]string{"User not found", "Хэрэглэгч буруу байна"}, err),
		)
		return
	}

	org.ID = usr.KaraokeID
	err = org.Get()

	if err != nil {

		c.JSON(
			http.StatusInternalServerError,
			utils.Error([]string{"Org not found", "Байгууллага буруу байна"}, err),
		)
		return

	}

	c.JSON(
		200,
		utils.Success([]string{"Org found", "Байгууллага олдлоо"}, org),
	)

}

func CreateOrgAndUser(c *gin.Context) {
	var input org.OrgAndUserInput

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(
			http.StatusBadRequest,
			utils.Error([]string{"Org and user fields required", "Байгууллагын мэдээлэл дутуу байна"}, err),
		)
		return
	}

	orgc, err := input.Create()
	if err != nil {
		c.JSON(
			http.StatusInternalServerError,
			utils.Error([]string{"Org and user creation failed or email or phone already exist in user", "Байгууллага бүртгэхэд алдаа гарлаа. Утасны дугаар аль эсвэл имейл хаяг бүртгэлтэй байна"}, err),
		)
		return
	}

	c.JSON(200, utils.Success([]string{"Org and user created", "Байгууллага бүртгэгдлээ"}, orgc))
}
