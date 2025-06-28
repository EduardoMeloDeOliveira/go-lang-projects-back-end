package dto

import (
	"go-api/model"
)


type ProductRequestDto struct {
	Name  string `json : "name"`
	Price  float64 `json : "price`
}

//criar uma function que popule um model
// esse pré parametro ali é o reciver basicamente ele ta se alto passando ja como parametro
//geralmente empregado em métods que só le coisas
func(dto ProductRequestDto) ToModel() model.Product{
	return model.Product{
		Name: dto.Name,
		Price: dto.Price,
	}
}


