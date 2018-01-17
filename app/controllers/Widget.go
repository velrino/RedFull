package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	Response "../responses"
	Models "../models"	
)

func ListWidget(c *gin.Context) {
	//var _model []Models.Widget
	var _response []Response.Widget
	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "data": _response })
}

func GetWidget(c *gin.Context) {
	//var _model []Models.Widget
	var _response []Response.Widget
	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "data": _response })
}

func UpdateWidget(c *gin.Context) {
	var _todos []Models.Widget
	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "data": _todos })
}