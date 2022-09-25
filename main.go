package main

import (
	"middleware/database"
	"middleware/router"
)

func main() {
	database.StartDB()
	r := router.StartApp()
	r.Run(":8080")
}
