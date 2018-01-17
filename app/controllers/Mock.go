package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	Response "../responses"
)

func MockFetchAll(c *gin.Context) {
	var _todos []Response.Mock
	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "data": _todos })
}