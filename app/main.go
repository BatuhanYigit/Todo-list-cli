package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"todolist/database"
	"todolist/helper"
)

func main() {
	helper.SelectDatabase()
	fmt.Println("Select Choice Db : ")
	var choiceDatabase int
	fmt.Scanln(&choiceDatabase)

	if choiceDatabase == 1 {
		for {

			database.DbCreate()
			helper.WelcomeMessage()

			fmt.Println("Select choice : ")
			var choice int
			fmt.Scanln(&choice)

			switch choice {
			case 1:
				database.GetAll()
				todos, err := database.GetAll()
				if err != nil {
					fmt.Println("Todo Get failed : ", err)
				}
				fmt.Println("Todo List : ")
				for _, todo := range todos {
					fmt.Printf("ID %d, Content: %s, Status: %s\n", todo.Id, todo.Content, todo.Status)
				}

			case 2:
				validStatusOptions := []string{"todo", "in progress", "done"}

				fmt.Print("Write todo names : ")
				todoName, _ := bufio.NewReader(os.Stdin).ReadString('\n')

				fmt.Print("Write todo status choice (todo,in progress,done) :  ")
				status, _ := bufio.NewReader(os.Stdin).ReadString('\n')

				status = strings.TrimSpace(status)

				if helper.IsValidStatus(status, validStatusOptions) {

					create := database.TodoCreate{
						Content: todoName,
						Status:  status,
					}

					err := database.CreateTodo(create)
					if err != nil {
						fmt.Println("Todo not create ", err)
						return
					}
					fmt.Println("Todo succesfuly create")

				} else {
					fmt.Println("Invalid status please choice (todo,in progress, done) your choice = ", status)
				}
			case 3:
				validStatusOptions := []string{"todo", "in progress", "done"}

				fmt.Print("Write Edit Todo ID : ")
				var todoId int
				fmt.Scanln(&todoId)

				fmt.Print("Edit status choice (todo,in progress,done) : ")
				var todoStatus string
				fmt.Scanln(&todoStatus)

				if helper.IsValidStatus(todoStatus, validStatusOptions) {
					database.Update(todoId, todoStatus)
					fmt.Println("Status Update")

				} else {
					fmt.Println("Invalid status please choice (todo,in progress, done) your choice = ", todoStatus)
				}
			case 4:
				fmt.Println("Write delete todo id : ")
				var todoId int
				fmt.Scanln(&todoId)

				err := database.Delete(todoId)
				if err != nil {
					fmt.Println("Delete failed", err)

				}
				fmt.Println("Todo deleted")

			default:
				fmt.Println("Invalid choice.")

			}

		}
	}

	if choiceDatabase == 2 {
		for {

			helper.FileCreate()
			helper.WelcomeMessage()

			fmt.Println("Select choice : ")
			var choice int
			fmt.Scanln(&choice)

			switch choice {
			case 1:
				todos, err := helper.ListTodo()
				if err != nil {
					fmt.Println("Error:", err)
					return
				}

				for _, todo := range todos {
					fmt.Printf("ID: %d, Content: %s, Status: %s\n", todo.ID, todo.Content, todo.Status)
				}

			case 2:
				validStatusOptions := []string{"todo", "in progress", "done"}
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
				todoName, _ := bufio.NewReader(os.Stdin).ReadString('\n')

				fmt.Print("Write todo status choice (todo,in progress,done) :  ")
				status, _ := bufio.NewReader(os.Stdin).ReadString('\n')

				todoName = strings.TrimSpace(todoName)
				status = strings.TrimSpace(status)

				if helper.IsValidStatus(status, validStatusOptions) {

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

				} else {
					fmt.Println("Invalid status please choice (todo,in progress, done) your choice = ", status)
				}

			case 3:
				validStatusOptions := []string{"todo", "in progress", "done"}

				fmt.Print("Write Edit Todo ID : ")
				var todoId string
				fmt.Scanln(&todoId)
				helper.DetailTodo(todoId)

				fmt.Print("Edit status choice (todo,in progress,done) : ")
				var todoStatus string
				fmt.Scanln(&todoStatus)

				if helper.IsValidStatus(todoStatus, validStatusOptions) {
					helper.EditTodoStatus(todoId, todoStatus)
					fmt.Println("Status Update")

				} else {
					fmt.Println("Invalid status please choice (todo,in progress, done) your choice = ", todoStatus)
				}

			// case 4:
			// 	fmt.Print("Write Todo ID : ")
			// 	var todoId string
			// 	fmt.Scanln(&todoId)

			// 	helper.DetailTodo(todoId)

			case 4:
				fmt.Print("Delete Todo ID : ")
				var todoId string
				fmt.Scanln(&todoId)

				err := helper.DeleteTodoByID(todoId)
				if err != nil {
					fmt.Println("Error deleting todo ", err)
				} else {
					fmt.Println("Todo Deleted id : ", todoId)
				}

			default:
				fmt.Println("Invalid choice.")
			}
		}
	}
}
