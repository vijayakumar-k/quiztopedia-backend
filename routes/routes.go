package routes

import (
	"net/http"
	"quiztopedia-backend/controllers"

	"github.com/gin-gonic/gin"
)

var (
	router = gin.Default()
)

// Run will start the server
func Run() {
	getRoutes()
	router.Run(":5000")
}

func getRoutes() {
	v1 := router.Group("/v1")
	addQuestionsRoutes(v1)
	addPingRoutes(v1)
}

func addQuestionsRoutes(rg *gin.RouterGroup) {
	questions := rg.Group("/questions")

	questions.GET("/", controllers.GetAllQuestions)
	questions.GET("/:id", controllers.GetQuestionById)
	questions.POST("/", controllers.CreateQuestion)
}

func addPingRoutes(rg *gin.RouterGroup) {
	ping := rg.Group("/ping")

	ping.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, "pong")
	})
}
