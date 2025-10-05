package cmd

import (
	"ecommerce/config"
	"ecommerce/rest"
	"ecommerce/rest/handlers/product"
	"ecommerce/rest/handlers/review"
	"ecommerce/rest/handlers/user"
	middleware "ecommerce/rest/middlewares"
)

func Serve() {
	cnf := config.GetConfig()

	middlewares := middleware.NewMiddlewares(cnf)

	productHandler := product.NewHandler(*middlewares)
	userHandler := user.NewHandler()
	reviewHandler := review.NewHandler()

	server := rest.NewServer(
		cnf, 
		productHandler, 
		userHandler, 
		reviewHandler,
	)
	server.Start()
}
