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
	want := "1 - Complete assignment - false"
	
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
	Display(&buffer, todo1, todo2)

	got := buffer.String()
	want := "1 - Wash dishes - false\n2 - Hoover the house - true\n"
	
	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}