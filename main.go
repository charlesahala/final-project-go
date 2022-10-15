package main

import (
	"final-project-go/database"
	"final-project-go/routers"
)

func main() {
	database.StartDB()
	r := router.StartApp()
	r.Run(":8080")
}
