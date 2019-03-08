package database

import (
	"github.com/jinzhu/gorm"
	_ "github.com/mattn/go-sqlite3"
)

type Handler struct {
	*gorm.DB
}

func NewHandler() *Handler {
	dbName := "myblog.sqlite3"

	db, err := gorm.Open("sqlite3", dbName)
	if err != nil {
		panic(err)
	}

	handler := Handler{db}

	return &handler
}
