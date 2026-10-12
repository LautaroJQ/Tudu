package database

import (
	"ToDo/models"
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"time"
)

var ErrTaskNotFound = errors.New("task not found")

func DeleteTasks(id int) error {

	FILENAME := os.Getenv("FILENAME")

	fileData, err := os.ReadFile(FILENAME)

	if err != nil {
		return err
	}
	tasks := []models.Task{}

	err = json.Unmarshal(fileData, &tasks)

	if err != nil {
		return err
	}

	result := []models.Task{}

	for _, task := range tasks {
		if task.ID != id {
			result = append(result, task)
		}
	}

	if len(tasks) == len(result) {
		return ErrTaskNotFound
	}

	fileData, err = json.MarshalIndent(result, "", "\t")

	if err != nil {
		return err
	}

	err = os.WriteFile(FILENAME, fileData, 0644)

	if err != nil {
		return err
	}

	return nil
}

func GetTasks(done *bool) ([]models.Task, error) {
	FILENAME := os.Getenv("FILENAME")

	fileData, err := os.ReadFile(FILENAME)

	if err != nil {
		return []models.Task{}, err
	}
	var tasks []models.Task
	err = json.Unmarshal(fileData, &tasks)

	if err != nil {
		return []models.Task{}, err
	}

	var filteredTasks []models.Task

	for _, task := range tasks {
		if done == nil || task.Done == *done {
			filteredTasks = append(filteredTasks, task)
		}
	}

	return filteredTasks, nil
}

func NewTask(title string) (models.Task, error) {

	fileName := os.Getenv("FILENAME")

	tasks := []models.Task{}

	fileData, err := os.ReadFile(fileName)

	if err == nil {
		err = json.Unmarshal(fileData, &tasks)

		if err != nil {
			return models.Task{}, errors.New("Cannot decode data from file")
		}
	} else if !errors.Is(err, os.ErrNotExist) {
		return models.Task{}, errors.New("Cannot read file")
	}

	id := 1
	actualTime := time.Now()

	task := models.Task{
		ID:        id,
		Title:     title,
		Done:      false,
		CreatedAt: actualTime,
	}

	tasks = append(tasks, task)

	data, err := json.MarshalIndent(tasks, "", "\t")

	if err != nil {
		return models.Task{}, errors.New("Cannot encode data")
	}

	err = os.WriteFile(fileName, data, 0644)

	if err != nil {
		return models.Task{}, errors.New("Cannot write data")
	}
	fmt.Print(1123123)
	return task, nil
}

func EditTask(id int, title string) (models.Task, error) {
	FILENAME := os.Getenv("FILENAME")

	fileData, err := os.ReadFile(FILENAME)

	if err != nil {
		return models.Task{}, errors.New("Cannot open file")
	}
	//Lee el archivo de guardado
	tasks := []models.Task{}
	var modifiedTaskIndex int = -1
	err = json.Unmarshal(fileData, &tasks)

	if err != nil {
		return models.Task{}, errors.New("Could not parse data")
	}
	//Busca coincidencia por Id
	for i := range tasks {
		if tasks[i].ID == id {
			tasks[i].Title = title
			modifiedTaskIndex = i
			break
		}
	}
	//Si no hay coincidencia devuelve error
	if modifiedTaskIndex == -1 {
		return models.Task{}, errors.New("Could not find task with that id")
	}

	//Si hay coincidencia edita y guarda
	fileData, err = json.MarshalIndent(tasks, "", "\t")

	if err != nil {
		return models.Task{}, err
	}

	err = os.WriteFile(FILENAME, fileData, 0644)

	if err != nil {
		return models.Task{}, err
	}

	return tasks[modifiedTaskIndex], nil
}
