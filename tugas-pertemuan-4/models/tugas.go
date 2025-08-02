package models

type Tugas struct {
	ID uint `gorm:"primaryKey"`
	Judul string `gorm:"type:varchar(100)"`
	Deskripsi string
	MahasiswaId uint
	Hasil Hasil `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}