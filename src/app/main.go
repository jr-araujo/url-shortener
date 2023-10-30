package main

import (
	"github.com/gin-gonic/gin"
	"urlshortener.com/devgym/jr/controllers"
	database "urlshortener.com/devgym/jr/repository"
)

func main() {
	DB := database.Init()
	ctlr := controllers.New(DB)

	server := gin.Default()

	// Routes
	server.POST("/shorteners", ctlr.GenerateUrlShorten)
	server.GET("/shorteners/:code", ctlr.RedirectOriginalUrl)

	// Get the server up
	server.Run()
}
