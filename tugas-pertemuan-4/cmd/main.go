package main

import (
	"fmt"
	"log"

	"tugas-pertemuan-4/config"
)

func main() {
	config.InitDatabase()
	defer config.DB.Close()

	// _, err := config.DB.Exec(`INSERT INTO mahasiswa (nama) VALUES ($1)`, "Aas Asrini")
	// if err != nil {
	// 	log.Fatal(err)
	// }

	fmt.Println("Data berhasil disimpan.")
}