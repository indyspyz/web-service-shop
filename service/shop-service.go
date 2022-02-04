package service

import (
	"github.com/gin-gonic/gin"

	"github.com/indyspyz/web-service-shop/db"
	"github.com/indyspyz/web-service-shop/entity"
)

type ShopService struct{}

func (s ShopService) GetAll() ([]entity.Shop, error) {
	db := db.GetDB()
	var res []entity.Shop

	if err := db.Find(&res).Error; err != nil {
		return nil, err
	}

	return res, nil
}

func (s ShopService) Create(ctx *gin.Context) (entity.Shop, error) {
	db := db.GetDB()
	var res entity.Shop

	if err := ctx.BindJSON(&res); err != nil {
		return res, err
	}

	if err := db.Create(&res).Error; err != nil {
		return res, err
	}

	return res, nil
}

func (s ShopService) GetByID(id string) (entity.Shop, error) {
	db := db.GetDB()
	var res entity.Shop

	if err := db.Where("id = ?", id).First(&res).Error; err != nil {
		return res, err
	}

	return res, nil
}

func (s ShopService) UpdateByID(id string, c *gin.Context) (entity.Shop, error) {
	db := db.GetDB()
	var res entity.Shop

	if err := db.Where("id = ?", id).First(&res).Error; err != nil {
		return res, err
	}

	if err := c.BindJSON(&res); err != nil {
		return res, err
	}

	db.Save(&res)

	return res, nil
}

func (s ShopService) DeleteByID(id string) error {
	db := db.GetDB()
	var res entity.Shop

	if err := db.Where("id = ?", id).Delete(&res).Error; err != nil {
		return err
	}

	return nil
}

func (s ShopService) GetAllProduct() ([]entity.ShopList, error) {
	db := db.GetDB()
	var res []entity.ShopList

	if err := db.Preload("ProductList").Find(&res).Error; err != nil {
		return nil, err
	}

	return res, nil
}
