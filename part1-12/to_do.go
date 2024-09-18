package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type ToDo struct {
	ID          int		`json:"id"`
	Description string	`json:"description"`
	Completed   bool	`json:"completed"`
}

func (t ToDo) String() string {
	return fmt.Sprintf("ID: %d, Description: %s, Completed: %t", t.ID, t.Description, t.Completed)
}

func SaveToFile(filename string, todos ...ToDo) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()
	
	todoBytes, _ := json.Marshal(todos)
	_, err = file.Write(todoBytes)
	if err != nil {
		return err
	}
	return nil
}

func main() {
	todos := generateToDos()
	SaveToFile("/tmp/todos.json", todos...)
}

func generateToDos() []ToDo {
	var todos []ToDo
	for i := 1; i <= 10; i++ {
		todos = append(todos, ToDo{
			ID:          i,
			Description: fmt.Sprintf("Task %d", i),
			Completed:   false,
		})
	}
	return todos
}