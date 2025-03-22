## **概要**

## **実行方法**

```sh
docker compose up -d # データベース起動
go run cmd/basic/main.go # データベースからCSVを出力
docker compose down # データベース停止
```

## **学習ポイント**

1. **`csv.NewWriter()`** を使って CSV ファイルを作成し、**`writer.Flush()`** でデータをファイルに書き込む。

## 作成者

- **池田虎太郎** | [GitHub プロフィール](https://github.com/kotaroikeda-apl-dev)
