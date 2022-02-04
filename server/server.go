package server

import (
	"io"
	"os"

	"github.com/gin-gonic/gin"
	ctrl "github.com/indyspyz/web-service-shop/controller"
	"github.com/indyspyz/web-service-shop/middlewares"
)

func Init() {
	r := Router()
	r.Run()
}

func setupLogOutPut() {
	f, _ := os.Create("gin.log")
	gin.DefaultWriter = io.MultiWriter(f, os.Stdout)
}

func Router() *gin.Engine {
	setupLogOutPut()
	g := gin.New()
	g.Use(gin.Recovery(), middlewares.Logger())

	shopCtrl := ctrl.ShopController{}
	productCtrl := ctrl.ProductController{}

	g.GET("/shop", shopCtrl.Index)
	g.GET("/shop/:id", shopCtrl.GetById)
	g.POST("/shop", shopCtrl.Create)
	g.PUT("/shop/:id", shopCtrl.Update)
	g.DELETE("/shop/:id", shopCtrl.Delete)

	g.GET("/product", productCtrl.Index)
	g.GET("/product/:id", productCtrl.GetById)
	g.POST("/product", productCtrl.Create)
	g.PUT("/product/:id", productCtrl.Update)
	g.DELETE("/product/:id", productCtrl.Delete)

	g.GET("/viewproduct", shopCtrl.GetAllProduct)

	return g
}
