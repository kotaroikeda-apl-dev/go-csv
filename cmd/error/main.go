package main

import (
	"database/sql"
	"encoding/csv"
	"fmt"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

const dbURL = "user:password@tcp(localhost:3306)/dbname"

func exportToCSV(filter string) error {
	db, err := sql.Open("mysql", dbURL)
	if err != nil {
		return fmt.Errorf("DB接続エラー: %w", err)
	}
	defer db.Close()

	query := "SELECT id, name, age FROM users WHERE name LIKE ?"
	rows, err := db.Query(query, "%"+filter+"%")
	if err != nil {
		return fmt.Errorf("クエリ実行エラー: %w", err)
	}
	defer rows.Close()

	file, err := os.Create("users.csv")
	if err != nil {
		return fmt.Errorf("ファイル作成エラー: %w", err)
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	writer.Write([]string{"ID", "Name", "Age"})

	count := 0
	for rows.Next() {
		var id int
		var name string
		var age int
		rows.Scan(&id, &name, &age)
		writer.Write([]string{fmt.Sprintf("%d", id), name, fmt.Sprintf("%d", age)})
		count++
	}

	if count == 0 {
		return fmt.Errorf("検索条件に一致するデータがありません")
	}

	fmt.Println("CSV 書き出し完了")
	return nil
}

func main() {
	err := exportToCSV("A") // "A" を含む名前のユーザーのみ取得
	if err != nil {
		fmt.Println("エラー:", err)
	}
}
