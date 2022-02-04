package main

import (
	"github.com/indyspyz/web-service-shop/db"
	"github.com/indyspyz/web-service-shop/server"
)

func main() {
	db.Init()
	server.Init()
}
