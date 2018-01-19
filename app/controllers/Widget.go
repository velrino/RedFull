package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"gopkg.in/validator.v1"
	Response "../responses"	
	Validate "../validate"	
	Config "../config"
)

func ListWidget(c *gin.Context) {
	var list []Response.Widget
	
	Config.Database().Find(&list)
	
	if len(list) <= 0 {
		c.JSON(http.StatusNotFound, gin.H{"message": "No widget found!"})
		return
	}
	
	c.JSON(http.StatusOK,list)
}

func GetWidget(c *gin.Context) {
	var _response []Response.Widget
	ID := c.Param("id")

	Config.Database().First(&_response,ID)

	if len(_response) <= 0 {
		c.JSON(http.StatusNotFound, gin.H{"message": "No widget found!"})
		return
	}

	c.JSON(http.StatusOK, _response[0])
}

func CreateWidget(c *gin.Context) {
	var create Validate.Widget
	c.Bind(&create)
	if valid, _ := validator.Validate(create); valid {
		db := Config.Database()
		db.Save(&create)
		c.JSON(http.StatusCreated, gin.H{"message": "Widget created successfully!"})
	} else {
		c.JSON(http.StatusNotFound, gin.H{"message": "Incorrect param"})
		return		
	}
}

func UpdateWidget(c *gin.Context) {
	var create Validate.Widget
	c.Bind(&create)
	ID := c.Param("id")

	if valid, _ := validator.Validate(create); valid {
		db := Config.Database()
		db.First(&create,ID)
		db.Save(&create)
		c.JSON(http.StatusOK, gin.H{"message": "Widget updated successfully!"})
	} else {
		c.JSON(http.StatusNotFound, gin.H{"message": "Incorrect param"})
		return		
	}
}