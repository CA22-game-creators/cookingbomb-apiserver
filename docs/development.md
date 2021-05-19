# ローカル環境構築

## Go のインストール

[goenv](https://qiita.com/koralle/items/7a16772ad1d2e2e34682)がおすすめ

## Docker の準備

Docker を[ここ](https://www.docker.com/get-started)からインストール

## アプリケーションの実行

```
$ make run
```

## CLI で用いる Go ライブラリ

`tools/tools.go`を参照

### wire

- 依存性注入(DI)ツール

```
$ go install github.com/google/wire/cmd/wire
```

### Skeema

- マイグレーションツール

```
$ go install github.com/skeema/skeema
```

### golangci-lint

- コードの静的解析を行う
- Go の Linter 詰め合わせ

```
$ go install github.com/golangci/golangci-lint/cmd/golangci-lint
```

### grpc_cli（任意）

- CLI で gRPC のリクエストができる
- [インストール](https://qiita.com/jackchuka/items/2072191efccec8a2d859)
