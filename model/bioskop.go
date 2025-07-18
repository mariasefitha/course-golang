package model

import "gorm.io/gorm"

type Bioskop struct {
	ID     uint    `gorm:"primaryKey" json:"id"`
	Nama   string  `json:"nama" binding:"required"`
	Lokasi string  `json:"lokasi" binding:"required"`
	Rating float32 `json:"rating"`
}

// Opsional: jika kamu mau buat validasi tambahan nanti di controller.
func (b *Bioskop) BeforeCreate(tx *gorm.DB) (err error) {
	// Tambahkan custom validation di sini kalau perlu
	return
}
