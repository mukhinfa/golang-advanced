package main

import (
	"fmt"
	"net/http"

	"github.com/mukhinfa/golang-advanced/4-order-api/configs"
	"github.com/mukhinfa/golang-advanced/4-order-api/internal/product"
	"github.com/mukhinfa/golang-advanced/4-order-api/pkg/db"
)

func main() {
	config := configs.LoadConfig()
	db := db.New(&config.DB)

	router := http.NewServeMux()

	productRepo := product.NewProductRepository(*db)
	productService := product.NewService(productRepo)

	product.NewProductHandler(router, product.ProductHandlerDeps{
		ServiceInterface: productService,
	})

	server := &http.Server{
		Addr:    ":8081",
		Handler: router,
	}
	fmt.Println("listen port 8081")

	server.ListenAndServe()

}
