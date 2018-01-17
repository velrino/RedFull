package main

import (
	"github.com/gin-gonic/gin"
	App "./app"
	Controllers "./app/controllers"
)

func main() {

	App.Migrations()

	router := gin.Default()

	v1 := router.Group("/api")
	{
		v1.GET("/", Controllers.MockFetchAll)
	}
	router.Run(":8979")
}
