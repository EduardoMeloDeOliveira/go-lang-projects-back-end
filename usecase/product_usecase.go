package usecase

import (
	"errors"
	"go-api/dto/message"
	"go-api/dto/product"
	"go-api/model"
	"go-api/repository"
)

type ProductUseCase struct {
	productRepository repository.ProductRepository
}

func NewProductUseCase(repository repository.ProductRepository) ProductUseCase {
	return ProductUseCase{
		productRepository: repository,
	}
}

func (pu *ProductUseCase) GetProducts() ([]model.Product, error) {
	return pu.productRepository.GetProducts()
}

func (pu *ProductUseCase) CreateProduct(product dto.ProductRequestDto) (model.Product, error) {

	if !pu.isValidProduct(product) {
		return model.Product{}, errors.New("Produto inválido, o campo name não pode estar vazio e o campo price deve ser maior que zero")
		
	}

	produtModel := product.ToModel()


	return pu.productRepository.CreateProduct(produtModel)

}

func (pu * ProductUseCase) DeleteProduct (id int64) (message.SuccessMessage,error){
	if id < 0 {
		return  message.SuccessMessage{}, errors.New("Id inválido ou não enviado")
	}

	return  pu.productRepository.DeleteProduct(id);
}

func (pu *ProductUseCase) isValidProduct(product dto.ProductRequestDto) bool {
	
	if product.Name == "" || product.Price <= 0 {
		return false
	}

	return true
}
