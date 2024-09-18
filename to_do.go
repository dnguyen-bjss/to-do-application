package main

import (
	"encoding/json"
	"fmt"
	"io"
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

func Display(writer io.Writer, todos ...ToDo) {
	for _, todo := range todos {
		todoBytes, _ := json.Marshal(todo)
		fmt.Fprintf(writer, "%s\n", string(todoBytes))
	}
}

func main() {
	todos := generateToDos()
	Display(os.Stdout, todos...)
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