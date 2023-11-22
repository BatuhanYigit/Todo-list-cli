package helper

import (
	"fmt"
	"os"
)

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
		fmt.Println("Todos List Ready")
	}

}
