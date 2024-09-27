package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/gaabrieleromiti/todo/internal"
)

const (
	fileName = "todo.json"
)


func main() {
	add := flag.String("add",  "", "Create a new task")
	list := flag.Bool("list", false, "List all tasks")
	done := flag.Int("done", 0, "Mark task as done")
	remove := flag.Int("remove", 0, "Remove task")

	flag.Parse()

	tasks := internal.TaskList{}

	tasks, err := internal.Load(fileName)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	switch {
	case *add != "":
		tasks.Add(*add)
		err := internal.Save(fileName, tasks)
		if err != nil {
			log.Fatal(err)
		}
	case *list:
		err := tasks.List() 
	 	if err != nil {
			log.Fatal(err)
		}
	case *done > 0:
		err := tasks.Complete(*done)
		if err != nil {
			log.Fatal(err)
		}

		err = internal.Save(fileName, tasks)
		if err != nil {
			log.Fatal(err)
		}
	case *remove >= 0:
		err := tasks.Remove(*remove)
		if err != nil {
			log.Fatal(err)
		}

		err = internal.Save(fileName, tasks)
		if err != nil {
			log.Fatal(err)
		}
	}
}
