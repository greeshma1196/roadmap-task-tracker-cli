package main

import (
	"fmt"
	"os"
)

func main() {

	input := os.Args[1]

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

}
