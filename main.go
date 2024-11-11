package main

import "fmt"

func main() {
	fmt.Print("Hello world......")

	todoList := TodoList{}
	todoList.add("create activity")
	todoList.add("create 2nd activity")
	fmt.Printf("%+v\n\n", todoList)
	todoList.toggle(0)
	todoList.displayAll()
}
