package routes

import (
	"github.com/devfurkankizmaz/go-mongo-rest-app/controllers"
	"github.com/gin-gonic/gin"
)

func MovieRoute(router *gin.Engine) {
	router.POST("/films", controllers.GetFilteredData())
}
