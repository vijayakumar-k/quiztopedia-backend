package routes

import (
	"net/http"
	dataaccess "quiztopedia-backend/data-access"

	"github.com/gin-gonic/gin"
)

func addQuestionsRoutes(rg *gin.RouterGroup) {
	questions := rg.Group("/questions")

	questions.GET("/", func(c *gin.Context) {
		var questions []dataaccess.Question

		questions, err := dataaccess.GetAllQuestions()

		if err != nil {
			c.JSON(http.StatusInternalServerError, err)
		} else {
			c.JSON(http.StatusOK, questions)
		}
	})
	questions.GET("/comments", func(c *gin.Context) {
		c.JSON(http.StatusOK, "users comments")
	})
	questions.GET("/pictures", func(c *gin.Context) {
		c.JSON(http.StatusOK, "users pictures")
	})
}
