package models

type Mahasiswa struct {
	ID uint `gorm:"primaryKey"`
	Nama string `gorm:"type:varchat(100)"`
	Tugas []Tugas `gorm:"foreignKey:MahasiswaID"`
}