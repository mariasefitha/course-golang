package main

import (
	"course-golang/config"
	"course-golang/controller"
	"course-golang/model"
	"os"

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

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080" // fallback default
	}
	r.Run(":" + port)

}
