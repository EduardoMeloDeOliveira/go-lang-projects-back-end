package controller

import (
	"go-api/model"
	"go-api/usecase"
	"net/http"

	"github.com/gin-gonic/gin"
)

type productController struct {
	productUseCase usecase.ProductUseCase
}

// pseudo contrutor da pseudo classe, fora da pseudo classe???

func NewProductController(usecase usecase.ProductUseCase) productController {
	return productController{
		productUseCase:  usecase,
	}
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
