package main

import (
	"tugas-14-bioskop/config"
	"tugas-14-bioskop/controller"
	"tugas-14-bioskop/model"

	"github.com/gin-gonic/gin"
)

func main() {
	config.ConnectDB()
	config.DB.AutoMigrate(&model.Bioskop{})

	r := gin.Default()

	api := r.Group("/api/bioskop")
	{
		api.POST("/", controller.CreateBioskop)
		api.GET("/", controller.GetAllBioskop)
		api.GET("/:id", controller.GetBioskopByID)
		api.PUT("/:id", controller.UpdateBioskop)
		api.DELETE("/:id", controller.DeleteBioskop)
	}

	r.Run(":8080")
}
