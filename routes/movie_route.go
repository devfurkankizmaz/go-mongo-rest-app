package routes

import (
	"github.com/devfurkankizmaz/go-mongo-rest-app/controllers"
	"github.com/gin-gonic/gin"
)

func MovieRoute(router *gin.Engine) {
	router.POST("/films", controllers.GetFilteredData())         //filter movies with request body
	router.POST("/film", controllers.CreateMovie())              // create a movie
	router.GET("/films", controllers.GetAllMovies())             // get all novies
	router.GET("/film/:userId", controllers.GetFilmByID())       //get a single movie by id
	router.DELETE("/film/:userId", controllers.DeleteFilmByID()) //delete movie
	router.PUT("/film", controllers.EditFilmByID())
}
