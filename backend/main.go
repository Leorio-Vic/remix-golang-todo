package main

import (
	"net/http"
	"rgt-backend/db"
)

func main() {
	db.InitDB()
	http.ListenAndServe(":8080", nil)
}
