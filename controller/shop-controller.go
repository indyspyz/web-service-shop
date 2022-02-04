package controller

import (
	"fmt"

	"github.com/gin-gonic/gin"
	service "github.com/indyspyz/web-service-shop/service"
)

type ShopController struct{}

func (c ShopController) Index(ctx *gin.Context) {
	var s service.ShopService
	p, err := s.GetAll()

	if err != nil {
		ctx.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		ctx.JSON(200, p)
	}
}

func (c ShopController) Create(ctx *gin.Context) {
	var s service.ShopService
	p, err := s.Create(ctx)

	if err != nil {
		ctx.AbortWithStatus(400)
		fmt.Println(err)
	} else {
		ctx.JSON(201, p)
	}
}

func (c ShopController) GetById(ctx *gin.Context) {
	id := ctx.Params.ByName("id")
	var s service.ShopService
	p, err := s.GetByID(id)

	if err != nil {
		ctx.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		ctx.JSON(200, p)
	}
}

func (c ShopController) Update(ctx *gin.Context) {
	id := ctx.Params.ByName("id")
	var s service.ShopService
	p, err := s.UpdateByID(id, ctx)

	if err != nil {
		ctx.AbortWithStatus(400)
		fmt.Println(err)
	} else {
		ctx.JSON(200, p)
	}
}

func (c ShopController) Delete(ctx *gin.Context) {
	id := ctx.Params.ByName("id")
	var s service.ShopService

	if err := s.DeleteByID(id); err != nil {
		ctx.AbortWithStatus(403)
		fmt.Println(err)
	} else {
		ctx.JSON(204, gin.H{"id #" + id: "deleted"})
	}
}

func (c ShopController) GetAllProduct(ctx *gin.Context) {
	var s service.ShopService
	p, err := s.GetAllProduct()

	if err != nil {
		ctx.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		ctx.JSON(200, p)
	}
}
