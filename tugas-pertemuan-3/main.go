package main

import (
	"fmt"
	"tugas-pertemuan-3/mahasiswa"
)


func main() {

	// Nama, umur
	// Nilai rata-rata →  menggunakan objek Mahasiswa menggunakan fungsi RataRata
	mhs1 := mahasiswa.BuatMahasiswa("Upin", 21, 80, 75, 70)
	mhs2 := mahasiswa.BuatMahasiswa("Ipin", 19, 95, 90, 92)
	mhs3 := mahasiswa.BuatMahasiswa("Apin", 20, 97, 85, 72)

	//buat jadi slice
	allMhs := []*mahasiswa.Mahasiswa{mhs1, mhs2, mhs3}

	//hitung total umur
	totalUmur := 0
	for _, m := range allMhs {
		fmt.Println(m.Info())
		fmt.Println("-----------------------")
		totalUmur += m.GetUmur()
	}

	fmt.Println("########################")
	fmt.Println("Versi Package:", mahasiswa.Versi)
    fmt.Println("Nilai Maksimum:", mahasiswa.GetMaxNilai())
	fmt.Println("Total Umur Mahasiswa:", totalUmur)

}


// Import package mahasiswa.
// Buat beberapa objek Mahasiswa menggunakan fungsi BuatMahasiswa.
// Tampilkan hasil:
// Nama, umur
// Nilai rata-rata →  menggunakan objek Mahasiswa menggunakan fungsi RataRata
// Versi package
// Nilai maksimal (dari GetMaxNilai)
// Buat closure function untuk menghitung jumlah seluruh umur mahasiswa.
// Contoh output
// Nama: Ali, Umur: 20
// Rata-rata nilai: 85.00
// ---
// Nama: Siti, Umur: 22
// Rata-rata nilai: 81.67
// ---
// Versi Package: v1.0.0
// Nilai Maksimum: 100
// Total Umur Mahasiswa: 42