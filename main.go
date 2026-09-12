package main

import (
	"log"
	"os"

	"manajemen-product/config"
	"manajemen-product/controllers"
	"manajemen-product/models"
	"manajemen-product/repositories"
	"manajemen-product/services"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Println("File .env tidak ditemukan, menggunakan environment OS")
	}

	// koneksi database
	db, err := config.ConnectDB()
	if err != nil {
		log.Fatalf("Gagal koneksi ke database: %v", err)
	}

	err = db.AutoMigrate(&models.Product{})

	if err != nil {
		log.Fatalf("Gagal melakukan migrasi: %v", err)
	}

	productRepo := repositories.NewProductRepository(db)
	productService := services.NewProductService(productRepo)
	productController := controllers.NewProductController(productService)

	r := gin.Default()

	api := r.Group("api/v1")
	{
		api.GET("/products", productController.GetProducts)
		api.GET("/products/:id", productController.FindById)
		api.POST("/products", productController.CreateProduct)
		api.PUT("/products/:id", productController.UpdateProduct)
		api.DELETE("products/:id", productController.DeleteProduct)
	}

	port := os.Getenv("PORT")

	if port == "" {
		port = "8085"
	}

	r.Run(":" + port)

}
