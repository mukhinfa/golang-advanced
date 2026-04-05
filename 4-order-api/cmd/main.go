package main

import (
	"fmt"
	"net/http"
	"os"

	log "github.com/sirupsen/logrus"

	"github.com/mukhinfa/golang-advanced/4-order-api/configs"
	"github.com/mukhinfa/golang-advanced/4-order-api/internal/product"
	"github.com/mukhinfa/golang-advanced/4-order-api/pkg/db"
	"github.com/mukhinfa/golang-advanced/4-order-api/pkg/middleware"
)

func init() {
	log.SetFormatter(&log.JSONFormatter{})

	log.SetOutput(os.Stdout)

	log.SetLevel(log.WarnLevel)
}

func main() {
	config := configs.LoadConfig()
	db := db.New(&config.DB)

	router := http.NewServeMux()

	productRepo := product.NewProductRepository(*db)
	productService := product.NewService(productRepo)

	product.NewProductHandler(router, product.ProductHandlerDeps{
		ServiceInterface: productService,
	})

	stack := middleware.Chain(
		middleware.CORS,
		middleware.Logging,
	)

	server := &http.Server{
		Addr:    ":8081",
		Handler: stack(router),
	}
	fmt.Println("listen port 8081")

	server.ListenAndServe()

}
