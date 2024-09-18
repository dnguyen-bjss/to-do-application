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

func ReadFromFileAsJson(filename string) ([]ToDo, error) {
	file, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	var todos []ToDo
	err = json.Unmarshal(file, &todos)
	if err != nil {
		return nil, err
	}
	return todos, nil
}

func main () {
	filename := "/tmp/todos.json"
	todos, err := ReadFromFileAsJson(filename)
	if err != nil {
		panic(err)
	}
	for _, todo := range todos {
		fmt.Println(todo)
	}
}