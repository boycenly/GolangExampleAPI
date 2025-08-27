package database

import (
	"log"

	"myfiberapi/models" // pastikan path sesuai module Anda

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
	dsn := "root:@tcp(localhost)/golang?charset=utf8mb4&parseTime=True&loc=Local"
	var err error
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Gagal koneksi database:", err)
	}

	// Auto migrate semua model
	if err := DB.AutoMigrate(&models.User{}); err != nil {
		log.Fatal("Gagal migrate:", err)
	}

	log.Println("Database connected & migrated successfully 🚀")
}
