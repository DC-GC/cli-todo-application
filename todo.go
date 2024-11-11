package main

import (
	"errors"
	"fmt"
	"time"
)

type Activity struct {
	Title       string
	Completed   bool
	CreatedOn   time.Time
	CompletedOn *time.Time
	LastUpdated time.Time
}

type TodoList []Activity

func (todoList *TodoList) add(t string) {
	newActivity := Activity{
		Title:       t,
		Completed:   false,
		CreatedOn:   time.Now(),
		CompletedOn: nil,
		LastUpdated: time.Now(),
	}

	*todoList = append(*todoList, newActivity)
}

func (todoList *TodoList) validateIndex(i int) error {
	if i < 0 || i >= len(*todoList) {
		err := errors.New("Invalid index")
		fmt.Println(err)
		return err
	}
	return nil
}

func (todoList *TodoList) delete(i int) error {
	t := *todoList

	if err := t.validateIndex(i); err != nil {
		return err
	}
	*todoList = append(t[:i], t[i+1:]...)
	return nil

}

func (todoList *TodoList) toggle(i int) error {
	t := *todoList

	if err := t.validateIndex(i); err != nil {
		return err
	}

	isCompleted := t[i].Completed

	if !isCompleted {
		completionTime := time.Now()
		t[i].CompletedOn = &completionTime
	}

	t[i].Completed = !isCompleted

	return nil
}

func (todoList *TodoList) update(i int, title string) error {
	t := *todoList

	if err := t.validateIndex(i); err != nil {
		return err
	}

	t[i].Title = title
	updateTime := time.Now()
	t[i].CompletedOn = &updateTime

	return nil
}
