package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"strings"

	"github.com/triplq/todo"
)

var filename = ".todo.json"

func main() {
	l := &todo.List{}

	if os.Getenv("TODO_FILENAME") != "" {
		filename = os.Getenv("TODO_FILENAME")
	}

	list := flag.Bool("l", false, "Show list")
	add := flag.Bool("a", false, "Add task")
	complete := flag.Int("c", 0, "Complete a task")

	flag.Parse()

	if err := l.Get(filename); err != nil {
		log.Fatal(err)
	}

	switch {
	case *list:
		fmt.Print(l)
	case *add:
		task, err := getTask(os.Stdin, flag.Args()...)

		if err != nil {
			log.Fatal(err)
		}

		l.Add(task)

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

func getTask(r io.Reader, args ...string) (string, error) {
	if len(args) > 0 {
		return strings.Join(args, " "), nil
	}
	s := bufio.NewScanner(r)
	s.Scan()
	if err := s.Err(); err != nil {
		return "", err
	}
	if len(s.Text()) == 0 {
		return "", fmt.Errorf("Task cannot be blank")
	}
	return s.Text(), nil
}
