package main

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/triplq/todo"
)

const filename = ".todo.json"

func main() {
	l := &todo.List{}

	if err := l.Get(filename); err != nil {
		log.Fatal(err)
	}

	switch {
	case len(os.Args) == 1:
		for _, item := range *l {
			fmt.Println(item.Task)
		}

	default:
		item := strings.Join(os.Args[1:], " ")
		l.Add(item)

		if err := l.Save(filename); err != nil {
			log.Fatal(err)
		}
	}

}
