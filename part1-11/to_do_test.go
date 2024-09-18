package main

import (
	"bytes"
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

func TestDisplay(t *testing.T) {
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
	buffer := bytes.Buffer{}
	DisplayAsJson(&buffer, todo1, todo2)

	got := buffer.String()
	want := "[{\"id\":1,\"description\":\"Wash dishes\",\"completed\":false},{\"id\":2,\"description\":\"Hoover the house\",\"completed\":true}]\n"
	
	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}