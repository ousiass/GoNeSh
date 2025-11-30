---
title: "7. 開発ロードマップ"
description: "GoNeShの開発ステップとマイルストーン"
weight: 8
---

## 7-1. ステップ一覧

| Step | 名称 | 内容 |
|------|------|------|
| **1** | The Shell Foundation | 基礎UI構築 |
| **2** | Local Command Execution | ローカルシェル機能 |
| **3** | Portal Implementation | SSH/リモート機能 |
| **4** | AI Integration | AI連携機能 |
| **5** | API Client | Postman風API機能 |
| **6** | Advanced Features | 高度な機能 |

---

## 7-2. Step 1: The Shell Foundation

**目標:** Bubbletea でステータスバー付きシェル画面。タブUIの基礎実装。

### 7-2-1. タスク

- [ ] プロジェクト初期化（Go modules, ディレクトリ構成）
- [ ] Bubbletea基本セットアップ
- [ ] Atomic Designに基づくUI構造の構築
- [ ] ステータスバー実装（CPU/メモリ表示）
- [ ] タブバー実装（複数タブの切り替え）
- [ ] `g+キー` ショートカットシステム実装
- [ ] 設定ファイル読み込み（viper）

### 7-2-2. 成果物

- 基本的なTUI画面が表示される
- タブ切り替えができる
- ステータスバーにローカルリソースが表示される

---

## 7-3. Step 2: Local Command Execution

**目標:** pty統合、ローカルコマンド実行、履歴の永続化。

### 7-3-1. タスク

- [ ] PTY統合（creack/pty）
- [ ] コマンド入力・実行
- [ ] 出力のスクロール表示
- [ ] コマンド履歴の保存・読み込み
- [ ] `Ctrl+R` 履歴検索
- [ ] クリップボード連携

### 7-3-2. 成果物

- ローカルコマンドが実行できる
- 履歴が永続化される
- 履歴検索ができる

---

## 7-4. Step 3: Portal Implementation

**目標:** SSH接続、リモートリソース取得、Tailscale連携、Quick Transfer。

### 7-4-1. タスク

- [ ] SSH接続（golang.org/x/crypto/ssh）
- [ ] connections.yaml読み込み
- [ ] 秘密鍵認証
- [ ] Tailscale MagicDNS対応
- [ ] リモートリソース取得（Agentless）
- [ ] Smart Context Switching（Prod/Dev色分け）
- [ ] Quick Transfer（push/pull）実装
- [ ] transfers.yaml読み込み

### 7-4-2. 成果物

- SSH接続ができる
- リモートリソースがステータスバーに表示される
- `push`/`pull` でファイル転送ができる

---

## 7-5. Step 4: AI Integration

**目標:** Gemini/OpenAI連携、プリセット、AIパネル、フローティングメニュー。

### 7-5-1. タスク

- [ ] Gemini API連携
- [ ] OpenAI API連携
- [ ] presets.yaml読み込み
- [ ] AIパネル（右側スプリット）実装
- [ ] フローティングメニュー実装
- [ ] `g+a` でコンテキスト転送
- [ ] `g+p` でプリセット選択
- [ ] スラッシュコマンド（/review等）

### 7-5-2. 成果物

- AIに質問できる
- プリセットが使える
- コード選択→AI転送ができる

---

## 7-6. Step 5: API Client

**目標:** Postman風API機能。Swagger読み込み、リクエスト送信、環境変数。

### 7-6-1. タスク

- [ ] API Client画面実装
- [ ] Swagger/OpenAPI解析（kin-openapi）
- [ ] エンドポイント一覧表示
- [ ] リクエスト作成・送信
- [ ] レスポンス表示
- [ ] 環境変数（api-envs.yaml）
- [ ] コレクション保存（api-collections.yaml）
- [ ] リクエスト履歴
- [ ] `g+r` でAPI Client起動

### 7-6-2. 成果物

- Swaggerを読み込んでAPIテストができる
- 環境を切り替えられる
- リクエストを保存・再利用できる

---

## 7-7. Step 6: Advanced Features

**目標:** ローカルLLM、ビジュアルファイルブラウザ、ペイン分割、多段SSH。

### 7-7-1. タスク

- [ ] ローカルLLM対応（llama.cpp）
- [ ] ビジュアルファイルブラウザ（2ペイン）
- [ ] ドラッグ&ドロップ対応
- [ ] ペイン分割
- [ ] 多段SSH（ProxyJump）
- [ ] Mock Server起動
- [ ] Proxy Mode

### 7-7-2. 成果物

- オフラインでAIが使える
- GUIでファイル転送ができる
- 踏み台経由のSSHができる

---

## 7-8. マイルストーン

| マイルストーン | 含まれるStep | 状態 |
|--------------|-------------|------|
| **v0.1.0** - MVP | Step 1-2 | 計画中 |
| **v0.2.0** - Portal | Step 3 | 計画中 |
| **v0.3.0** - AI | Step 4 | 計画中 |
| **v0.4.0** - API Client | Step 5 | 計画中 |
| **v1.0.0** - Full Release | Step 6 | 計画中 |
