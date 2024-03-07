package controllers

import (
	"net/http"
	"neversitup-test-template/internal/pkg/persistence"

	"github.com/gin-gonic/gin"
)

func Hellow(c *gin.Context) {
	c.JSON(http.StatusOK, ResponseOK("Hellow NeverSitup"))
}

func Test2(c *gin.Context) {
	input := c.Query("input")
	if input == "" {
		c.JSON(http.StatusBadRequest, ResponseError("param is require"))
		return
	}
	t := persistence.Test()
	result := t.Test2(input)
	c.JSON(http.StatusOK, result)
}

func Test3(c *gin.Context) {
	input := c.Query("input")
	if input == "" {
		c.JSON(http.StatusBadRequest, ResponseError("param is require"))
		return
	}
	t := persistence.Test()
	result := t.Test3(input)
	c.JSON(http.StatusOK, result)
}

func Test4(c *gin.Context) {
	var param map[string][]string
	err := c.BindJSON(&param)
	if err != nil {
		c.JSON(http.StatusBadRequest, ResponseError("param is require"))
		return
	}
	t := persistence.Test()
	result1, result2 := t.Test4(param["input"])
	c.JSON(http.StatusOK, map[string]interface{}{"smileys": result1, "countOfSmileys": result2})
}
