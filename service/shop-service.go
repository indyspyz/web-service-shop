package service

import (
	"github.com/indyspyz/web-service-market/entity"
)

type ShopService interface {
	AddShop(entity.Shop) entity.Shop
	FindShop() []entity.Shop
}

type shopService struct {
	shops []entity.Shop
}

func NewShop() ShopService {
	return &shopService{}
}

func (service *shopService) AddShop(shop entity.Shop) entity.Shop {
	service.shops = append(service.shops, shop)
	return shop
}

func (service *shopService) FindShop() []entity.Shop {
	return service.shops
}
