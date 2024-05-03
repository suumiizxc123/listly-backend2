package middleware

import (
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
)

func Paginate() gin.HandlerFunc {
	return func(c *gin.Context) {
		page, ok := c.GetQuery("page")
		if !ok {
			page = "1"
		}

		limit, ok := c.GetQuery("limit")

		if !ok {
			limit = "20"
		}

		sort, ok := c.GetQuery("sort")

		if !ok {
			sort = "id"
		}

		order, ok := c.GetQuery("order")

		if !ok {
			order = "asc"
		}

		if order != "asc" && order != "desc" {
			order = "asc"
		}

		if sort != "id" && sort != "created_at" && sort != "updated_at" {
			sort = "id"
		}

		pageInt, _ := strconv.Atoi(page)
		limitInt, _ := strconv.Atoi(limit)

		offset := (pageInt - 1) * limitInt

		c.Set("page", page)
		c.Set("limit", limit)
		c.Set("sort", sort)
		c.Set("order", order)
		c.Set("offset", fmt.Sprintf("%d", offset))

		c.Next()
	}
}
