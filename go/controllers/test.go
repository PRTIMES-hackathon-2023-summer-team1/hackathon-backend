package controllers

import (
	"net/http"

	"github.com/PRTIMES-hackathon-2023-summer-team1/hackathon-backend/models"
	"github.com/gin-gonic/gin"
)

type TestController struct{}

var testModel = new(models.TestModel)

func (t TestController) Set(c *gin.Context) {
	var testInfo models.Test
	err := c.ShouldBindJSON(&testInfo)
	if err != nil {
		c.Error(err).SetType(gin.ErrorTypePublic)
		return
	}
	err = testModel.Set(testInfo)
	if err != nil {
		c.Error(err).SetType(gin.ErrorTypePublic)
		return
	}
	c.JSON(http.StatusOK, testInfo)
	return
}
