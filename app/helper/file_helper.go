package helper

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type TodoCreate struct {
	ID      int
	Content string
	Status  string
}

func FileCreate() {
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

		fmt.Println("todos.txt control OK")
	}

}

func CreateTodo(todo TodoCreate) error {

	// TODO: Update todoId uniq id
	filePath := "./todos.txt"

	file, err := os.OpenFile(filePath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0755)
	if err != nil {
		return err
	}

	defer file.Close()

	line := fmt.Sprintf("%d:%s:%s\n", todo.ID, todo.Content, todo.Status)

	_, err = file.WriteString(line)
	if err != nil {
		return err
	}

	return nil

}

func ListTodo() ([]TodoCreate, error) {
	filePath := "./todos.txt"

	file, err := os.Open(filePath)
	if err != nil {
		fmt.Println("file not open", err)
		return nil, err
	}
	defer file.Close()
	var todos []TodoCreate
	reader := bufio.NewScanner(file)

	for reader.Scan() {
		line := reader.Text()
		parts := strings.Split(line, ":")

		id, err := strconv.Atoi(strings.TrimSpace(parts[0]))
		if err != nil {
			fmt.Println("Error converting ID to integer:", err)
			continue
		}

		todo := TodoCreate{
			ID:      id,
			Content: strings.TrimSpace(parts[1]),
			Status:  strings.TrimSpace(parts[2]),
		}

		todos = append(todos, todo)
	}

	if err := reader.Err(); err != nil {
		fmt.Println("file read error: ", err)
		return nil, err
	}

	return todos, nil
}

func EditTodoStatus(todoId string, newStatus string) error {
	filePath := "./todos.txt"

	file, err := os.OpenFile(filePath, os.O_RDWR, 0644)
	if err != nil {
		fmt.Println("file not open", err)
	}
	defer file.Close()

	var updatedLines []string

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()

		if strings.Contains(line, todoId) {
			parts := strings.Split(line, ":")
			if len(parts) == 3 {
				parts[2] = newStatus
				line = strings.Join(parts, ":")
			}
		}
		updatedLines = append(updatedLines, line)
	}
	if err := scanner.Err(); err != nil {
		fmt.Println("file read error: ", err)
		return err
	}

	if err := file.Truncate(0); err != nil {
		fmt.Println("file truncate error: ", err)
		return err
	}

	if _, err := file.Seek(0, 0); err != nil {
		fmt.Println("file seek error: ", err)
		return err
	}

	writer := bufio.NewWriter(file)

	for _, line := range updatedLines {
		_, err := writer.WriteString(line + "\n")
		if err != nil {
			fmt.Println("file write error: ", err)
			return err
		}

	}
	writer.Flush()
	return nil
}

func DetailTodo(todoId string) {

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

}

func PrintTodo(todos []TodoCreate) {
	fmt.Println("Todo List : ")
	for _, todo := range todos {
		fmt.Printf("Todo Id : %d Content :  %s Status : %s\n", todo.ID, todo.Content, todo.Status)
	}
}

func IsValidStatus(status string, validOptions []string) bool {
	for _, option := range validOptions {
		if strings.ToLower(status) == option {
			return true
		}
	}
	return false
}

func DeleteTodoByID(todoId string) error {
	filePath := "./todos.txt"

	file, err := os.OpenFile(filePath, os.O_RDWR, 0644)
	if err != nil {
		return err
	}
	defer file.Close()

	var updatedLines []string

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()

		if !strings.Contains(line, todoId) {
			updatedLines = append(updatedLines, line)
		}
	}

	if err := scanner.Err(); err != nil {
		return err
	}

	if err := file.Truncate(0); err != nil {
		return err
	}

	if _, err := file.Seek(0, 0); err != nil {
		return err
	}

	writer := bufio.NewWriter(file)
	for _, line := range updatedLines {
		_, err := writer.WriteString(line + "\n")
		if err != nil {
			return err
		}
	}

	writer.Flush()
	return nil
}
