package main

import (
	"go-users-manager-api/database"
	"go-users-manager-api/internal/routes"
)

func main() {
	database.InitDB()
	routes.Routes()
}
