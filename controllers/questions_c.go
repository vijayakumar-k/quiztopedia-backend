package controllers

import (
	"net/http"
	"quiztopedia-backend/models"

	"github.com/gin-gonic/gin"
)

func GetQuestionById(c *gin.Context) {
	id := c.Params.ByName("id")
	question, err := models.GetQuestionById(id)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, question)
	}
}

func GetAllQuestions(c *gin.Context) {
	questions, err := models.GetAllQuestions()
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, questions)
	}
}

func CreateQuestion(c *gin.Context) {
	var question models.Question
	c.BindJSON(&question)
	err := models.CreateQuestion(&question)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, question)
	}
}
