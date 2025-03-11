package main

import (
	"fmt"
	"todo-app/config"
	"todo-app/models"
	"todo-app/routes"
)

func main() {
	// Memanggil fungsi koneksi ke database
	config.ConnectDatabase()

	// Memanggil migrasi otomatis di Go
	config.DB.AutoMigrate(&models.Todo{})
	fmt.Println("Migrasi database berhasil!")

	// Setup Router
	r := routes.SetupRouter()

	// Jalankan Server di PORT 8080
	fmt.Println("Server berjalan di https://localhost:8080")
	r.Run(":8080")
}
