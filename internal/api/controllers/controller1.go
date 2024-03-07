package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func HellowNeverSitup(c *gin.Context) {
	c.JSON(http.StatusOK, ResponseOK("Hellow NeverSitup"))
}
