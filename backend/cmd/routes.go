package cmd

import (
	"ecommerce/handlers"
	"ecommerce/middleware"
	"net/http"
)

func initRoutes(mux *http.ServeMux, manager *middleware.Manager) {
	// mux.Handle(
	// 	"GET /rahim",
	// 	manager.With(
	// 		http.HandlerFunc(handlers.GetProducts),
	// 	),
	// )

	// mux.Handle("GET /karim", middleware.Test(middleware.Logger(http.HandlerFunc(handlers.GetProducts))))

	mux.Handle(
		"GET /products",
		manager.With(
			http.HandlerFunc(handlers.GetProducts),
		),
	)
	mux.Handle(
		"GET /products/{productId}",
		manager.With(
			http.HandlerFunc(handlers.GetProductByID),
		),
	)

	// mux.HandleFunc("/create-product", createProducts)   // previous 1.22 version
	// mux.Handle("POST /create-product", http.HandlerFunc(createProducts))

	mux.Handle(
		"POST /products",
		manager.With(
			http.HandlerFunc(handlers.CreateProducts),
		),
	)

}
