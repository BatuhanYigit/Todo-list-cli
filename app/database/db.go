package database

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/mattn/go-sqlite3"
)

type TodoCreate struct {
	Content string
	Status  string
}

type TodoAll struct {
	Id      int
	Content string
	Status  string
}

func DbCreate() {
	// Database file location change name or path
	dbPath := "./database/test.db"

	if _, err := os.Stat(dbPath); os.IsNotExist(err) {
		db, err := sql.Open("sqlite3", dbPath)
		if err != nil {
			fmt.Println("Db connection failed", err)
			return
		}
		defer db.Close()

		_, err = db.Exec(`CREATE TABLE IF NOT EXISTS todo(
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			content TEXT NOT NULL,
			status TEXT NOT NULL

		)
		`)
		if err != nil {
			fmt.Println("Table create failed: ", err)
			return
		}

		fmt.Println("Db file and table successfully create")
	}

}

func CreateTodo(todo TodoCreate) error {
	dbPath := "./database/test.db"
	db, err := sql.Open("sqlite3", dbPath)
	if err != nil {
		log.Fatal(err)
	} else {
		_, err = db.Exec("INSERT INTO todo (content,status) VALUES (?,?)", todo.Content, todo.Status)
		if err != nil {
			log.Fatal(err)
		}
	}
	defer db.Close()
	return nil

}

func GetAll() ([]TodoAll, error) {
	dbPath := "./database/test.db"
	db, err := sql.Open("sqlite3", dbPath)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	defer db.Close()

	rows, err := db.Query("SELECT * FROM todo")
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	defer rows.Close()

	var todos []TodoAll

	for rows.Next() {
		var todo TodoAll
		err := rows.Scan(&todo.Id, &todo.Content, &todo.Status)
		if err != nil {
			log.Fatal(err)
			return nil, err
		}
		todos = append(todos, todo)
	}

	if err = rows.Err(); err != nil {
		log.Fatal(err)
		return nil, err
	}

	return todos, nil
}

func Update(id int, status string) error {
	dbPath := "./database/test.db"
	db, err := sql.Open("sqlite3", dbPath)
	if err != nil {
		log.Fatal(err)
		return err
	}
	defer db.Close()

	_, err = db.Exec("UPDATE todo SET status = ? WHERE id = ?", status, id)
	if err != nil {
		log.Fatal(err)
		return err
	}
	return nil

}

func Delete(id int) error {
	dbPath := "./database/test.db"
	db, err := sql.Open("sqlite3", dbPath)
	if err != nil {
		log.Fatal(err)
		return err
	}
	defer db.Close()

	_, err = db.Exec("DELETE FROM todo WHERE id = ?", id)
	if err != nil {
		log.Fatal(err)
		return err
	}
	return nil
}
