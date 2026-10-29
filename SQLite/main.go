package main

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/jmoiron/sqlx" // более удобная для Go
	_ "modernc.org/sqlite"    // зависимость, чтобы запускать без CGO (компилятора C)
)

type Task struct {
	ID    int    `db:"id"`
	Title string `db:"title"`
	Done  bool   `db:"done"`
}

func main() {
	if err := run(); err != nil {
		log.Fatal(err)
	}
}

func run() error {
	// создаем файл, если его нет
	db, err := sqlx.Connect("sqlite", "tasks.db")
	if err != nil {
		return err
	}
	defer db.Close()

	// делаем 1 подключение на данном примере
	db.SetMaxOpenConns(1)

	// создаем таблицу один раз
	_, err = db.Exec(`
	CREATE TABLE IF NOT EXISTS tasks (
	id INTEGER PRIMARY KEY,
	title TEXT NOT NULL,
	done INTEGER NOT NULL DEFAULT 0
	)
	`)
	if err != nil {
		return err
	}

	// добавляем задачу
	res, err := db.Exec(
		"INSERT INTO tasks (title) VALUES (?)",
		"Разобраться с SQLite", // при каждом запуске будет создаваться задача в дб (можно менять название)
	)
	if err != nil {
		return err
	}

	id, err := res.LastInsertId()
	if err != nil {
		return err
	}
	fmt.Println("Создана задача:", id)

	// получаем все задачи
	var tasks []Task

	err = db.Select(
		&tasks,
		"SELECT id, title, done FROM tasks ORDER BY id",
	)
	if err != nil {
		return err
	}

	for _, task := range tasks {
		fmt.Printf(
			"%d. %s | выполнена: %t\n",
			task.ID, task.Title, task.Done,
		)
	}

	// Пользователь выбирает задачу из списка.
	var taskID int64

	fmt.Print("Введите номер задачи для выполнения: ")
	_, err = fmt.Scanln(&taskID)
	if err != nil {
		return err
	}

	// Получаем выбранную задачу
	var task Task

	err = db.Get(
		&task,
		"SELECT id, title, done FROM tasks WHERE id = ?",
		taskID,
	)
	if err == sql.ErrNoRows {
		fmt.Println("Задача с таким номером не найдена")
		return nil
	}
	if err != nil {
		return err
	}

	if task.Done {
		fmt.Println("Задача уже выполнена")
		return nil
	}

	_, err = db.Exec(
		"UPDATE tasks SET done = 1 WHERE id = ?",
		taskID,
	)

	if err != nil {
		return err
	}

	fmt.Println("Задача отмечена выполненной:", task.Title)
	return nil
}
