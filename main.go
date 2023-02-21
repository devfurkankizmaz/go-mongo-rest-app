package main

import (
	"github.com/devfurkankizmaz/go-mongo-rest-app/configs"
	"github.com/devfurkankizmaz/go-mongo-rest-app/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	configs.ConnectDB()
	routes.MovieRoute(router)
	router.Run("localhost:2121")
}
