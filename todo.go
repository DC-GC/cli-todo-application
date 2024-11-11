package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/aquasecurity/table"
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

func (todoList *TodoList) displayAll() {
	displayTable := table.New(os.Stdout)
	displayTable.SetRowLines(false)
	displayTable.SetHeaders("", "Activity", "Completed", "Created At", "Completed At", "Last Updated")
	displayTable.SetAlignment(table.AlignCenter, table.AlignLeft, table.AlignCenter, table.AlignLeft, table.AlignLeft, table.AlignLeft)
	for index, t := range *todoList {
		completed := "❌"
		completedOn := ""

		if t.Completed {
			completed = "✅"
			if t.CompletedOn != nil {
				completedOn = t.CompletedOn.Format(time.UnixDate)
			}
		}

		displayTable.AddRow(strconv.Itoa(index), t.Title, completed, t.CreatedOn.Format(time.UnixDate), completedOn, t.LastUpdated.Format(time.UnixDate))
	}
	displayTable.Render()
}
