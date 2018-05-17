package routes

import (
	"github.com/gin-gonic/gin"
	"gitlab.com/drxos/simple-api/handlers/todos"
)

// R exposes todos routes
func R(router *gin.Engine) *gin.Engine {
	router.GET("/todos", todos.List)
	router.POST("/todos", todos.Create)
	router.GET("/todos/:id", todos.Read)
	router.PUT("/todos/:id", todos.Update)
	router.DELETE("/todos/:id", todos.Delete)
	return router
}
