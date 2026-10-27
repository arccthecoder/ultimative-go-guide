package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "modernc.org/sqlite" // зависимость, чтобы запускать без CGO (компилятора C)
)

func main() {
	if err := run(); err != nil {
		log.Fatal(err)
	}
}

func run() error {
	// создаем файл, если его нет
	db, err := sql.Open("sqlite", "tasks.db")
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
		"Разораться с SQLite", // при каждом запуске будет создаваться задача в дб (можно менять название)
	)
	if err != nil {
		return err
	}

	id, err := res.LastInsertId()
	if err != nil {
		return err
	}
	fmt.Println("Создана задача:", id)

	// выполняем задачу
	var taskID int

	fmt.Println("Введите номер задачи:")
	_, err = fmt.Scanln(&taskID)
	if err != nil {
		return err
	}

	res, err = db.Exec(
		"UPDATE tasks SET done = 1 WHERE id = ?",
		taskID,
	)
	if err != nil {
		return err
	}

	count, err := res.RowsAffected()
	if err != nil {
		return err
	}

	if count == 0 {
		fmt.Println("Задача с таким номером не найдена")
	} else {
		fmt.Println("Задача отмечена выполненной")
	}

	// получаем все задачи
	rows, err := db.Query(
		"SELECT id, title, done FROM tasks ORDER BY id",
	)
	if err != nil {
		return err
	}
	defer rows.Close()

	for rows.Next() {
		var id int64
		var title string
		var done bool

		if err := rows.Scan(&id, &title, &done); err != nil {
			return err
		}

		fmt.Printf("%d. %s | выполнена: %t\n", id, title, done)
	}

	return rows.Err()
}
