package controllers

import (
	"net/http"

	"github.com/PRTIMES-hackathon-2023-summer-team1/hackathon-backend/models"
	"github.com/PRTIMES-hackathon-2023-summer-team1/hackathon-backend/repository"
	"github.com/gin-gonic/gin"
)


type TestController struct{
	testModelRepository repository.ITestRepository
}

func NewTestController(repo repository.ITestRepository) *TestController {
	return &TestController{testModelRepository: repo}
}

func (t TestController) Set(c *gin.Context) {
	var testInfo models.TestModel
	err := c.ShouldBindJSON(&testInfo)
	if err != nil {
		c.Error(err).SetType(gin.ErrorTypePublic)
		return
	}
	err = t.testModelRepository.Set(testInfo)
	if err != nil {
		c.Error(err).SetType(gin.ErrorTypePublic)
		return
	}
	c.JSON(http.StatusOK, testInfo)
}
