package todo

import (
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
