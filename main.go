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

//eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6InRlc3RAZ21haWwuY29tIiwiaWQiOjF9.aCPlDs9mlCuO-kLzppoh63XvhbbFOlGD0SLm5UREoGg
