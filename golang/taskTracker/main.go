package main

import (
	"fmt"
	"os"
)

const (
	StatusTodo              = "todo"
	StatusInProgress        = "in-progress"
	StatusDone              = "done"
	version          string = "1.0.0"
)

type Task struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Status      string `json:"status"`
	CreatedAt   string `json:"created_at"`
	UpdatedAt   string `json:"updated_at"`
}

func printUsage() {
	fmt.Println("Usage: task-cli <command> [arguments]")
	fmt.Println("\nCommands:")
	fmt.Println("  add \"<description>\"          Add a new task")
	fmt.Println("  update <id> \"<description>\"  Update a task's description")
	fmt.Println("  delete <id>                Delete a task")
	fmt.Println("  mark-done <id>             Mark a task as 'done'")
	fmt.Println("  mark-in-progress <id>    Mark a task as 'in-progress'")
	fmt.Println("  list [status]              List tasks (all, todo, in-progress, done)")
	fmt.Println("  help                       Show this help message")
}

func main() {
	if len(os.Args) < 2 {
		printUsage()
		os.Exit(1)
	}

	action := os.Args[1]
	args := os.Args[2:]

	fmt.Println("args: ", args)

	fmt.Println("Task Tracker Version:", version)

	switch action {
	case "add":
		fmt.Println("Adding task")

	default:
		fmt.Println("Default action")
	}

}
