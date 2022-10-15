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

//$2a$08$XAkMaV/zqDlj1hDLC0BEZevVsl.ImLEQIHqVCWOMf0TDV5UFW.lCW
//eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6ImNoc0BnbWFpbC5jb20iLCJwYXNzd29yZCI6IiQyYSQwOCRYQWtNYVYvenFEbGoxaERMQzBCRVpldlZzbC5JbUxFUUlIcVZDV09NZjBURFY1VUZXLmxDVyJ9.OBwVAYLO8Do-3pzNkdR8s3NUn-G3rcnCRheYqD-M2MM