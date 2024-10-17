package main

import (
	"encoding/json"
	"fmt"
	"os"
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
		fmt.Printf("Task updated successfully (ID: %d)\n", 0)
	case "delete":
		fmt.Printf("Task deleted successfully (ID: %d)\n", 0)
	case "mark-in-progress":
		fmt.Printf("Task marked in progress successfully (ID: %d)\n", 0)
	case "mark-done":
		fmt.Printf("Task marked as done successfully (ID: %d)\n", 0)
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
