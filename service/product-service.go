package service

import (
	"github.com/gin-gonic/gin"

	"github.com/indyspyz/web-service-shop/db"
	"github.com/indyspyz/web-service-shop/entity"
)

type ProductService struct{}

func (s ProductService) GetAll() ([]entity.Product, error) {
	db := db.GetDB()
	var res []entity.Product

	if err := db.Find(&res).Error; err != nil {
		return nil, err
	}

	return res, nil
}

func (s ProductService) Create(ctx *gin.Context) (entity.Product, error) {
	db := db.GetDB()
	var res entity.Product

	if err := ctx.BindJSON(&res); err != nil {
		return res, err
	}

	if err := db.Create(&res).Error; err != nil {
		return res, err
	}

	return res, nil
}

func (s ProductService) GetByID(id string) (entity.Product, error) {
	db := db.GetDB()
	var res entity.Product

	if err := db.Where("id = ?", id).First(&res).Error; err != nil {
		return res, err
	}

	return res, nil
}

func (s ProductService) UpdateByID(id string, c *gin.Context) (entity.Product, error) {
	db := db.GetDB()
	var res entity.Product

	if err := db.Where("id = ?", id).First(&res).Error; err != nil {
		return res, err
	}

	if err := c.BindJSON(&res); err != nil {
		return res, err
	}

	db.Save(&res)

	return res, nil
}

func (s ProductService) DeleteByID(id string) error {
	db := db.GetDB()
	var res entity.Product

	if err := db.Where("id = ?", id).Delete(&res).Error; err != nil {
		return err
	}

	return nil
}
