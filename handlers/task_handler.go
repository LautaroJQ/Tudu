package handlers

import (
	"ToDo/database"
	"ToDo/dtos"
	"encoding/json"
	"net/http"
	"strconv"
)

func EditTaskHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")

	id, err := strconv.Atoi(r.PathValue("id"))

	if err != nil {
		http.Error(w, "Cannot parse id", http.StatusBadRequest)
		return
	}

	if id < 0 {
		http.Error(w, "Id cannot be lower than zero", http.StatusBadRequest)
		return
	}

	taskRequest := dtos.EditTaskRequest{Title: r.URL.Query().Get("title")}

	if taskRequest.Title == "" {
		http.Error(w, "Task titlte should not be an empty string", http.StatusBadRequest)
		return
	}

	task, err := database.EditTask(id, taskRequest.Title)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(task)
}

func CreateTaskHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")

	taskRequest := dtos.CreateTaskRequest{}

	err := json.NewDecoder(r.Body).Decode(&taskRequest)

	if err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	if taskRequest.Title == "" {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	task, err := database.NewTask(taskRequest.Title)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(201)
	err = json.NewEncoder(w).Encode(task)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
