package main

import (
	"github.com/keiya01/myblog/database"
	"github.com/keiya01/myblog/http"
	"github.com/keiya01/myblog/migration"
)

func main() {
	db := database.NewHandler()
	migration.Set(db.DB)
	s := http.NewServer()
	s.Route()
	s.Start(":8686")
}
