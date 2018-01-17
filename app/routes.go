package app

import (
	"time"
	"github.com/appleboy/gin-jwt"
	"github.com/gin-gonic/gin"
	Controllers "./controllers"
)


func Routes() {

	Middleware := &jwt.GinJWTMiddleware{
		Realm:      "test zone",
		Key:        []byte("secret key"),
		Timeout:    time.Hour,
		MaxRefresh: time.Hour,
		Authenticator: func(userId string, password string, c *gin.Context) (string, bool) {
			if (userId == "admin" && password == "admin") || (userId == "test" && password == "test") {
				return userId, true
			}

			return userId, false
		},
		Authorizator: func(userId string, c *gin.Context) bool {
			if userId == "admin" {
				return true
			}

			return false
		},
		Unauthorized: func(c *gin.Context, code int, message string) {
			c.JSON(code, gin.H{
				"code":    code,
				"message": message,
			})
		},

		TokenLookup: "header:Authorization",
		TokenHeadName: "Bearer",
		TimeFunc: time.Now,
	}

	router := gin.Default()

	router.GET("/", Controllers.MockFetchAll)
	router.POST("/login", Middleware.LoginHandler)

	Auth := router.Group("/api")
	Auth.Use(Middleware.MiddlewareFunc())
	{
		Auth.GET("/token", Middleware.RefreshHandler)

		Users := router.Group("/users")
		{
			Users.GET("/",  Controllers.MockFetchAll)
			Users.GET("/:id",  Controllers.MockFetchAll)
		}

		Widgets := router.Group("/widgets")
		{
			Widgets.POST("/",  Controllers.MockFetchAll)
			Widgets.GET("/",  Controllers.MockFetchAll)
			Widgets.GET("/:id",  Controllers.MockFetchAll)
		}

	}
	router.Run(":7878")
}