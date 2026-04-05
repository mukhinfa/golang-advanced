package db

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"github.com/mukhinfa/golang-advanced/4-order-api/configs"
)

type Db struct {
	*gorm.DB
}

func New(c *configs.DbConfig) *Db {
	db, err := gorm.Open(postgres.Open(c.DSN), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	return &Db{DB: db}
}
