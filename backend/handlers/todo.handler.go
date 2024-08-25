package handlers

import (
	"encoding/json"
	"io"
	"net/http"
	"rgt-backend/models"
	"rgt-backend/services"

	"github.com/gorilla/mux"
)

type TodoHandler struct {
	service *services.TodoService
}

func RegisterTodoHandlers(router *mux.Router, service *services.TodoService) {
	handler := &TodoHandler{service: service}

	router.HandleFunc("/todos", handler.GetTodo).Methods(http.MethodGet)
	router.HandleFunc("/todos", handler.CreateTodo).Methods(http.MethodPost)
}

func (h *TodoHandler) GetTodo(res http.ResponseWriter, req *http.Request) {
	todo, err := h.service.GetTodo(req.Context())

	if err != nil {
		http.Error(res, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(res).Encode(todo)
}

func (h *TodoHandler) CreateTodo(res http.ResponseWriter, req *http.Request) {
	body, err := io.ReadAll(req.Body)

	if err != nil {
		http.Error(res, err.Error(), http.StatusInternalServerError)
		return
	}

	var todoBody models.Todos
	defer req.Body.Close()
	err = json.Unmarshal(body, &todoBody)

	if err != nil {
		http.Error(res, err.Error(), http.StatusBadRequest)
		return
	}

	result, err := h.service.CreateTodo(req.Context(), &todoBody)

	if err != nil {
		http.Error(res, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(res).Encode(result)
}
