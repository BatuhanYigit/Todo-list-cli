package main

import (
	"fmt"
	"todolist/helper"
)

func main() {
	helper.FileCreate()
	helper.WelcomeMessage()

	fmt.Println("Select choice : ")
	var choice int
	fmt.Scanln(&choice)

	if choice == 1 {
		fmt.Print("****** TODO LIST ******")
		helper.ListTodo()
	}

	if choice == 2 {
		fmt.Print("Write todo names : ")
		var todoName string
		fmt.Scanln(&todoName)

		helper.CreateTodo(todoName)
	}

	if choice == 4 {
		fmt.Print("Write Todo ID : ")
		var todoId string
		fmt.Scanln(&todoId)

		helper.DetailTodo(todoId)
	}

}
