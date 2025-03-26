package main

import (
	"fmt"
	"os"
)

func main() {
	file, err := os.Create("insert_users.sql")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	fmt.Fprintln(file, "INSERT INTO users (name, age) VALUES")

	for i := 1; i <= 100000; i++ {
		name := fmt.Sprintf("'User%d'", i)
		age := 20 + (i % 30) // 20〜49歳くらいの年齢を分散

		// 最後の行だけセミコロンをつける
		if i == 100000 {
			fmt.Fprintf(file, "(%s, %d);\n", name, age)
		} else {
			fmt.Fprintf(file, "(%s, %d),\n", name, age)
		}
	}

	fmt.Println("insert_users.sql ファイルが作成されました。")
}
