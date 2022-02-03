package service

import "github.com/indyspyz/web-service-market/entity"

type ProductService interface {
	FindProduct() []entity.Product
	AddProduct(entity.Product) entity.Product
}

type productService struct {
	products []entity.Product
}

func NewProduct() ProductService {
	return &productService{}
}

func (service *productService) AddProduct(product entity.Product) entity.Product {
	service.products = append(service.products, product)
	return product
}

func (service *productService) FindProduct() []entity.Product {
	return service.products
}
