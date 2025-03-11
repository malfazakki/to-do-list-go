package routes

import (
	"todo-app/controllers"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	// Routes untuk To-Do List
	r.GET("/", controllers.GetTodos)
	r.POST("/todos", controllers.CreateTodo)
	r.GET("/todos/:id", controllers.GetTodo)
	r.PUT("/todos/:id", controllers.UpdateTodo)
	r.DELETE("/todos/:id", controllers.DeleteTodo)

	return r
}
