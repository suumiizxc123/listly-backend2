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
			utils.Error("Org fields required", err),
		)
		return
	}

	org.CreatedAt = time.Now()
	org.ExpireDate = time.Now().AddDate(1, 0, 0)

	if err = org.Create(); err != nil {
		c.JSON(
			http.StatusInternalServerError,
			utils.Error("Org creation failed", err),
		)
		return
	}

	c.JSON(
		http.StatusCreated,
		utils.Success("Org created", org),
	)
}

func UpdateOrg(c *gin.Context) {
	var org org.Org
	var err error

	if err := c.ShouldBindJSON(&org); err != nil {
		c.JSON(
			http.StatusBadRequest,
			utils.Error("Org fields required", err),
		)
		return
	}

	if err = org.Update(); err != nil {

		c.JSON(
			http.StatusInternalServerError,
			utils.Error("Org update failed", err),
		)
		return
	}

	c.JSON(
		200,
		utils.Success("Org updated", org),
	)
}

func DeleteOrg(c *gin.Context) {
	var org org.Org
	var err error

	id, ok := c.GetQuery("id")

	if !ok {
		c.JSON(
			http.StatusBadRequest,
			utils.Error("Org id required", err),
		)
		return
	}

	org.ID, err = strconv.ParseInt(id, 10, 64)

	if err != nil {
		c.JSON(
			http.StatusBadRequest,
			utils.Error("Org id cannot be parsed", err),
		)
		return
	}

	if err = org.Delete(); err != nil {
		c.JSON(
			http.StatusInternalServerError,
			utils.Error("Org deletion failed", err),
		)
		return
	}

	c.JSON(
		200,
		utils.Success("Org deleted", struct{}{}),
	)

}

func GetOrg(c *gin.Context) {
	var org org.Org
	var err error

	id, ok := c.GetQuery("id")

	if !ok {
		c.JSON(
			http.StatusBadRequest,
			utils.Error("Org id required", err),
		)
		return
	}

	org.ID, err = strconv.ParseInt(id, 10, 64)

	if err != nil {
		c.JSON(
			http.StatusBadRequest,
			utils.Error("Org id cannot be parsed", err),
		)
		return
	}

	if err = org.Get(); err != nil {
		c.JSON(
			http.StatusInternalServerError,
			utils.Error("Org not found", err),
		)
		return
	}

	c.JSON(
		200,
		utils.Success("Org found", org),
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
			utils.Error("Org list not found", err),
		)
		return
	}

	c.JSON(
		200,
		utils.Success("Org list", orgs),
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
			utils.Error("User id required", nil),
		)
		return
	}

	usr.ID = userID.(int64)

	if err != nil {

		c.JSON(
			http.StatusBadRequest,
			utils.Error("User id cannot be parsed", err),
		)
		return
	}

	err = usr.Get()

	if err != nil {

		c.JSON(
			http.StatusInternalServerError,
			utils.Error("User not found", err),
		)
		return
	}

	org.ID = usr.KaraokeID
	err = org.Get()

	if err != nil {

		c.JSON(
			http.StatusInternalServerError,
			utils.Error("Org not found", err),
		)
		return

	}

	c.JSON(
		200,
		utils.Success("Org found", org),
	)

}

func CreateOrgAndUser(c *gin.Context) {
	var input org.OrgAndUserInput

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(
			http.StatusBadRequest,
			utils.Error("Org and user fields required", err),
		)
		return
	}

	if err := input.Create(); err != nil {
		c.JSON(
			http.StatusInternalServerError,
			utils.Error("Org and user creation failed", err),
		)
		return
	}

	c.JSON(200, utils.Success("Org and user created", struct{}{}))
}
