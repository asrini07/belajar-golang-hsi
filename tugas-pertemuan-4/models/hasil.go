package models

type Hasil struct {
	ID uint `gorm:"primaryKey"`
	TugasID uint
	Nilai float32
	
}