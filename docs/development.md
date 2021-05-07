# ローカル環境構築

## Go のインストール

[goenv](https://qiita.com/koralle/items/7a16772ad1d2e2e34682)がおすすめ

## Docker の準備

Docker を[ここ](https://www.docker.com/get-started)からインストール

## アプリケーションの実行

```
$ make run
```

## go.mod とは別に インストールが必要な Go ライブラリ

### wire

- 依存性注入(DI)ツール
- [インストール](https://github.com/google/wire)

### Skeema

- マイグレーションツール
- [インストール](https://www.skeema.io/download/)

### grpc_cli

- CLI で gRPC のリクエストができる
- [インストール](https://qiita.com/jackchuka/items/2072191efccec8a2d859)

### golangci-lint

- コードの静的解析を行う
- Go の Linter 詰め合わせ
- [インストール](https://golangci-lint.run/usage/install/#local-installation)
