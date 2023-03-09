package main

import (
	"github.com/devfurkankizmaz/go-mongo-rest-app/configs"
	"github.com/devfurkankizmaz/go-mongo-rest-app/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	configs.ConnectDB()
	router := gin.Default()
	router.Use(gin.Logger())
	routes.MovieRoute(router)
	routes.DirectorRoute(router)
	router.Run("localhost:2121")
}
