package main

import (
	"reflect"
	"testing"
)

func TestReadFromFile(t *testing.T) {
		got, err := ReadFromFileAsJson("/tmp/todos.json")
		want := []ToDo{{ID: 1, Description: "Wash dishes", Completed: false}, {ID: 2, Description: "Hoover the house", Completed: true}}

		if err != nil {
			t.Errorf("Error reading todos from file: %v", err)
		}
		
		for i, todo := range got {
			if !reflect.DeepEqual(todo, want[i]) {
				t.Errorf("got %v, want %v", todo, want[i])
			}
		}
}