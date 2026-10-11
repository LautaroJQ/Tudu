package database

import (
	"ToDo/models"
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"time"
)

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

	tasks := []models.Task{}
	var modifiedTaskIndex int = -1
	err = json.Unmarshal(fileData, &tasks)

	if err != nil {
		return models.Task{}, errors.New("Could not parse data")
	}

	for i := range tasks {
		if tasks[i].ID == id {
			tasks[i].Title = title
			modifiedTaskIndex = i
			break
		}
	}

	if modifiedTaskIndex == -1 {
		return models.Task{}, errors.New("Could not find task with that id")
	}

	fileData, err = json.MarshalIndent(tasks, "", "\t")

	if err != nil {
		return models.Task{}, err
	}

	err = os.WriteFile(FILENAME, fileData, 0644)

	if err != nil {
		return models.Task{}, err
	}

	return tasks[modifiedTaskIndex], nil

	// FILENAME := os.Getenv("FILENAME")
	// task, err := getTask(id)

	// if err != nil {
	// 	return models.Task{}, err
	// }

	// task.Title = title

	// data, err := os.ReadFile(FILENAME)

	// if err != nil {
	// 	return models.Task{}, err
	// }

	// tasks := []models.Task{}

	// err = json.Unmarshal(data,&tasks)

	// if err != nil {
	// 	return models.Task{}, err
	// }

	// for

	// return task, nil
}
