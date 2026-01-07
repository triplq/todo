package todo

import (
	"encoding/json"
	"fmt"
	"os"
	"time"
)

type item struct {
	Task        string
	CompletedAt time.Time
	CreatedAt   time.Time
	Done        bool
}

type List []item

func (l *List) Add(task string) {
	t := item{
		Task:        task,
		CreatedAt:   time.Now(),
		CompletedAt: time.Time{},
		Done:        false,
	}

	*l = append(*l, t)
}

func (l *List) Complete(i int) error {
	if i <= 0 || i > len(*l) {
		return fmt.Errorf("There are no task with id: %d", i)
	}

	(*l)[i-1].CompletedAt = time.Now()
	(*l)[i-1].Done = true

	return nil
}

func (l *List) Delete(i int) error {
	if i <= 0 || i > len(*l) {
		return fmt.Errorf("There are no task with id: %d", i)
	}

	*l = append((*l)[:i-1], (*l)[i:]...)
	return nil
}

func (l *List) Save(filename string) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	err = encoder.Encode(l)
	if err != nil {
		return err
	}

	return nil
}

func (l *List) Get(filename string) error {
	file, err := os.Open(filename)
	if err != nil {
		if os.IsNotExist(err) {
			return nil
		} else {
			return err
		}
	}
	defer file.Close()

	err = json.NewDecoder(file).Decode(l)

	if err != nil {
		return err
	}

	return nil
}
