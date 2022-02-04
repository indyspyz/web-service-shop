package controller

import (
	"fmt"

	"github.com/gin-gonic/gin"
	service "github.com/indyspyz/web-service-shop/service"
)

type ProductController struct{}

func (c ProductController) Index(ctx *gin.Context) {
	var s service.ProductService
	p, err := s.GetAll()

	if err != nil {
		ctx.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		ctx.JSON(200, p)
	}
}

func (c ProductController) Create(ctx *gin.Context) {
	var s service.ProductService
	p, err := s.Create(ctx)

	if err != nil {
		ctx.AbortWithStatus(400)
		fmt.Println(err)
	} else {
		ctx.JSON(201, p)
	}
}

func (c ProductController) GetById(ctx *gin.Context) {
	id := ctx.Params.ByName("id")
	var s service.ProductService
	p, err := s.GetByID(id)

	if err != nil {
		ctx.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		ctx.JSON(200, p)
	}
}

func (c ProductController) Update(ctx *gin.Context) {
	id := ctx.Params.ByName("id")
	var s service.ProductService
	p, err := s.UpdateByID(id, ctx)

	if err != nil {
		ctx.AbortWithStatus(400)
		fmt.Println(err)
	} else {
		ctx.JSON(200, p)
	}
}

func (c ProductController) Delete(ctx *gin.Context) {
	id := ctx.Params.ByName("id")
	var s service.ProductService

	if err := s.DeleteByID(id); err != nil {
		ctx.AbortWithStatus(403)
		fmt.Println(err)
	} else {
		ctx.JSON(204, gin.H{"id #" + id: "deleted"})
	}
}
