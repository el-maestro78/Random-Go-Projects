package task_tracker_cli

import (
	"encoding/json"
	"fmt"
	"os"
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

func printTasks(tasks []Task) {
	if len(tasks) > 0 {
		fmt.Println(tasks)
	} else {
		fmt.Println("There are no tasks stored")
	}
}

func addTask() {

}

func updateTask() {

}

func deleteTask() {

}

func main() {
	var tasks []Task
	tasks, err := loadTasks("tasks.json")
	if err != nil {
		fmt.Println("Empty Tasks")
	}
	args := os.Args
	command := args[1]

	if len(args) < 3 {
		fmt.Println("Usage: go run main.go [command] [age]")
		return
	}

	dateOnly := time.Now().Truncate(24 * time.Hour)
	fmt.Println("Current Date and Time:", dateOnly)
	switch command {
	case add:
		printTasks(tasks)
	case update:
		printTasks(tasks)
	case Delete:
		printTasks(tasks)
	case list:
		printTasks(tasks)
	case mark_in_progress:
		printTasks(tasks)
	case mark_done:
		printTasks(tasks)
	case exit:
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
