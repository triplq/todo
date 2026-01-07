package todo_test

import (
	"fmt"
	"os"
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

func TestComplete(t *testing.T) {
	l := todo.List{}

	l.Add("Eat")
	l.Add("Sleep")

	l.Complete(1)

	if l[0].Done != true {
		t.Errorf("NOT DONE")
	} else {
		fmt.Println(l[0].CompletedAt)
	}
}

func TestDelete(t *testing.T) {
	l := todo.List{}

	l.Add("Eat")
	l.Add("Sleep")
	l.Add("Run")

	if l[0].Task != "Eat" {
		t.Errorf("Expected %q, got %q instead.", "Eat", l[0].Task)
	}
	l.Delete(2)
	if len(l) != 2 {
		t.Errorf("Expected list length %d, got %d instead.", 2, len(l))
	}
	if l[1].Task != "Run" {
		t.Errorf("Expected %q, got %q instead.", "Run", l[1].Task)
	}

}

func TestSaveGet(t *testing.T) {
	l1 := todo.List{}
	l2 := todo.List{}

	l1.Add("Eat")
	l1.Add("Sleep")
	l1.Add("Run")

	l1.Save("testing1.txt")
	defer os.Remove("testing1.txt")

	l2.Get("testing1.txt")

	if l1[0].Task != l2[0].Task {
		t.Errorf("Slices r not same")
	}
}
