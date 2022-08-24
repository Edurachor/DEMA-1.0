package main

import (
	"apidema/api"
	"apidema/api/database"
)

func main() {
	database.ConnectDatabase()
	database.AutoMigration()
	api.Run()
}
