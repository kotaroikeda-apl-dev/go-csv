## **概要**

## **実行方法**

```sh
docker compose up -d # データベース起動
go run cmd/basic/main.go # データベースからCSVを出力
go run cmd/error/main.go # エラーハンドリングを含めたCSV出力
docker compose down # データベース停止
```

## **学習ポイント**

1. **`csv.NewWriter()`** を使って CSV ファイルを作成し、**`writer.Flush()`** でデータをファイルに書き込む。
2. 検索条件付き SQL にプレースホルダ（**`LIKE ?`**）を使うことで安全にフィルタ検索ができる。
3. CSV 出力時のエラーやデータ未取得時にも適切なハンドリングでユーザーに明確なメッセージを返せる。

## 作成者

- **池田虎太郎** | [GitHub プロフィール](https://github.com/kotaroikeda-apl-dev)
