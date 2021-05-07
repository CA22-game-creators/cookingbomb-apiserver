# アーキテクチャ

## 概要

DDD(ドメイン駆動設計)を意識した構成になっています。

## 主なレイヤーと責務

### Domain 層

ドメインモデルとその振る舞いを定義する。
ドメイン知識を記述する。

- エンティティ

  同一性によって識別されるドメインオブジェクト

- 値オブジェクト

  同一性によって識別されないドメインオブジェクト

- ファクトリ

  エンティティの新規生成処理を担う

- リポジトリ <I\>

  エンティティの永続化・再構築のインターフェース

- ドメインサービス

  エンティティや値オブジェクトの責務ではないドメインモデルのロジック

### Infrastructure 層

エンティティの永続化・再構築や、DB との通信が責務。

- DBModeler

  エンティティから DB モデルへの詰め替え

- リポジトリ実装クラス

  エンティティの永続化・再構築を行う

### Application 層

ドメイン層の記述を元にユースケースの進行を担う。

- インタラクタ

  ユースケースの進行

- インプットポート <I\>

  インタラクタ のインターフェース

- インプットデータ <DS\>

  インタラクタの入力値

- アウトプットデータ <DS\>

  インタラクタの出力値

### Presentation 層

リクエスト/レスポンスを行う。