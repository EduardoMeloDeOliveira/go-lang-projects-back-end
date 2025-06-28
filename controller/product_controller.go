package controller

import (
	"go-api/model"
	"net/http"

	"github.com/gin-gonic/gin"
)

type productController struct {
	// Use case aqui na pseudo classe
}

// pseudo contrutor da pseudo classe, fora da pseudo classe???

func NewProductController() productController {
	return productController{}
}

func (p *productController) GetProducts(ctx *gin.Context){
	products := []model.Product{
		{
			ID: 1,
			Name: "Mac-book",
			Price: 10000,
		},
		{
			ID: 2,
			Name: "Roubar empréstimo do banco",
			Price: 20000,
		},

	}

			ctx.JSON(http.StatusOK,products)

}
