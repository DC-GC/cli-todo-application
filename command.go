package main

import (
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type CommandFlags struct {
	Add        string
	Del        int
	Update     string
	Toggle     int
	DisplayAll bool
}

func NewCommandFlags() *CommandFlags {
	cf := CommandFlags{}
	flag.StringVar(&cf.Add, "Add", "", "Add a new activity.")
	flag.StringVar(&cf.Update, "Update", "", "Update activity.")

	flag.IntVar(&cf.Del, "Delete", -1, "Remove Activity by giving index.")
	flag.IntVar(&cf.Toggle, "Toggle", -1, "Mark activity as complete / incomplete.")

	flag.BoolVar(&cf.DisplayAll, "DisplayAll", false, "View all activities.")

	flag.Parse()
	return &cf
}

func (cf *CommandFlags) Execute(todos *TodoList) {
	switch {
	case cf.DisplayAll:
		todos.displayAll()
	case cf.Add != "":
		todos.add(cf.Add)
	case cf.Update != "":
		parts := strings.SplitN(cf.Update, ";", 2)
		if len(parts) != 2 {
			fmt.Println("Error: invalid command format. Please use \"<id>:<new_title>\" ")
			os.Exit(1)
		}

		index, err := strconv.Atoi(parts[0])

		if err != nil {
			fmt.Println("Error: invalid index provided.")
			os.Exit(1)
		}
		todos.update(index, parts[1])

	case cf.Del != -1:
		todos.delete(cf.Del)
	case cf.Toggle != -1:
		todos.toggle(cf.Toggle)
	default:
		fmt.Println("Error: invalid command.")
	}

}
