package internal

import (
	"encoding/json"
	"errors"
	"os"
	"slices"
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

func (tl *TaskList) List() error {
	if len(*tl) == 0 {
		println("No tasks to display")
		os.Exit(0)
	}
	for i, task := range *tl {
		check := " "
		if task.Completed {
			check = "X"
		}
		println(i+1, task.Title ,check)
	}
	return nil
}

func (tl *TaskList) Complete(i int) error {
	if i <= 0 || i > len(*tl) {
		return errors.New("Index not valid")
	}
	
	(*tl)[i-1].Completed = true
	return nil
}

func (tl *TaskList) Remove(i int) error {
	if i <= 0 || i > len(*tl) {
		return errors.New("Index not valid")
	}

	*tl = slices.Delete(*tl, i-1, i) 
	return nil
}

func Load(fileName string) (TaskList, error) {
	fileContent, err := os.ReadFile(fileName)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			return TaskList{}, nil
		}
		return TaskList{}, err
	}

	if len(fileContent) == 0 {
		return TaskList{}, nil
	}


	tl := TaskList{}
	err = json.Unmarshal(fileContent, &tl)
	if err != nil {
		return TaskList{}, err
	}

	return tl, nil
}

func Save(fileName string, tl TaskList) error {
	fileContent, err := json.Marshal(tl)
	if err != nil {
		return err
	}

	err = os.WriteFile(fileName, fileContent, 0644)
	if err != nil {
		return err
	}

	return nil
}


