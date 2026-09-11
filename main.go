package main

import (
	"log"
	"os"

	"manajemen-product/config"
	"manajemen-product/controllers"
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
		log.Fatalf("Gagal koneksi ke database; %v", err)
	}
	defer db.Close()

	productRepo := repositories.NewProductRepository(db)
	productService := services.NewProductService(productRepo)
	productController := controllers.NewProductController(productService)

	r := gin.Default()

	api := r.Group("api/v1")
	{
		api.GET("/products", productController.GetProducts)
		api.POST("/products", productController.CreateProduct)
		api.PUT("/products/:id", productController.UpdateProduct)
	}

	port := os.Getenv("PORT")

	if port == "" {
		port = "8085"
	}

	r.Run(":" + port)

}
