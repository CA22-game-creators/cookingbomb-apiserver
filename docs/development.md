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
# install
$ go install github.com/google/wire/cmd/wire

# DI(account)
$ wire ./di/account/wire.go
```

### Skeema

- マイグレーションツール

```
# install
$ go install github.com/skeema/skeema

# コードと DB の差分を確認する
$ skeema diff local -ppassword

# コードの情報を DB に反映する
$ skeema push local -ppassword

# DB の情報をコードに反映する
$ skeema pull local -ppassword
```

### golangci-lint

- コードの静的解析を行う
- Go の Linter 詰め合わせ

```
# install
$ go install github.com/golangci/golangci-lint/cmd/golangci-lint

# 実行
$ golangci-lint run
```

### gRPCurl

- CLI で gRPC のリクエストができる
- [公式/インストール](https://github.com/fullstorydev/grpcurl)

```
# アカウント登録
$ grpcurl -plaintext -d '{"name": "Mario"}' localhost:8080 proto.AccountServices/Signup
```
