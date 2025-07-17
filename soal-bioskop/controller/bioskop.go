package controller

import (
	"net/http"
	"soal-bioskop/config"
	"soal-bioskop/model"

	"github.com/gin-gonic/gin"
)

func TambahBioskop(c *gin.Context) {
	var bioskop model.Bioskop

	if err := c.ShouldBindJSON(&bioskop); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Data tidak valid atau ada field kosong"})
		return
	}

	if err := config.DB.Create(&bioskop).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal menambahkan bioskop"})
		return
	}

	c.JSON(http.StatusOK, bioskop)
}
