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
	nama, _ := reader.ReadString('\n')
	nama = strings.TrimSpace(nama)
	fmt.Print("Masukkan umur Anda: ")
	age, _ := reader.ReadString('\n')
	age = strings.TrimSpace(age)

	umur, err := strconv.Atoi(age)
	if err != nil {
		msgErr := "Input tidak valid!"
		fmt.Println(msgErr)
		log.WithError(err).Error(msgErr)
		return
	}

	if umur < 18 {
		fmt.Printf("Nama: %d\n", nama)
		fmt.Printf("Umur: %d\n", umur)
		msgErr := "Error: umur tidak valid (minimal 18 tahun)"
		fmt.Printf(msgErr)
		log.WithError(err).Error(msgErr)
		return
	}

	if nama == "" {
		msgErr := "Error: nama harus diisi"
        fmt.Printf(msgErr)
	    log.WithError(err).Error(msgErr)
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