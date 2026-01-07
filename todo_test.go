package todo_test

import (
	"testing"

	"github.com/triplq/todo"
)

func TestAdd(t *testing.T) {
	l := todo.List{}

	taskName := "Eat"
	l.Add(taskName)

	if l[0].Task != taskName {
		t.Errorf("Error add")
	}
}
