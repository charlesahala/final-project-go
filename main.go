package main

import (
	"final-project-go/database"
	"final-project-go/routers"
)

func main() {
	database.StartDB()
	r := router.StartService()
	r.Run(":8080")
}
