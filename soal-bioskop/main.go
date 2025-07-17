package main

import (
	"soal-bioskop/config"
	"soal-bioskop/controller"
	"soal-bioskop/model"

	"github.com/gin-gonic/gin"
)

func main() {
	config.ConnectDB()
	config.DB.AutoMigrate(&model.Bioskop{})

	r := gin.Default()

	r.POST("/bioskop", controller.TambahBioskop)

	r.Run(":8080")
}
