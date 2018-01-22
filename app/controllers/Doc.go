package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func Doc(c *gin.Context) {
	c.Redirect(http.StatusMovedPermanently, "/doc")
}