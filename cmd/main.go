package main

import (
	"github.com/ChemicalZ/GoTemplate/cmd/api"
	"github.com/ChemicalZ/GoTemplate/cmd/internal/db"
	_ "github.com/joho/godotenv/autoload"
)

func main() {
	db := db.NewSQLStorage()
	server := api.NewServer(":8080", db)
	server.Run()

}
