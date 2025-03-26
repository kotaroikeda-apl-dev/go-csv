package main

import (
	"database/sql"
	"encoding/csv"
	"fmt"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

const dbURL = "user:password@tcp(localhost:3306)/dbname"

func exportLargeCSV() error {
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

	file, err := os.Create("large_users.csv")
	if err != nil {
		return err
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	writer.Write([]string{"ID", "Name", "Age"}) // ヘッダー書き込み

	count := 0
	for rows.Next() {
		var id int
		var name string
		var age int
		rows.Scan(&id, &name, &age)
		writer.Write([]string{fmt.Sprintf("%d", id), name, fmt.Sprintf("%d", age)})
		count++

		if count%1000 == 0 { // 1000行ごとにフラッシュしてメモリ使用量を抑える
			writer.Flush()
		}
	}

	fmt.Printf("CSV 書き出し完了（%d件）\n", count)
	return nil
}

func main() {
	err := exportLargeCSV()
	if err != nil {
		fmt.Println("エラー:", err)
	}
}
