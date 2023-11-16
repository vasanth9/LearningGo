package todo_test

import (
	todo "github.com/vasanth9/TodoCli"
	"os"
	"testing"
)

func TestAdd(t *testing.T) {
	l := todo.List{}
	taskName := "New Task"
	l.Add(taskName)

	if l[0].Task != taskName {
		t.Errorf("Expected %q , got %q instead", taskName, l[0].Task)
	}

}
func TestComplete(t *testing.T) {
	l := todo.List{}
	taskName := "New Task"
	l.Add(taskName)

	if l[0].Task != taskName {
		t.Errorf("Expected %q , got %q instead", taskName, l[0].Task)
	}

	if l[0].Done {
		t.Errorf("New Task should not be completed")
	}
	l.Complete(1)

	if !l[0].Done {
		t.Errorf("New Task should be completed")
	}
}

func TestSaveGet(t *testing.T) {
	// two lists
	l1 := todo.List{}
	l2 := todo.List{}
	// saving task into l1 and loading it into l2 -- error if fails
	taskName := "New Task"
	l1.Add(taskName)

	if l1[0].Task != taskName {
		t.Errorf("Expected %q , got %q instead", taskName, l1[0].Task)
	}
	tf, err := os.CreateTemp("", "")
	if err != nil {
		t.Errorf("Error creating temp file: %s", err)
	}
	defer os.Remove(tf.Name())
	if err := l1.Save(tf.Name()); err != nil {
		t.Fatalf("Error saving list to file: %s", err)
	}
	if err := l2.Get(tf.Name()); err != nil {
		t.Fatalf("Error getting list from file: %s", err)
	}
	if l1[0].Task != l2[0].Task {
		t.Errorf("Task %q should match %q task.", l1[0].Task, l2[0].Task)
	}
}
