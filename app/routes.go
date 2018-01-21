package app

import (
	"time"
	"github.com/appleboy/gin-jwt"
	"github.com/gin-gonic/gin"
	Controllers "github.com/velrino/RedFull/app/controllers"
)

func Routes() {
	Middleware := &jwt.GinJWTMiddleware{
		Realm:      "RedVentures",
		Key:        []byte("RedVentures"),
		Timeout:    time.Hour,
		MaxRefresh: time.Hour,
		Authenticator: func(userId string, password string, c *gin.Context) (string, bool) {
			return userId, Controllers.Auth(userId,password)
		},
		Authorizator: func(userId string, c *gin.Context) bool {
			return Controllers.AuthEmail(userId)
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
		Auth.GET("/users",  Controllers.ListUser)
		Auth.GET("/users/:id",  Controllers.GetUser)
		Auth.POST("widgets/",  Controllers.CreateWidget)
		Auth.GET("widgets/",  Controllers.ListWidget)
		Auth.GET("widgets/:id",  Controllers.GetWidget)
		Auth.PUT("widgets/:id",  Controllers.UpdateWidget)
	}
	router.Run(":7878")
}