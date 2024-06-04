package main

import (
	"kcloudb1/internal/config"
	"kcloudb1/internal/routes/client_route"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		c.Header("Access-Control-Allow-Headers", "Content-Type, Authorization")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(http.StatusNoContent)
			return
		}

		c.Next()
	}
}

func main() {
	config.ConnectDatabase()
	config.RedisConfig()

	r := gin.Default()

	r.Use(Cors())
	// gin.DefaultWriter = io.MultiWriter(f)

	v1 := r.Group("/api/v1")

	client_route.ClientRoute(v1)

	// user_route.SysUserRoute(v1)
	// user_route.UserRoute(v1)
	// user_route.ServiceLogRoute(v1)

	// song_route.SongRoute(v1)
	// song_route.SongCategoryRoute(v1)
	// song_route.SongCategoryCombinationRoute(v1)
	// song_route.SongCounterRoute(v1)
	// song_route.SongLanguageRoute(v1)

	// org_route.OrgRoute(v1)
	// org_route.OrgSysRoute(v1)
	// org_route.OrgUserRoute(v1)
	// org_route.OrgAccountRoute(v1)
	// org_route.OrgAccountTxnRoute(v1)
	// org_route.OrgAccountTxnLogRoute(v1)

	// common_route.LanguageRoute(v1)

	// artist_route.ArtistTypeRoute(v1)
	// artist_route.ArtistSongRoute(v1)
	// artist_route.ArtistProfileRoute(v1)
	// artist_route.ArtistMemberRoute(v1)
	// artist_route.ArtistMemberSongRoute(v1)
	// artist_route.ArtistBandRoute(v1)
	// artist_route.ArtistBandCombinationRoute(v1)

	// moscap_route.MoscapUserRoute(v1)
	// moscap_route.MoscapUserLogRoute(v1)

	r.Run(":8080")
}
