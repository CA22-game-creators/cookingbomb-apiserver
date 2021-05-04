# データベースについて

## スキーマの管理

migration ツールには[skeema](https://github.com/skeema/skeema)を採用した

### 特徴

- 現在の DB 状態を sql ファイルに記載でき、差分を積み上げる[golang-migrate](https://github.com/golang-migrate/migrate)などより可視性が高い
- sql ファイルと DB で相互に反映し合う

### 使い方

[インストール](https://www.skeema.io/download/)

- コードと DB の diff を確認する

```
$ skeema diff local -p${MYSQL_ROOT_PASSWORD}
```

- コードを DB に反映する(破壊的なら`--allow-unsafe`を付ける)

```
$ skeema push local -p${MYSQL_ROOT_PASSWORD}
```

- DB をコードに反映する

```
$ skeema pull local -p${MYSQL_ROOT_PASSWORD}
```
