package controller

import (
	"net/http"
	"strconv"
	"tugas-14-bioskop/config"
	"tugas-14-bioskop/model"

	"github.com/gin-gonic/gin"
)

// CREATE
func CreateBioskop(c *gin.Context) {
	var bioskops []model.Bioskop // pakai slice

	if err := c.ShouldBindJSON(&bioskops); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := config.DB.Create(&bioskops).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal menyimpan data"})
		return
	}

	c.JSON(http.StatusCreated, bioskops)
}

// READ - All
func GetAllBioskop(c *gin.Context) {
	var bioskops []model.Bioskop

	if err := config.DB.Find(&bioskops).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal mengambil data"})
		return
	}

	c.JSON(http.StatusOK, bioskops)
}

// READ - by ID
func GetBioskopByID(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID tidak valid"})
		return
	}

	var bioskop model.Bioskop
	if err := config.DB.First(&bioskop, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Data tidak ditemukan"})
		return
	}

	c.JSON(http.StatusOK, bioskop)
}

// UPDATE
func UpdateBioskop(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID tidak valid"})
		return
	}

	var bioskop model.Bioskop
	if err := config.DB.First(&bioskop, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Data tidak ditemukan"})
		return
	}

	var updateData model.Bioskop
	if err := c.ShouldBindJSON(&updateData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	bioskop.Nama = updateData.Nama
	bioskop.Lokasi = updateData.Lokasi
	bioskop.Rating = updateData.Rating

	config.DB.Save(&bioskop)

	c.JSON(http.StatusOK, bioskop)
}

// DELETE
func DeleteBioskop(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID tidak valid"})
		return
	}

	var bioskop model.Bioskop
	if err := config.DB.First(&bioskop, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Data tidak ditemukan"})
		return
	}

	config.DB.Delete(&bioskop)

	c.JSON(http.StatusOK, gin.H{"message": "Data berhasil dihapus"})
}
