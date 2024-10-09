package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strconv"
	"time"
)

const (
	todo = iota
	inProgress
	done
)

type Task struct {
	id          int       `json:"id"`          //A unique identifier for the task
	description string    `json:"description"` // A short description of the task
	status      int       `json:"status"`      //The status of the task (todo, in-progress, done)
	createdAt   time.Time `json:"createdAt"`   //The date and time when the task was created
	updatedAt   time.Time `json:"updatedAt"`   //The date and time when the task was last updated
}

func saveTasks(tasks []Task, filename string) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()
	encoder := json.NewEncoder(file)
	encoder.SetIndent("", " ") //format for readability
	return encoder.Encode(tasks)
}

func loadTasks(filename string) ([]Task, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var tasks []Task
	decoder := json.NewDecoder(file)
	err = decoder.Decode(&tasks)
	return tasks, err
}

func currentTaskID(tasks []Task) int {
	if len(tasks) == 0 {
		return 0
	}
	maxId := 0
	for i := range tasks {
		if tasks[i].id > maxId {
			maxId = tasks[i].id
		}
	}
	return maxId
}

func printTasks(tasks []Task) {
	if len(tasks) > 0 {
		fmt.Println(tasks)
	} else {
		fmt.Println("There are no tasks stored")
	}
}

func getStatus(status string) int {
	switch status {
	case "todo":
		return 0
	case "in-progress":
		return 1
	case "done":
		return 2
	default:
		return -1
	}
}

func printTasksByStatus(tasks []Task, status int) {
	if len(tasks) == 0 {
		fmt.Println("There are no tasks stored")
	} else {
		var filtered []Task
		for _, task := range tasks {
			if task.status == status {
				filtered = append(filtered, task)
			}
		}
		fmt.Println(filtered)
	}
}

func addTask(tasks []Task, description string, tasksId *int) {
	task := Task{
		id:          *tasksId,
		description: description,
		status:      0,
		createdAt:   time.Now().Truncate(24 * time.Hour), //time.Time{},
		updatedAt:   time.Now().Truncate(24 * time.Hour),
	}
	tasks = append(tasks, task)
	*tasksId++
}

func updateTask(tasks []Task, taskId int, description string) {
	var changed = false
	for i := range tasks {
		if tasks[i].id == taskId {
			tasks[i].description = description
			changed = true
			break
		}
	}
	if !changed {
		fmt.Println("No task found")
	} else {
		fmt.Println("Task updated")
	}
}

func deleteTask(tasks []Task, taskId int) []Task {
	var changed = false
	for i := range tasks {
		if tasks[i].id == taskId {
			tasks = append(tasks[:i], tasks[i+1:]...)
			changed = true
			break
		}
	}
	if !changed {
		fmt.Println("No task found")
	} else {
		fmt.Println("Task updated")
	}
	return tasks
}

func markInProgress(tasks []Task, taskId int) {
	var changed = false
	for i := range tasks {
		if tasks[i].id == taskId {
			tasks[i].status = 1
			changed = true
			break
		}
	}
	if !changed {
		fmt.Println("No task found")
	} else {
		fmt.Println("Task updated")
	}
}

func markDone(tasks []Task, taskId int) {
	var changed = false
	for i := range tasks {
		if tasks[i].id == taskId {
			tasks[i].status = 2
			changed = true
			break
		}
	}
	if !changed {
		fmt.Println("No task found")
	} else {
		fmt.Println("Task updated")
	}
}

func main() {
	var tasks []Task
	var tasksId = currentTaskID(tasks)
	tasks, err := loadTasks("tasks.json")
	if err != nil {
		fmt.Println("Empty Tasks")
	}

	for { //while True
		args := os.Args
		if len(args) < 3 {
			fmt.Println("Usage: go run main.go [command] [options]")
			return
		}
		command := args[1]

		switch command {
		case "add":
			description := args[2]
			addTask(tasks, description, &tasksId)
		case "update":
			taskID, err := strconv.Atoi(args[2])
			if err != nil {
				fmt.Println("Invalid ID type, try again")
			}
			description := args[3]
			updateTask(tasks, taskID, description)
		case "delete":
			taskID, err := strconv.Atoi(args[2])
			if err != nil {
				fmt.Println("Invalid ID type, try again")
			}
			tasks = deleteTask(tasks, taskID)
		case "list":
			if len(args) == 1 {
				printTasks(tasks)
			} else {
				status := getStatus(args[2])
				if status < 0 {
					fmt.Println("Invalid status type, try again")
				}
				printTasksByStatus(tasks, status)
			}
		case "mark_in_progress":
			taskID, err := strconv.Atoi(args[2])
			if err != nil {
				fmt.Println("Invalid ID type, try again")
			}
			markInProgress(tasks, taskID)
		case "mark_done":
			taskID, err := strconv.Atoi(args[2])
			if err != nil {
				fmt.Println("Invalid ID type, try again")
			}
			markDone(tasks, taskID)
		case "exit":
			err := saveTasks(tasks, "tasks.json")
			if err != nil {
				fmt.Println("Saved Successfully")
				os.Exit(0)
			} else {
				fmt.Println("Error saving tasks in json", err)
				os.Exit(1)
			}
		default:
			fmt.Println("Unknown command:", command)
			fmt.Println("Unknown command:", command)
		}
	}

}
