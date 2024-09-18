package main

import (
	"testing"
)

func TestToDo(t *testing.T) {
	todo := ToDo{
		ID:          1,
		Description: "Complete assignment",
		Completed:   false,
	}
	
	got := todo.String()
	want := "ID: 1, Description: Complete assignment, Completed: false"
	
	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}

func TestSaveToFile(t *testing.T) {
	todo1 := ToDo{
		ID:          1,
		Description: "Wash dishes",
		Completed:   false,
	}
	todo2 := ToDo{
		ID:          2,
		Description: "Hoover the house",
		Completed:   true,
	}
	
	err := SaveToFile("/tmp/todos.json", todo1, todo2)
	if err != nil {
		t.Errorf("Error saving todos to file: %v", err)
	}
}