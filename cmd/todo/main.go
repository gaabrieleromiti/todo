package main

import (
	"flag"
	"fmt"
	"os"
	"log"
	"github.com/gaabrieleromiti/todo"
)

const (
	fileName = "todo.json"
)

func main() {
	add := flag.Bool("new",  false,"Create a new task")
	list := flag.Bool("list", false, "List all tasks")
	done := flag.Int("done", 0, "Mark task as done")
	remove := flag.Int("remove", 0, "Remove task")

	flag.Parse()

	tasks := todo.TaskList{}

	err := tasks.Load(fileName)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	switch {
	case *add:
		tasks.Add(flag.Args())

		err := tasks.Save(fileName)
		if err != nil {
			log.Fatal(err)
		}
	case *list:
		err := tasks.List() 
		if err != nil {
			log.Fatal(err)
		}
		
		err = tasks.Save(fileName)
		if err != nil {
			log.Fatal(err)
		}
	case *done > 0:
		err := tasks.Complete(*done)
		if err != nil {
			log.Fatal(err)
		}

		err = tasks.Save(fileName)
		if err != nil {
			log.Fatal(err)
		}
	case *remove > 0:
		err := tasks.Remove(*remove)
		if err != nil {
			log.Fatal(err)
		}

		err = tasks.Save(fileName)
		if err != nil {
			log.Fatal(err)
		}
	}
}
