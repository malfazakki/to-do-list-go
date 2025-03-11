package controllers

import (
	"net/http"
	"todo-app/config"
	"todo-app/models"

	"github.com/gin-gonic/gin"
)

// GET /todos - ambil semua To-Do
func GetTodos(c *gin.Context) {
	var todos []models.Todo
	config.DB.Find(&todos)

	// Buat response Todos
	

	c.JSON(http.StatusOK, gin.H{"todos": todos})
}

// GET /todos/:id - Ambil satu To-Do berdasarkan ID
func GetTodo(c *gin.Context) {
	var todo models.Todo
	id := c.Param("id")

	if err := config.DB.First(&todo, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Data tidak"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"todo": todo})
}

// POST /todos - Tambah To-Do baru
func CreateTodo(c *gin.Context) {
	var todo models.Todo

	if err := c.ShouldBindJSON(&todo); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	config.DB.Create(&todo)
	c.JSON(http.StatusOK, gin.H{"todo": todo})
}

// PUT /todos:id - Update To-Do
func UpdateTodo(c *gin.Context) {
	var todo models.Todo
	id := c.Param("id")

	if err := config.DB.First(&todo, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Data tidak ditemukan"})
		return
	}

	if err := c.ShouldBindJSON(&todo); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	config.DB.Save(&todo)
	c.JSON(http.StatusOK, gin.H{"todo": todo})
}

func DeleteTodo(c *gin.Context) {
	var todo models.Todo
	id := c.Param("id")

	if err := config.DB.First(&todo, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Data tidak ditemukan"})
		return
	}

	config.DB.Delete(todo)
	c.JSON(http.StatusOK, gin.H{"message": ""})
}
