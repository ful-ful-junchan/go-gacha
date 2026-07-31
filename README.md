# go-gacha

ガチャ(ソーシャルゲームのランダム抽選)の仕組みを模した学習用のGo製Web APIです。レアリティごとの重み付き抽選、天井(一定回数はずれるとSSR確定)、抽選履歴の保存を備えています。

## 技術スタック

- Go 1.22
- MySQL 8.0
- Docker / Docker Compose

## 機能

### ガチャ情報取得

```
GET /gachas/{gacha_id}
```

指定したガチャの基本情報(名前・天井回数など)を返します。

### ガチャを引く

```
POST /gachas/{gacha_id}/draw
Content-Type: application/json

{
  "user_id": 1,
  "count": 10
}
```

`count`は`1`または`10`のみ指定可能です。レアリティ(SSR/SR/R)は`gacha_items.weight`の重みに基づいて抽選され、天井回数(`gachas.pity_threshold`)に到達するとSSRが確定します。抽選結果は`gacha_histories`に保存され、天井カウントは`user_pity_counters`に保存されます。

## セットアップ

1. `.env`を作成し、以下を設定します。

   ```
   MYSQL_ROOT_PASSWORD=xxxx
   MYSQL_DATABASE=gacha_db
   MYSQL_USER=xxxx
   MYSQL_PASSWORD=xxxx
   ```

2. コンテナを起動します。

   ```
   make up
   ```

   初回起動時、`docker/mysql/init.sql`によってテーブルが自動作成されます。

3. 動作確認します。

   ```
   curl http://localhost:8080/gachas/1
   ```

## 開発用コマンド(Makefile)

| コマンド | 内容 |
| --- | --- |
| `make build` | イメージをビルド |
| `make up` | コンテナを起動 |
| `make down` | コンテナを削除して停止 |
| `make stop` | コンテナを停止(削除はしない) |
| `make restart` | 全コンテナを再起動 |
| `make go` | ローカルでビルドチェックした上で`api`コンテナのみ再起動 |

`api`コンテナは`go run`でソースを実行しているだけでホットリロードは無いため、コードを変更したら`make go`(または`docker compose restart api`)でコンテナを再起動しないと変更が反映されません。

## DBスキーマ概要

| テーブル | 役割 |
| --- | --- |
| `gachas` | ガチャの基本情報(開催期間・天井回数) |
| `items` | 排出アイテムのマスタ(レアリティ含む) |
| `gacha_items` | ガチャとアイテムの中間テーブル。`weight`が排出重み |
| `gacha_histories` | ユーザーの抽選履歴 |
| `user_pity_counters` | ユーザー・ガチャごとの天井カウント |

## ディレクトリ構成

```
cmd/api            エントリポイント(main.go)
internal/handler    HTTPハンドラ
internal/service     ビジネスロジック(抽選・天井計算など)
internal/repository  DBアクセス
internal/model       ドメインモデル
docker/mysql/init.sql  DB初期化スキーマ
```

