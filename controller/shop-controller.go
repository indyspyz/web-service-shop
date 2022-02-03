package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/indyspyz/web-service-market/entity"
	"github.com/indyspyz/web-service-market/service"
)

type ShopController interface {
	FindShop() []entity.Shop
	AddShop(ctx *gin.Context) entity.Shop
}

type shopController struct {
	service service.ShopService
}

func NewShop(service service.ShopService) ShopController {
	return &shopController{
		service: service,
	}
}

func (c *shopController) FindShop() []entity.Shop {
	return c.service.FindShop()
}

func (c *shopController) AddShop(ctx *gin.Context) entity.Shop {
	var shop entity.Shop
	ctx.BindJSON(&shop)
	c.service.AddShop(shop)
	return shop
}
