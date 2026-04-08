package main

import (
	"fmt"
	"net/http"
	"os"

	log "github.com/sirupsen/logrus"

	"github.com/mukhinfa/golang-advanced/4-order-api/configs"
	"github.com/mukhinfa/golang-advanced/4-order-api/internal/auth"
	"github.com/mukhinfa/golang-advanced/4-order-api/internal/product"
	"github.com/mukhinfa/golang-advanced/4-order-api/pkg/db"
	"github.com/mukhinfa/golang-advanced/4-order-api/pkg/jwt"
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
	jwtSvc := jwt.NewJWT(config.Auth.Secret)

	// Repos
	productRepo := product.NewProductRepository(*db)

	// Services
	productService := product.NewService(productRepo)
	authService := auth.NewAuthService(jwtSvc)

	// Dependencies
	authDeps := auth.NewAuthHandlerDeps(authService)

	// Handlers
	product.NewProductHandler(router, product.ProductHandlerDeps{
		ServiceInterface: productService,
	})
	auth.NewAuthHandler(router, authDeps)

	// Middlewares
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
