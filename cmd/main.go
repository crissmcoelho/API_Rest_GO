package main

import (
	"API_Rest_GO/controller"
	"API_Rest_GO/db"
	"API_Rest_GO/repository"
	"API_Rest_GO/usecase"

	"github.com/gin-gonic/gin"
)

func main() {

	dbConnection, err:= db.ConnectDB()
	if err != nil {
		panic(err)
	}	

	server := gin.Default()

	ProductRepository:= repository.NewProductRespository(dbConnection)

	ProductUseCase:= usecase.NewProductUseCase(ProductRepository)

	ProductController := controller.NewProducterController(ProductUseCase)

	server.GET("/ping", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "pong",
		})
	})

	server.GET("/products", ProductController.GetProducts)

	server.Run(":8000")

}
