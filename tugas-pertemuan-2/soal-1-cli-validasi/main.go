package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	log "github.com/sirupsen/logrus"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Masukkan nama Anda: ")
	nama, err := reader.ReadString('\n')
	nama = strings.TrimSpace(nama)

	if nama == "" {
		msgError := "Error: nama harus diisi"
		log.Error(msgError)
		fmt.Println(msgError)
		return
	}

	fmt.Print("Masukkan umur Anda: ")
	inputUmur, err := reader.ReadString('\n')
	if err != nil {
		log.Errorf("Gagal membaca umur: %v", err)
		return
	}
	inputUmur = strings.TrimSpace(inputUmur)

	umur, err := strconv.Atoi(inputUmur)
	if err != nil {
		msgError := "Error: input umur harus berupa angka"
		log.Error(msgError)
		fmt.Println(msgError)
		return
	}

	if umur < 18 {
		msgError := "Error: umur tidak valid (minimal 18 tahun)"
		log.Error(msgError)
		fmt.Println(msgError)
		return
	}

	infoSuccess(nama, umur)
	
}

func infoSuccess(nama string, umur int) {
	fmt.Printf("Nama: %s\n", nama)
	fmt.Printf("Umur: %d\n", umur)
	fmt.Printf("Selamat Datang, %s!\n", nama)
	log.Infof("Data %s berhasil masuk dengan umur %d", nama, umur)
}