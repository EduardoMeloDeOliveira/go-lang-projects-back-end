package usecase

import "go-api/model"


type ProductUseCase struct {
	// repository aqui jaja
}

func newProductUseCase() ProductUseCase {
	return ProductUseCase{}
}

func (pu * ProductUseCase) GetProducts() ([]model.Product,error){
	return []model.Product{}, nil
}