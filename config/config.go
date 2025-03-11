package config

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {
	// Konfigurasi koneksi ke mysql
	dsn := "root:221100@tcp(127.0.0.1:3306)/todo_db?charset=utf8mb4&parseTime=True&loc=Local"

	// Membuka koneksi dengan GORM
	database, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("Gagal terhubung ke database")
	}

	fmt.Println("Berhasil terhubung ke database")

	DB = database
}
