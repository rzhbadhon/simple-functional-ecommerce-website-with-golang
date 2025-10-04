package cmd

import (
	"ecommerce/config"
	"ecommerce/rest"
	"ecommerce/rest/handlers/product"
	"ecommerce/rest/handlers/user"
)

func Serve() {
	cnf := config.GetConfig()

	productHandler := product.NewHandler()
	userHandler := user.NewHandler()

	server := rest.NewServer(productHandler, userHandler)
	server.Start(cnf)
}
