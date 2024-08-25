package main

import (
	"net/http"
	"rgt-backend/db"
	"rgt-backend/handlers"
	"rgt-backend/services"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()

	client := db.ConnectDB()
	todoService := services.NewTodoService(client)

	handlers.RegisterTodoHandlers(router, todoService)

	http.ListenAndServe(":8080", router)
}
