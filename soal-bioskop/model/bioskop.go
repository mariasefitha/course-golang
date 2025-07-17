package model

type Bioskop struct {
	ID     int     `json:"id" gorm:"primaryKey"`
	Nama   string  `json:"nama" binding:"required"`
	Lokasi string  `json:"lokasi" binding:"required"`
	Rating float32 `json:"rating"`
}
