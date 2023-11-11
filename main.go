package main

import (
	"flag"
	"fmt"
	"os"
)

const tasksFilename = "tasks.json"

func main() {
	taskManager := TaskManager{}

	err := taskManager.loadTasksFromFile(tasksFilename)

	if err != nil {
		fmt.Println("Error loading tasks from file:", err)
	}

	addFlag := flag.String("add", "", "Add a new task")
	removeFlag := flag.Int("remove", 0, "Remove a task by ID")
	listFlag := flag.Bool("list", false, "List all tasks")

	flag.Parse()

	switch {
	case *addFlag != "":
		taskManager.AddTask(*addFlag)
		fmt.Println("Task added successfully")
		err := taskManager.saveTasksToFile(tasksFilename)
		if err != nil {
			fmt.Println("Error saving tasks:", err)
		}
	case *removeFlag != 0:
		taskManager.RemoveTask(*removeFlag)
		fmt.Println("Task removed successfully")
		err := taskManager.saveTasksToFile(tasksFilename)
		if err != nil {
			fmt.Println("Error saving tasks:", err)
		}
	case *listFlag:
		tasks := taskManager.ListTasks()
		fmt.Println("Tasks:")
		for _, task := range tasks {
			fmt.Printf("%d. %s (created at %s)\n", task.ID, task.Title, task.CreatedAt.Format("2006-01-02 15:04:05"))
		}
	default:
		fmt.Println("Usage: task-tracker [options]")
		flag.PrintDefaults()
		os.Exit(1)
	}
}
