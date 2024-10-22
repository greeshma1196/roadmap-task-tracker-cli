package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strconv"
	"time"
)

const fileName = "data.json"

type status string

const (
	StatusToDo       status = "todo"
	StatusInProgress status = "in-progress"
	StatusDone       status = "done"
)

type task struct {
	ID          int    `json:"ID"`
	Description string `json:"Description"`
	Status      status `json:"Status"`
	CreatedAt   int64  `json:"CreatedAt"`
	UpdatedAt   int64  `json:"UpdatedAt"`
}

func addTask(data []task, desc string) ([]task, int, error) {

	id := len(data) + 1

	// set struct task
	t := task{ID: id, Description: desc, Status: StatusToDo, CreatedAt: time.Now().Unix(), UpdatedAt: time.Now().Unix()}

	// append task to unmarshalled data
	data = append(data, t)

	return data, id, nil
}

func updateTask(data []task, id int, desc string) ([]task, error) {

	for i := range data {
		if data[i].ID == id {
			data[i].Description = desc
			data[i].UpdatedAt = time.Now().Unix()
			return data, nil
		}
	}

	return data, fmt.Errorf("task not present")
}

func deleteTask(data []task, id int) ([]task, error) {

	var dataUpdated []task
	isTaskPresent := false

	//delete task
	for i := range data {
		if data[i].ID == id {
			isTaskPresent = true
		} else {
			dataUpdated = append(dataUpdated, data[i])
		}
	}

	if !isTaskPresent {
		return dataUpdated, fmt.Errorf("task not present")
	}

	return dataUpdated, nil
}

func markInProgressTask(data []task, id int) ([]task, error) {

	// mark task in progress
	for i := range data {
		if data[i].ID == id {
			data[i].Status = StatusInProgress
			data[i].UpdatedAt = time.Now().Unix()
			return data, nil
		}
	}

	return data, fmt.Errorf("task not present")
}

func markDoneTask(data []task, id int) ([]task, error) {

	// mark task in progress
	for i := range data { // why does this only work with index?
		if data[i].ID == id {
			data[i].Status = StatusDone
			data[i].UpdatedAt = time.Now().Unix()
			return data, nil
		}
	}

	return data, fmt.Errorf("task not present")
}

func main() {

	input := os.Args[1]

	// create the data file if it does not exist
	_, err := os.Stat(fileName)
	if os.IsNotExist(err) {
		_, err := os.Create(fileName)
		if err != nil {
			panic(err)
		}
	}

	// read existing tasks if any
	file, err := os.ReadFile(fileName)

	if err != nil {
		panic(err)
	}

	data := []task{}

	if len(file) > 0 {
		err = json.Unmarshal(file, &data)
		if err != nil {
			panic(err)
		}
	}

	switch input {
	case "add":
		if len(os.Args) != 3 || os.Args[2] == "" {
			panic(fmt.Errorf("missing description"))
		}

		var id int

		data, id, err = addTask(data, os.Args[2])
		if err != nil {
			panic(err)
		}

		fmt.Printf("Task added successfully (ID: %d)\n", id)
	case "update":
		if len(os.Args) != 4 || os.Args[3] == "" {
			panic(fmt.Errorf("missing task id or description"))
		}

		id, err := strconv.Atoi(os.Args[2])
		if err != nil {
			panic(err)
		}
		if id < 1 {
			panic(fmt.Errorf("invalid id"))
		}

		if len(data) == 0 {
			panic(fmt.Errorf("no tasks present, please add"))
		}

		data, err = updateTask(data, id, os.Args[3])

		if err != nil {
			panic(err)
		}

		fmt.Printf("Task updated successfully (ID: %d)\n", id)
	case "delete":
		if len(os.Args) != 3 {
			panic(fmt.Errorf("missing task id"))
		}

		id, err := strconv.Atoi(os.Args[2])
		if err != nil {
			panic(err)
		}

		if id < 1 {
			panic(fmt.Errorf("invalid id"))
		}

		if len(data) == 0 {
			panic(fmt.Errorf("no tasks present, please add"))
		}

		data, err = deleteTask(data, id)
		if err != nil {
			panic(err)
		}

		fmt.Printf("Task deleted successfully (ID: %d)\n", id)
	case "mark-in-progress":
		if len(os.Args) != 3 {
			panic(fmt.Errorf("missing task id"))
		}

		id, err := strconv.Atoi(os.Args[2])
		if err != nil {
			panic(err)
		}

		if id < 1 {
			panic(fmt.Errorf("invalid id"))
		}

		if len(data) == 0 {
			panic(fmt.Errorf("no tasks present, please add"))
		}

		data, err = markInProgressTask(data, id)
		if err != nil {
			panic(err)
		}

		fmt.Printf("Task marked in progress successfully (ID: %d)\n", id)
	case "mark-done":
		if len(os.Args) != 3 {
			panic(fmt.Errorf("missing task id"))
		}

		id, err := strconv.Atoi(os.Args[2])
		if err != nil {
			panic(err)
		}

		if id < 1 {
			panic(fmt.Errorf("invalid id"))
		}

		if len(data) == 0 {
			panic(fmt.Errorf("no tasks present, please add"))
		}

		data, err = markDoneTask(data, id)
		if err != nil {
			panic(err)
		}

		fmt.Printf("Task marked in progress successfully (ID: %d)\n", id)
	case "list":
		fmt.Printf("No Tasks\n")
	}

	b, err := json.MarshalIndent(data, "", " ")
	if err != nil {
		panic(err)
	}

	// write tasks to file
	err = os.WriteFile(fileName, b, os.ModePerm)
	if err != nil {
		panic(err)
	}

}
