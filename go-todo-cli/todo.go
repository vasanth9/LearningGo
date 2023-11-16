package todo

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"time"
)

type item struct {
	Task        string
	Done        bool
	CreatedAt   time.Time
	CompletedAt time.Time
}

type List []item

// adding the todo
func (l *List) Add(task string) {
	t := item{
		Task:        task,
		Done:        false,
		CreatedAt:   time.Now(),
		CompletedAt: time.Time{},
	}
	*l = append(*l, t)
}

// marking the todo as completed
func (l *List) Complete(i int) error {
	ls := *l
	if i <= 0 || i > len(ls) {
		return fmt.Errorf("item %d does not exist", i)
	}
	ls[i-1].Done = true
	ls[i-1].CompletedAt = time.Now()
	return nil
}

// saving the json format in a file
func (l *List) Save(fileName string) error {
	json, err := json.Marshal(l)
	if err != nil {
		return err
	}
	return os.WriteFile(fileName, json, 0644)
}

// reading the stored todo-s from the saved file
func (l *List) Get(fileName string) error {
	file, err := os.ReadFile(fileName)
	if err != nil {
		// if the given files does not exists
		if errors.Is(err, os.ErrNotExist) {
			return nil
		}
		return err
	}
	if len(file) == 0 {
		return nil
	}
	return json.Unmarshal(file, l)
}
