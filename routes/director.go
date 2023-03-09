package routes

import (
	"github.com/devfurkankizmaz/go-mongo-rest-app/controllers"
	"github.com/gin-gonic/gin"
)

func DirectorRoute(router *gin.Engine) {
	router.GET("/directors", controllers.GetAllDirectors())
	router.POST("/director", controllers.CreateDirector())
}
