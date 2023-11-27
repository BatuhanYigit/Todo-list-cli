package main

import (
	"fmt"
	"os"
	"todolist/helper"
)

func main() {
	for {
		// helper.FileCreate()
		helper.WelcomeMessage()

		fmt.Println("Select choice : ")
		var choice int
		fmt.Scanln(&choice)

		if choice == 1 {
			todos, err := helper.ListTodo()
			if err != nil {
				fmt.Println("Error:", err)
				return
			}

			for _, todo := range todos {
				fmt.Printf("ID: %d, Content: %s, Status: %s\n", todo.ID, todo.Content, todo.Status)
			}
		}

		if choice == 2 {
			filePath := "todos.txt"

			file, err := os.OpenFile(filePath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0755)

			if err != nil {
				fmt.Println("Dosya açılamadı:", err)
				return
			}
			defer file.Close()

			todos, err := helper.ListTodo()
			if err != nil {
				fmt.Println("Todo List Not Read : ", err)
			}
			fmt.Print("Write todo names : ")
			var todoName string
			fmt.Scanln(&todoName)
			fmt.Print("Write todo status : ")
			var status string
			fmt.Scanln(&status)

			fmt.Println("Todosss : ", len(todos))
			fmt.Println("Todos list : ", todos)

			id := len(todos) + 1
			create := helper.TodoCreate{
				ID:      id,
				Content: todoName,
				Status:  status,
			}

			todos = append(todos, create)

			err = helper.CreateTodo(create)
			if err != nil {
				fmt.Println("Todo not create ", err)
				return
			}

			helper.PrintTodo(todos)

		}

		if choice == 4 {
			fmt.Print("Write Todo ID : ")
			var todoId string
			fmt.Scanln(&todoId)

			helper.DetailTodo(todoId)
		}
	}
}
