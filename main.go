package main

import (
	"encoding/json"
	"fmt"
	"os"
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
		fmt.Printf("Task added successfully (ID: %d)\n", 0)
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
