package mahasiswa

import "fmt"

// - Di package `mahasiswa`, buat variabel `maxNilai int = 100` (huruf kecil = private).
// - Buat variabel `Versi string = "v1.0.0"` yang bisa langsung diakses dari `main.go`.
var maxNilai int = 100
var Versi string = "v1.0.0"

type Mahasiswa struct {
	Nama string
	Nilai []int
	umur int
	nilaiAvg float64
}

type Deskripsi interface {
    Info() string
    RataRata() float64
    GetUmur() int
}

func (m Mahasiswa) Info() string {
	return fmt.Sprintf("Nama: %s, Umur: %d, \nNilai Rata-Rata: %.2f", m.Nama, m.umur, m.nilaiAvg)
}

func (m Mahasiswa) GetUmur() int {
	return m.umur
}

func (m Mahasiswa) RataRata() float64 {
	return m.nilaiAvg
}

// - Buat fungsi `GetMaxNilai() int` agar variabel `maxNilai` bisa diakses dari luar.
func GetMaxNilai() int {
	return maxNilai
}



// - Buat struct `Mahasiswa` dengan atribut:
//     - `Nama string`
//     - `Nilai []int`
//     - `umur int`→ huruf kecil
//     - `nilaiAvg float64` → huruf kecil
// - Buat interface `Deskripsi`:
//     type Deskripsi interface {
//         Info() string
//         RataRata() float64
//         GetUmur() int
//     }
// - Implementasikan interface di struct `Mahasiswa`.



