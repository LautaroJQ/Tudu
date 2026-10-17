package routes

import (
	"ToDo/handlers"
	"net/http"
)

func RegisterTaskRoutes(server *http.ServeMux) {

	server.HandleFunc("POST /tasks", handlers.CreateTaskHandler)
	server.HandleFunc("PATCH /tasks/{id}", handlers.EditTaskHandler)
	server.HandleFunc("GET /tasks", handlers.GetTasksHandler)
	server.HandleFunc("DELETE /tasks/{id}", handlers.DeleteTaskHandler)

}
