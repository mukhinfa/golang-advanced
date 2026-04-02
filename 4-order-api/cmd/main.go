package main

import (
	"github.com/mukhinfa/golang-advanced/4-order-api/configs"
	"github.com/mukhinfa/golang-advanced/4-order-api/pkg/db"
)

func main() {
	config := configs.LoadConfig()
	_ = db.New(&config.DB)
}
