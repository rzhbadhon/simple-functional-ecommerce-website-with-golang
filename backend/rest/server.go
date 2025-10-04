package rest

import (
	"ecommerce/config"
	"ecommerce/rest/handlers/product"
	"ecommerce/rest/handlers/user"
	"ecommerce/rest/middlewares"
	"fmt"
	"net/http"
	"os"
	"strconv"
)

type Server struct{
	productHandler *product.Handler
	userHandler *user.Handler
}

func NewServer(
	productHandler *product.Handler,
	userHandler *user.Handler,
	) *Server{
	return &Server{
		productHandler: productHandler,
		userHandler: userHandler,
	}
}

func (server *Server) Start(cnf config.Config) {
	manager := middleware.NewManager()

	manager.Use(
		middleware.Preflight,
		middleware.Cors,
		middleware.Logger,
	)

	mux := http.NewServeMux()
	wrappedMux := manager.WrapMux(mux)

	//InitRoutes(mux, manager)

	server.productHandler.RegisterRoutes(mux, manager)
	server.userHandler.RegisterRoutes(mux, manager)

	addr := ":" + strconv.Itoa(cnf.HttpPort)

	fmt.Println("Server running on ", addr)

	err := http.ListenAndServe(addr, wrappedMux)

	if err != nil {
		fmt.Println("Error starting the server", err)
		os.Exit(1)
	}
}
