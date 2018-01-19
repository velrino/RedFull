package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	Response "../responses"
	Config "../config"
)

func ListUser(c *gin.Context) {
	var list []Response.User

	Config.Database().Find(&list)

	if len(list) <= 0 {
		c.JSON(http.StatusNotFound, gin.H{"message": "No users found!"})
		return
	}

	c.JSON(http.StatusOK,list)
}

func GetUser(c *gin.Context) {
	var _response []Response.User
	ID := c.Param("id")
	
	Config.Database().First(&_response,ID)

	if len(_response) <= 0 {
		c.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "message": "No User found!"})
		return
	}

	c.JSON(http.StatusOK, _response[0])
}