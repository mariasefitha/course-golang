package config

import (
	"fmt"
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() {
	var err error

	// Ganti sesuai konfigurasi database kamu
	host := "localhost"
	port := 5432
	user := "postgres"
	password := "msfitha0880" // Ganti dengan password PostgreSQL kamu
	dbname := "bioskop_db"    // Nama database yang akan kamu buat

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=disable",
		host, user, password, dbname, port)

	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Gagal koneksi ke database: ", err)
		os.Exit(1)
	}
	fmt.Println("Berhasil koneksi ke database")
}
