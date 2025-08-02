package config

import (
	"database/sql"
	"fmt"
	"hsi-sandbox/belajar-golang-hsi/tugas-pertemuan-4/models"
	"log"
	//_ "github.com/lib/pq"
	"tugas-pertemuan-4/models" // sesuaikan dengan module kamu
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

//var DB *sql.DB
var DB *gorm.DB
func InitDatabase() {
	var err error
	//// Format: postgres://username:password@host:port/dbname
	connStr := "postgres://odoo:odoo@localhost:5434/mahasiswa_db?sslmode=disable"
	// DB, err = sql.Open("postgres", connStr)
	// if err != nil {
	// 	log.Fatal("Tidak bisa konek ke database", err)
	// }

	DB, err = gorm.Open(postgres.Open(connStr), &gorm.Config{})
	if err != nil {
		log.Fatal("Tidak bisa konek ke database", err)
	}
	fmt.Println("Berhasil konek ke database")
	//Auto migration
	err = DB.AutoMigrate(&models.Mahasiswa{}, &models.Tugas{}, &models.Hasil{})
	if err != nil {
		log.Fatal("Gagal Migration", err)
	}
}