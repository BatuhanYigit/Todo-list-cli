package helper

import "fmt"

func WelcomeMessage() {
	fmt.Println("********************ToDo List V0.1********************")
	fmt.Println("")
	fmt.Println("1 - List All")
	fmt.Println("2 - Add Todo")
	fmt.Println("3 - Status Edit Todo")
	fmt.Println("4 - Delete Todo")
}

func SelectDatabase() {
	fmt.Println("********************ToDo List V0.1********************")
	fmt.Println("")
	fmt.Println("Select Database")
	fmt.Println("1 - Sqlite3")
	fmt.Println("2 - Storage to txt file")
}
