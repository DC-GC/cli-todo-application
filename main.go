package main

func main() {

	todoList := TodoList{}
	storage := NewStorage[TodoList]("TodoList.json")
	storage.Load(&todoList)
	todoList.displayAll()
	storage.Save(todoList)
}
