package middleware

import (
	"encoding/json"
	"kcloudb1/internal/config"
	"kcloudb1/internal/models/user"
	"kcloudb1/internal/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

func AuthSysUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("Authorization")
		objToken, err := config.RS.Get(token).Result()

		if err != nil {
			c.JSON(http.StatusUnauthorized, utils.Error(
				[]string{"Unauthorized", "Хэрэглэгч нэвтрээгүй байна"},
				err.Error(),
			))
			c.Abort()
			return
		}

		var sysUser user.SysUser
		var sysUserDB user.SysUser

		if err := json.Unmarshal([]byte(objToken), &sysUser); err != nil {

			c.JSON(http.StatusUnauthorized, utils.Error(
				[]string{"Unauthorized", "Хэрэглэгч нэвтрээгүй байна"},
				err.Error(),
			))
			c.Abort()
			return
		}

		if err := config.DB.Find(&sysUserDB, "id = ?", sysUser.ID).Error; err != nil {

			c.JSON(http.StatusUnauthorized, utils.Error(
				[]string{"Unauthorized", "Хэрэглэгч нэвтрээгүй байна"},
				err.Error(),
			))
			c.Abort()
			return
		}

		if sysUserDB.ID == 0 && sysUserDB.UID != sysUser.UID {
			c.JSON(http.StatusForbidden, utils.Error(
				[]string{"Forbidden", "Хэрэглэгч нэвтрэх эрхгүй байна"},
				"Forbidden",
			))

			c.Abort()
			return
		}

		c.Set("sys_user_id", sysUser.ID)
		c.Next()
	}
}

func AuthUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("Authorization")
		objToken, err := config.RS.Get(token).Result()

		if err != nil {
			c.JSON(http.StatusUnauthorized,
				utils.Error(
					[]string{"Unauthorized", "Хэрэглэгч нэвтрээгүй байна"},
					err.Error(),
				))

			c.Abort()
			return
		}

		var userRS user.User
		var userDB user.User

		if err := json.Unmarshal([]byte(objToken), &userRS); err != nil {

			c.JSON(http.StatusUnauthorized,
				utils.Error(
					[]string{"Unauthorized", "Хэрэглэгч нэвтрээгүй байна"},
					err.Error(),
				))
			c.Abort()
			return
		}

		if err := config.DB.Find(&userDB, "id = ?", userRS.ID).Error; err != nil {

			c.JSON(http.StatusUnauthorized,
				utils.Error(
					[]string{"Unauthorized", "Хэрэглэгч нэвтрээгүй байна"},
					err.Error(),
				))

			c.Abort()
			return
		}

		if userDB.ID == 0 && userDB.UID != userRS.UID {

			c.JSON(http.StatusForbidden,
				utils.Error(
					[]string{"Forbidden", "Хэрэглэгч нэвтрэх эрхгүй байна"},
					"Forbidden",
				))
			c.Abort()
			return
		}

		// if userDB.RoleID >= role {

		// 	c.JSON(http.StatusForbidden,
		// 		utils.Error(
		// 			"Forbidden",
		// 			"Forbidden user permission needed",
		// 		))
		// 	c.Abort()
		// 	return
		// }

		c.Set("user_id", userRS.ID)
		c.Next()
	}
}
