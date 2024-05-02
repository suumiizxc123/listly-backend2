package org_handler

import (
	"kcloudb1/internal/models/org"
	"kcloudb1/internal/utils"
	"net/http"
	"strconv"

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
		utils.Success("Org deleted", org),
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
	var org org.Org
	var err error

	orgs, err := org.GetList()

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
