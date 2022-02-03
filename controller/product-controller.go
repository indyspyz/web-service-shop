package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/indyspyz/web-service-market/entity"
	"github.com/indyspyz/web-service-market/service"
)

type ProductController interface {
	FindProduct() []entity.Product
	AddProduct(ctx *gin.Context) entity.Product
}

type productController struct {
	service service.ProductService
}

func NewProduct(service service.ProductService) ProductController {
	return &productController{
		service: service,
	}
}

func (c *productController) FindProduct() []entity.Product {
	return c.service.FindProduct()
}

func (c *productController) AddProduct(ctx *gin.Context) entity.Product {
	var product entity.Product
	ctx.BindJSON(&product)
	c.service.AddProduct(product)
	return product
}
