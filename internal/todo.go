package todo

import (
	"errors"
	"os"
)

type Task struct {
	Title string
	Completed bool
}

type TaskList []Task

func (tl *TaskList) Add(title string) {
	task := Task {
		Title: title,
		Completed: false,
	}

	*tl = append(*tl, task)
}

func (tl *TaskList) Complete(i int) error {
	s := *tl
	if i <= 0 || i > len(s) {
		return errors.New("Index not valid")
	}
	
	s[i-1].Completed = true
	return nil
}

func (tl *TaskList) Remove(i int) error {
	s := *tl
	if i <= 0 || i > len(s) {
		return errors.New("Index not valid")
	}

	*tl = append(s[:i-1], s[:i]...)
	return nil
}

func (tl *TaskList) Open(fileName string) error {
	_, err := os.ReadFile(fileName)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			return nil
		}
		return err
	}

	return nil
}
