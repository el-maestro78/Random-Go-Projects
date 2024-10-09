package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

const (
	todo = iota
	inProgress
	done
)

type Task struct {
	Id          int       `json:"id"`          //A unique identifier for the task
	Description string    `json:"description"` // A short description of the task
	Status      int       `json:"status"`      //The status of the task (todo, in-progress, done)
	CreatedAt   time.Time `json:"createdAt"`   //The date and time when the task was created
	UpdatedAt   time.Time `json:"updatedAt"`   //The date and time when the task was last updated
}

func saveTasks(tasks []Task, filename string) error {
	var file *os.File
	//var err error
	if _, err := os.Stat(filename); err == nil {
		file, err = os.OpenFile(filename, os.O_WRONLY, 0644) //O_APPEND|os.O_WRONLY
		if err != nil {
			return err
		}
	} else if os.IsNotExist(err) {
		file, err = os.Create(filename)
		if err != nil {
			return err
		}
	} else {
		return err
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			fmt.Println(err)
		}
	}(file)
	encoder := json.NewEncoder(file)
	encoder.SetIndent("", " ") //format for readability
	return encoder.Encode(tasks)
}

func loadTasks(filename string) ([]Task, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			fmt.Println(err)
		}
	}(file)

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
		if tasks[i].Id > maxId {
			maxId = tasks[i].Id
		}
	}
	return maxId
}

func prettyPrint(tasks []Task) error {
	encoder := json.NewEncoder(os.Stdout)
	encoder.SetIndent("", "  ")
	if err := encoder.Encode(tasks); err != nil {
		return err
	}
	return nil
}

func printTasks(tasks []Task) {
	if len(tasks) > 0 {
		err := prettyPrint(tasks)
		if err != nil {
			fmt.Println("Error:", err)
		}
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
	if len(tasks) > 0 {
		var filtered []Task
		for _, task := range tasks {
			if task.Status == status {
				filtered = append(filtered, task)
			}
		}
		if len(filtered) > 0 {
			err := prettyPrint(tasks)
			if err != nil {
				fmt.Println("Error:", err)
			}
		}
		fmt.Println("There are no tasks stored")
	} else {
		fmt.Printf("There are no tasks with status %d\n", status)
	}
}

func addTask(tasks []Task, description string, tasksId *int) []Task {
	task := Task{
		Id:          *tasksId,
		Description: description,
		Status:      todo,
		CreatedAt:   time.Time{},
		UpdatedAt:   time.Time{}, //time.Now().Truncate(24 * time.Hour),
	}
	tasks = append(tasks, task)
	*tasksId++
	fmt.Printf("%s Added successfully\n", description)
	return tasks
}

func updateTask(tasks []Task, taskId int, description string) []Task {
	var changed = false
	for i := range tasks {
		if tasks[i].Id == taskId {
			tasks[i].Description = description
			tasks[i].UpdatedAt = time.Time{}
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

func deleteTask(tasks []Task, taskId int) []Task {
	var changed = false
	for i := range tasks {
		if tasks[i].Id == taskId {
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

func markInProgress(tasks []Task, taskId int) []Task {
	var changed = false
	for i := range tasks {
		if tasks[i].Id == taskId {
			tasks[i].Status = inProgress
			tasks[i].UpdatedAt = time.Time{}
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

func markDone(tasks []Task, taskId int) []Task {
	var changed = false
	for i := range tasks {
		if tasks[i].Id == taskId {
			tasks[i].Status = done
			tasks[i].UpdatedAt = time.Time{}
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

func main() {
	var tasks []Task
	tasks, err := loadTasks("tasks.json")
	var tasksId = currentTaskID(tasks)
	tasksId++
	if err != nil {
		//fmt.Println("Empty Tasks")
	}
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Enter command new command. Type help to see all commands")
	for { //while True
		//fmt.Print("Enter command: ")
		fmt.Print("task-cli ")
		if !scanner.Scan() {
			break
		}
		line := scanner.Text()
		args := strings.Fields(line) // Split input into arguments
		if len(args) < 1 {
			fmt.Print("Error. Try again")
			continue
		}
		command := args[0]

		switch command {
		case "add":
			if len(args) < 2 {
				fmt.Println("Usage: add [description]")
				continue
			}
			//description := args[1]
			description := strings.Join(args[1:], " ")
			tasks = addTask(tasks, description, &tasksId)
		case "update":
			if len(args) < 3 {
				fmt.Println("Usage: update [taskID] [description]")
				continue
			}
			taskID, err := strconv.Atoi(args[1])
			if err != nil {
				fmt.Println("Invalid ID type, try again")
			}
			description := strings.Join(args[1:], " ")
			//description := args[2]
			tasks = updateTask(tasks, taskID, description)
		case "delete":
			if len(args) < 2 {
				fmt.Println("Usage: delete [taskID]")
				continue
			}
			taskID, err := strconv.Atoi(args[1])
			if err != nil {
				fmt.Println("Invalid ID type, try again")
			}
			tasks = deleteTask(tasks, taskID)
		case "list":
			if len(args) == 1 {
				printTasks(tasks)
			} else {
				if len(args) < 2 {
					printTasks(tasks)
					continue
				}
				status := getStatus(args[1])
				if status < 0 {
					fmt.Println("Invalid status type, try again")
				}
				printTasksByStatus(tasks, status)
			}
		case "mark-in-progress":
			if len(args) < 2 {
				fmt.Println("Usage: mark_in_progress [taskID]")
				continue
			}
			taskID, err := strconv.Atoi(args[1])
			if err != nil {
				fmt.Println("Invalid ID type, try again")
			}
			tasks = markInProgress(tasks, taskID)
		case "mark-done":
			if len(args) < 2 {
				fmt.Println("Usage: mark_done [taskID]")
				continue
			}
			taskID, err := strconv.Atoi(args[1])
			if err != nil {
				fmt.Println("Invalid ID type, try again")
			}
			tasks = markDone(tasks, taskID)
		case "exit":
			err := saveTasks(tasks, "tasks.json")
			if err != nil {
				fmt.Println("Saved Successfully")
				os.Exit(0)
			} else {
				fmt.Println("Error saving tasks in json", err)
				os.Exit(1)
			}
		case "help":
			fmt.Println("- add [description]")
			fmt.Println("- update [id] [description]")
			fmt.Println("- delete [id]")
			fmt.Println("- list {options}: done, todo, in-progress")
			fmt.Println("- mark-in-progress [id]")
			fmt.Println("- mark-done [id]")
			fmt.Println("- exit. Save and exit")
		default:
			fmt.Println("Unknown command:", command)
			continue
		}
	}

}
