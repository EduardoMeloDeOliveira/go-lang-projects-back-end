package controller

import (
	"go-api/dto/product"
	"go-api/usecase"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type productController struct {
	productUseCase usecase.ProductUseCase
}

func NewProductController(usecase usecase.ProductUseCase) productController {
	return productController{
		productUseCase: usecase,
	}
}

func (p *productController) GetProducts(ctx *gin.Context) {

	productsList, err := p.productUseCase.GetProducts()

	if err != nil {
		ctx.JSON(500, err)
	}

	ctx.JSON(http.StatusOK, productsList)

}

func (p *productController) CreateProduct(ctx *gin.Context) {

	var productRequestDto dto.ProductRequestDto

	// tenta extrair o json do body batendo com os campos de productRequestDTO
	if err := ctx.ShouldBindJSON(&productRequestDto); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "JSON INVÁLIDO"})
		return
	}

	product, err := p.productUseCase.CreateProduct(productRequestDto)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, product)
}

func (p *productController) DeleteProduct(ctx *gin.Context) {
	idParam := ctx.Param("id")

	if idParam == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Id obrigatório"})
		return
	}

	id, err := strconv.ParseInt(idParam, 10, 64)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Id obrigatório"})
		return
	}

	msg, err := p.productUseCase.DeleteProduct(id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return

	}

	ctx.JSON(http.StatusNoContent, msg)

}

func (p *productController) GetProductById(ctx *gin.Context) {
	idParam := ctx.Param("id")

	if idParam == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Id é obrigatório 1"})
		return
	}

	id, err := strconv.ParseInt(idParam, 10, 64)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Id é obrigatório 2"})
		return
	}
	proudct, err := p.productUseCase.GetProductById(id)

	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, proudct)
}
