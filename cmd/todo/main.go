package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/triplq/todo"
)

const filename = ".todo.json"

func main() {
	l := &todo.List{}

	list := flag.Bool("l", false, "Show list")
	task := flag.String("t", "", "Add task")
	complete := flag.Int("c", 0, "Complete a task")

	flag.Parse()

	if err := l.Get(filename); err != nil {
		log.Fatal(err)
	}

	switch {
	case *list:
		for _, item := range *l {
			if !item.Done {
				fmt.Println(item.Task)
			}
		}
	case *task != "":
		l.Add(*task)

		if err := l.Save(filename); err != nil {
			log.Fatal(err)
		}
	case *complete > 0:
		if err := l.Complete(*complete); err != nil {
			log.Fatal(err)
		}
		if err := l.Save(filename); err != nil {
			log.Fatal(err)
		}
	default:
		fmt.Fprintln(os.Stderr, "Invalid flag is provided")
		os.Exit(1)
	}
}
