package helper

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

func FileCreate() {
	//Delete print create
	filePath := "./todos.txt"

	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		file, err := os.Create(filePath)

		if err != nil {
			fmt.Println("File not create: ", err)
			return
		}
		defer file.Close()

		fmt.Println("File create : ", filePath)

	} else {

		fmt.Println("Todos List Ready")
	}

}

func CreateTodo(todo string) error {
	// TODO: Update todoId uniq id
	rand.Seed(time.Now().UnixNano())

	todoId := rand.Intn(101)

	filePath := "./todos.txt"
	file, err := os.OpenFile(filePath, os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0644)
	if err != nil {
		return fmt.Errorf("file not open : %v", err)
	}
	defer file.Close()

	_, err = file.Write([]byte(todo + " " + strconv.Itoa(todoId) + "\n"))

	if err != nil {
		return fmt.Errorf("file not write : %v", err)
	}

	fmt.Println("File write success : ", filePath)
	return nil
}

func ListTodo() error {
	//TODO: Add todo list column todoname todoid
	filePath := "./todos.txt"

	file, err := os.Open(filePath)
	if err != nil {
		return fmt.Errorf("file not open : %v", err)
	}
	defer file.Close()

	reader := bufio.NewScanner(file)

	for reader.Scan() {
		line := reader.Text()
		fmt.Println(line)
	}

	if err := reader.Err(); err != nil {
		fmt.Println("Dosya okunurken hata oluştu:", err)
	}

	return nil

}

func DetailTodo(todoId string) error {

	filePath := "./todos.txt"

	file, err := os.Open(filePath)
	if err != nil {
		fmt.Println("file not open", err)
	}

	defer file.Close()

	reader := bufio.NewScanner(file)
	for reader.Scan() {
		line := reader.Text()

		if strings.Contains(line, todoId) {
			fmt.Println(line)
		}
	}

	if err := reader.Err(); err != nil {
		fmt.Println("file read error: ", err)
	}

	return nil

}
