# entdemo
[このzennの記事](https://zenn.dev/tkb/articles/d1e6e3b7d62051)で説明している内容のリポジトリです。

## 1. チュートリアルが完了した状態で始める場合
以下のコマンドでpostgresqlを起動して下さい。

```bash
docker-compose up -d
```

そのあと、`go run start.go`を実行してauto migrationを行なって下さい。<br>
これで、`ent/schema/`で定義されている`Car`と`Group`と`User`のスキーマに対応するテーブルが構築されます。

## 2. チュートリアルを最初から開始する場合
本リポジトリはチュートリアル終了の状態なので、最初から始める場合は、以下の手順に従って下さい。

### 2-1. postgresqlの実行
以下のコマンドでpostgresqlを起動して下さい。

```bash
docker-compose up -d
```

### 2-2. `ent/`の削除
`entdemo`のルートディレクトリで、以下のコマンドを実行して下さい。
```bash
rm -rf ent
```

### 2-3. チュートリアルを開始する
[zennの記事に記載している4-3節](https://zenn.dev/tkb/articles/d1e6e3b7d62051#4-3.-%E3%81%AF%E3%81%98%E3%82%81%E3%81%A6%E3%81%AE%E3%82%B9%E3%82%AD%E3%83%BC%E3%83%9E%E4%BD%9C%E6%88%90)
から手順に従って開始して下さい。
