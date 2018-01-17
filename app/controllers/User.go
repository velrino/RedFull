package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	Response "../responses"
	Models "../models"
	Config "../config"
)

func ListUser(c *gin.Context) {
	var todos []Models.User
	var _todos []Response.User

	db := Config.Database()
	db.Find(&todos)

	if len(todos) <= 0 {
		c.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "message": "No todo found!"})
		return
	}

	//transforms the todos for building a good response
	for _, item := range todos {
		_todos = append(_todos, Response.User{ID: item.ID, Name: item.Name})
	}
	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "data": _todos})
}

func GetUser(c *gin.Context) {
	//var _model []Models.User
	var _response []Response.User
	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "data": _response })
}