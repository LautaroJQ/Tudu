package handlers

import (
	"ToDo/database"
	"ToDo/dtos"
	"ToDo/models"
	"encoding/json"
	"net/http"
	"strconv"
)

func DeleteTaskHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	var dto dtos.DeleteTaskRequest
	id, err := strconv.Atoi(r.PathValue("id"))

	if err != nil {
		http.Error(w, "Cannot parse id", http.StatusBadRequest)
		return
	}

	dto.Id = id

	err = database.DeleteTasks(dto.Id)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = json.NewEncoder(w).Encode("Task succesfully deleted")

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func GetTasksHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")

	// Query params: ?done=&limit=&offset=

	doneParam := r.URL.Query().Get("done")
	var donePtr *bool
	if doneParam != "" {
		v, err := strconv.ParseBool(doneParam)

		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		donePtr = &v
	}

	tasks, err := database.GetTasks(donePtr)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	offsetParam := r.URL.Query().Get("offset")
	var offset int

	if offsetParam != "" {
		offset, err = strconv.Atoi(offsetParam)

		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
	} else {
		offset = 0
	}

	if offset < 0 {
		http.Error(w, "Cannot have an offset less than zero", http.StatusBadRequest)
		return
	}

	if offset >= len(tasks) {
		json.NewEncoder(w).Encode(models.Task{})
	}

	limitParam := r.URL.Query().Get("limit")
	var limit int

	if limitParam != "" {
		limit, err = strconv.Atoi(limitParam)

		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
	} else {
		limit = 5
	}

	if limit <= 0 {
		http.Error(w, "Cannot have an limit equal or less than zero", http.StatusBadRequest)
		return
	}

	end := min(offset+limit, len(tasks))
	filteredTasks := []models.Task{}
	filteredTasks = append(filteredTasks, tasks[offset:end]...)

	err = json.NewEncoder(w).Encode(filteredTasks)

	if err != nil {
		http.Error(w, "Error ocurred during listing", http.StatusInternalServerError)
	}
}

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

	taskRequest := dtos.EditTaskRequest{
		Id:    id,
		Title: r.URL.Query().Get("title"),
	}

	if taskRequest.Title == "" {
		http.Error(w, "Task titlte should not be an empty string", http.StatusBadRequest)
		return
	}

	task, err := database.EditTask(taskRequest.Id, taskRequest.Title)

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
