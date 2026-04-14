package main

import (
	"github.com/mukhinfa/golang-advanced/4-order-api/configs"
	"github.com/mukhinfa/golang-advanced/4-order-api/internal/order"
	"github.com/mukhinfa/golang-advanced/4-order-api/internal/product"
	"github.com/mukhinfa/golang-advanced/4-order-api/internal/user"
	"github.com/mukhinfa/golang-advanced/4-order-api/pkg/db"
)

func main() {
	conf := configs.LoadConfig()
	db := db.New(&conf.DB)

	if err := db.AutoMigrate(
		&product.Product{},
		&user.User{},
		&order.Order{},
	); err != nil {
		panic(err)
	}
}
