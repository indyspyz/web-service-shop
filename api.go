package main

import (
	"io"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/indyspyz/web-service-market/controller"
	"github.com/indyspyz/web-service-market/middlewares"
	"github.com/indyspyz/web-service-market/service"
)

var (
	shopService    service.ShopService       = service.NewShop()
	ShopController controller.ShopController = controller.NewShop(shopService)

	productService    service.ProductService       = service.NewProduct()
	ProductController controller.ProductController = controller.NewProduct(productService)
)

func setuoLogOutPut() {
	f, _ := os.Create("gin.log")
	gin.DefaultWriter = io.MultiWriter(f, os.Stdout)
}

func main() {
	setuoLogOutPut()
	api := gin.New()

	api.Use(gin.Recovery(), middlewares.Logger())

	api.GET("/shop/findshop", func(ctx *gin.Context) {
		ctx.JSON(200, ShopController.FindShop())
	})

	api.POST("/shop/addshop", func(ctx *gin.Context) {
		ctx.JSON(200, ShopController.AddShop(ctx))
	})

	api.GET("/products/findproduct", func(ctx *gin.Context) {
		ctx.JSON(200, ProductController.FindProduct())
	})

	api.POST("/products/addproduct", func(ctx *gin.Context) {
		ctx.JSON(200, ProductController.AddProduct(ctx))
	})

	api.Run(":8080")
}
