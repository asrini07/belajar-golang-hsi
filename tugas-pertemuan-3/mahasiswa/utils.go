package mahasiswa

// - Buat fungsi `hitungRataRata(nilai ...int) float64` → variadic function.
func hitungRataRata(nilai ...int) float64 {
	if len(nilai) == 0 {
		return 0
	}

	totalNilai := 0
	for _, n := range nilai {
		totalNilai += n
	}

	nilaiRata := float64(totalNilai) / float64(len(nilai))
	return nilaiRata
}

// - Buat fungsi `BuatMahasiswa(nama string, umur int, nilai ...int) Mahasiswa`
// - Gunakan pointer pada struct `Mahasiswa`  untuk mengubah umur mahasiswa di dalam fungsi.
func BuatMahasiswa(nama string, umur int, nilai ...int) *Mahasiswa {
	mahasiswa := &Mahasiswa{
		Nama: nama,
		umur: umur,
		Nilai: nilai,
		nilaiAvg: hitungRataRata(nilai...),
	}

	return mahasiswa
}