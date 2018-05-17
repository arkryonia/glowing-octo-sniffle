package main

import (
	"net/http"
	"os"

	"github.com/drxos/blog-api/routes"

	"github.com/drxos/blog-api/db"
	"github.com/drxos/blog-api/middlewares"
	"github.com/gin-gonic/gin"
)

const (
	// Port is default server port
	Port = "8000"
)

func init() {
	db.Connect()
}

func main() {
	// Configure
	r := gin.Default()

	// Middlewares
	r.Use(middlewares.Connect)
	r.Use(middlewares.ErrorHandler)

	// Routes
	r.GET("/", func(c *gin.Context) {
		c.Redirect(http.StatusMovedPermanently, "/articles")
	})

	routes.R(r)

	// Start listening
	port := Port
	if len(os.Getenv("PORT")) > 0 {
		port = os.Getenv("PORT")
	}
	r.Run(":" + port)
}
