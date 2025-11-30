---
title: "4. API Client機能"
description: "GoNeShのPostman風API Client機能"
weight: 5
---

## 4-1. 概要

GoNeSh内蔵のAPIテストクライアント。Swagger/OpenAPI仕様を読み込み、ターミナル上でAPIリクエストの送信・レスポンス確認ができる。

**ショートカット:** `g + r` で起動

---

## 4-2. 画面構成

```
+-----------------------------------------------------------------------+
| [API Client] POST /api/users                                    [×]  |
+----------------------------------+------------------------------------+
| Endpoints          | Request                                         |
| ─────────────────  | ───────────────────────────────────────────────  |
| ▼ Users            | Method: [POST ▼]                                |
|   GET  /users      | URL: http://localhost:3000/api/users            |
|   POST /users      | ───────────────────────────────────────────────  |
|   GET  /users/{id} | Headers:                                        |
| ▼ Products         |   Content-Type: application/json                |
|   GET  /products   |   Authorization: Bearer {{token}}               |
|                    | ───────────────────────────────────────────────  |
|                    | Body:                                           |
|                    | {                                               |
|                    |   "name": "test user",                          |
|                    |   "email": "test@example.com"                   |
|                    | }                                               |
+----------------------------------+------------------------------------+
| Response (200 OK - 145ms)                                            |
| ──────────────────────────────────────────────────────────────────── |
| {                                                                    |
|   "id": 1,                                                           |
|   "name": "test user",                                               |
|   "email": "test@example.com"                                        |
| }                                                                    |
+-----------------------------------------------------------------------+
| [Send Request: Enter] [Save: g+s] [History: g+h] [Close: Esc]        |
+-----------------------------------------------------------------------+
```

---

## 4-3. 機能一覧

### 4-3-1. Swagger/OpenAPI読み込み

```bash
# コマンドで読み込み
api load ./swagger.yaml
api load https://api.example.com/openapi.json

# 設定ファイルから自動読み込み
# ~/.gonesh/api-specs.yaml
```

読み込み後、エンドポイント一覧が左ペインにツリー表示される。

### 4-3-2. リクエスト送信

| 機能 | 説明 |
|------|------|
| **HTTP メソッド** | GET, POST, PUT, PATCH, DELETE, HEAD, OPTIONS |
| **Headers** | カスタムヘッダー設定（認証トークン等） |
| **Body** | JSON / Form Data / Raw テキスト |
| **Query Params** | URLパラメータ |
| **Path Params** | `/users/{id}` の `{id}` を自動検出 |

### 4-3-3. 環境変数

```yaml
# ~/.gonesh/api-envs.yaml

environments:
  - name: "local"
    variables:
      base_url: "http://localhost:3000"
      token: "dev-token-xxx"

  - name: "staging"
    variables:
      base_url: "https://staging-api.example.com"
      token: "staging-token-xxx"

  - name: "production"
    variables:
      base_url: "https://api.example.com"
      token: "prod-token-xxx"
```

リクエスト内で `{{base_url}}` や `{{token}}` として参照可能。

### 4-3-4. リクエスト履歴

- 直近のリクエスト履歴を自動保存
- `g + h` で履歴一覧を表示
- 履歴から再実行可能

### 4-3-5. コレクション保存

```yaml
# ~/.gonesh/api-collections.yaml

collections:
  - name: "User API Tests"
    requests:
      - name: "Create User"
        method: "POST"
        url: "{{base_url}}/api/users"
        headers:
          Content-Type: "application/json"
          Authorization: "Bearer {{token}}"
        body: |
          {
            "name": "test user",
            "email": "test@example.com"
          }

      - name: "Get User"
        method: "GET"
        url: "{{base_url}}/api/users/{{user_id}}"
```

---

## 4-4. キーバインド（API Client内）

| 操作 | キーバインド |
|------|------------|
| リクエスト送信 | `Enter` |
| 環境切り替え | `g + e` |
| 履歴表示 | `g + h` |
| コレクション保存 | `g + s` |
| レスポンスをコピー | `g + c` |
| 閉じる | `Esc` |
| エンドポイント検索 | `/` |

---

## 4-5. ローカルサーバー連携

### 4-5-1. Mock Server起動

Swagger/OpenAPI仕様からモックサーバーを自動起動。

```bash
# コマンドで起動
api mock ./swagger.yaml --port 8080

# レスポンス例を自動生成してモックとして返す
```

### 4-5-2. Proxy Mode

既存のAPIサーバーへのリクエストをプロキシし、リクエスト/レスポンスをログ。

```bash
# プロキシ起動
api proxy https://api.example.com --port 9000

# localhost:9000 へのリクエストが api.example.com に転送される
# 全てのリクエスト/レスポンスが履歴に記録される
```

---

## 4-6. AI連携

### 4-6-1. リクエスト生成

AIにリクエストボディを生成させる。

```
> /api-gen Create a request body for user registration with validation errors
```

AIパネルが生成したJSONをそのままBodyにペースト可能。

### 4-6-2. レスポンス解析

レスポンスを選択して `g + a` でAIに解析させる。

```
> このレスポンスのエラーの原因を教えて
```

---

## 4-7. 設定ファイル

| ファイル | 用途 |
|----------|------|
| `~/.gonesh/api-specs.yaml` | Swagger/OpenAPI仕様のパス一覧 |
| `~/.gonesh/api-envs.yaml` | 環境変数定義 |
| `~/.gonesh/api-collections.yaml` | 保存したリクエストコレクション |
| `~/.gonesh/api-history.json` | リクエスト履歴 |
