package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func HealthCheck(c *gin.Context) {
	c.JSON(http.StatusOK, map[string]interface{}{
		"success": true,
	})
}
