## **概要**

このプロジェクトは、Go 言語と Docker を使って MySQL から CSV ファイルを出力する一連の処理を学べる実践的なサンプル集です。  
基本的な CSV 出力から、エラーハンドリング、高速化、大量データの生成まで幅広くカバーしています。

## **実行方法**

```sh
docker compose up -d # データベース起動
go run cmd/basic/main.go # データベースからCSVを出力
go run cmd/error/main.go # エラーハンドリングを含めたCSV出力
go run cmd/insert-create/main.go # 100000件のデータをINSERTするSQLを作成
go run cmd/high-speed/main.go # 高速なCSV出力
docker compose down # データベース停止
```

## **学習ポイント**

1. **`csv.NewWriter()`** を使って CSV ファイルを作成する。
2. 検索条件付き SQL にプレースホルダ（**`LIKE ?`**）を使うことで安全にフィルタ検索ができる。
3. CSV 出力時のエラーやデータ未取得時にも適切なハンドリングでユーザーに明確なメッセージを返せる。
4. **`os.Create`** を使って、新しいファイルを作成する。
5. **`writer.Write()`** を使って、そのファイルにデータを追加する。
6. **`writer.Flush()`** を使って、データをファイルに書き込む。

## 作成者

- **池田虎太郎** | [GitHub プロフィール](https://github.com/kotaroikeda-apl-dev)
