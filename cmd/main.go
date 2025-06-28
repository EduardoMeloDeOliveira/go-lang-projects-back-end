package main

import (
	"go-api/controller"
	"go-api/db"
	"go-api/repository"
	"go-api/usecase"

	"github.com/gin-gonic/gin"
)

func main() {

	// Server things
	server := gin.Default()

	//Db things
	dbConnection, err := db.ConnectDb()

	if err != nil {
		panic(err)
	}

	//Product things
	ProductRepository := repository.NewProductRepository(dbConnection)
	ProductUseCase := usecase.NewProductUseCase(ProductRepository)
	ProductController := controller.NewProductController(ProductUseCase)
	server.GET("/products", ProductController.GetProducts)
	server.POST("/products", ProductController.CreateProduct)
	server.DELETE("/products/:id", ProductController.DeleteProduct)
	server.GET("/products/:id",ProductController.GetProductById)

	//After de process everthing from project u can start de server setting port
	server.Run(":8080")

}
