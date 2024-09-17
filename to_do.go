package main

import (
	"fmt"
	"io"
	"os"
)

type ToDo struct {
	ID          int
	Description string
	Completed   bool
}

func (t ToDo) String() string {
	return fmt.Sprintf("%d - %s - %t", t.ID, t.Description, t.Completed)
}

func Display(writer io.Writer, todos ...ToDo) {
	for _, todo := range todos {
		fmt.Fprintf(writer, "%s\n", todo.String())
	}
}

func main() {
	var todos []ToDo
	for i := 1; i <= 10; i++ {
		todos = append(todos, ToDo{
			ID:          i,
			Description: fmt.Sprintf("Task %d", i),
			Completed:   false,
		})
	}
	Display(os.Stdout, todos...)
}