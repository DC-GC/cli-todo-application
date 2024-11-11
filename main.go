package main

func main() {

	todoList := TodoList{}
	storage := NewStorage[TodoList]("TodoList.json")
	storage.Load(&todoList)
	commandFlags := NewCommandFlags()
	commandFlags.Execute(&todoList)
	// todoList.displayAll()
	storage.Save(todoList)
}
