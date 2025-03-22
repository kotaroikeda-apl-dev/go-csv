package main

import (
	"database/sql"
	"encoding/csv"
	"fmt"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

const dbURL = "user:password@tcp(localhost:3306)/dbname"

func exportToCSV() error {
	db, err := sql.Open("mysql", dbURL)
	if err != nil {
		return err
	}
	defer db.Close()

	rows, err := db.Query("SELECT id, name, age FROM users")
	if err != nil {
		return err
	}
	defer rows.Close()

	file, err := os.Create("users.csv")
	if err != nil {
		return err
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	writer.Write([]string{"ID", "Name", "Age"})

	for rows.Next() {
		var id int
		var name string
		var age int
		rows.Scan(&id, &name, &age)
		writer.Write([]string{fmt.Sprintf("%d", id), name, fmt.Sprintf("%d", age)})
	}

	fmt.Println("CSV 書き出し完了")
	return nil
}

func main() {
	err := exportToCSV()
	if err != nil {
		fmt.Println("エラー:", err)
	}
}
