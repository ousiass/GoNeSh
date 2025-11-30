---
title: "5. データ構造定義"
description: "GoNeShの設定ファイルとYAML定義"
weight: 6
---

## 5-1. 設定ファイル一覧

| ファイル | 用途 |
|----------|------|
| `~/.gonesh/presets.yaml` | AIプリセット定義 |
| `~/.gonesh/connections.yaml` | SSH接続先設定 |
| `~/.gonesh/transfers.yaml` | Quick Transfer設定 |
| `~/.gonesh/api-specs.yaml` | Swagger/OpenAPI仕様パス |
| `~/.gonesh/api-envs.yaml` | API環境変数 |
| `~/.gonesh/api-collections.yaml` | APIリクエストコレクション |
| `~/.gonesh/ai-tools.yaml` | 外部AIツール設定 |
| `~/.gonesh/git.yaml` | Git Auto Commit設定 |
| `~/.gonesh/history` | コマンド履歴 |
| `~/.gonesh/api-history.json` | APIリクエスト履歴 |

---

## 5-2. AIプリセット設定

```yaml
# ~/.gonesh/presets.yaml

presets:
  - name: "Code Review Expert"
    trigger: "/review"
    model: "gemini-1.5-pro"
    system_prompt: |
      あなたはGoogleとUberでシニアエンジニアを務めた経験のあるGo言語のスペシャリストです。
      以下の観点でコードレビューを行ってください：
      1. パフォーマンス（メモリアロケーションの効率化）
      2. 可読性（Effective Goに準拠しているか）
      3. エラーハンドリングの適切さ

      回答は日本語で、具体的な修正コードブロックを含めてください。
    context: "selection"

  - name: "Git Commit Generator"
    trigger: "/commit"
    model: "gemini-1.5-flash"
    system_prompt: |
      git diffの出力結果を解析し、Conventional Commits (feat, fix, refactor等) に従ったコミットメッセージ案を3つ作成してください。
      出力形式：
      - <type>: <description>
    command: "git diff --cached"

  - name: "Explain Error"
    trigger: "/explain"
    model: "gpt-4o"
    system_prompt: |
      発生したエラーログを初心者にわかりやすく解説し、解決策をステップバイステップで提示してください。
    context: "last_output"

  - name: "API Request Generator"
    trigger: "/api-gen"
    model: "gemini-1.5-pro"
    system_prompt: |
      指定された要件に基づいてAPIリクエストのJSONボディを生成してください。
      バリデーションエラーのテストケースも含めてください。
    context: "selection"
```

---

## 5-3. SSH接続設定

```yaml
# ~/.gonesh/connections.yaml

connections:
  - name: "dev"
    host: "192.168.1.100"
    user: "ubuntu"
    key: "~/.ssh/id_ed25519"
    env: "dev"

  - name: "gpu"
    host: "gpu-server"  # Tailscale MagicDNS名
    user: "admin"
    env: "dev"

  - name: "prod"
    host: "api-prod"
    user: "deploy"
    env: "prod"  # 赤みのラインで警告表示
```

---

## 5-4. Quick Transfer設定

```yaml
# ~/.gonesh/transfers.yaml

transfers:
  - name: "dev"
    connection: "dev"
    remote_path: "/home/ubuntu/projects/"
    local_path: "~/dev/"

  - name: "logs"
    connection: "prod"
    remote_path: "/var/log/app/"
    local_path: "~/logs/"
```

---

## 5-5. API仕様設定

```yaml
# ~/.gonesh/api-specs.yaml

specs:
  - name: "Local Backend"
    path: "./swagger.yaml"
    auto_load: true

  - name: "External API"
    url: "https://api.example.com/openapi.json"
    auto_load: false
```

---

## 5-6. API環境変数

```yaml
# ~/.gonesh/api-envs.yaml

default: "local"

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

---

## 5-7. APIコレクション

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
        headers:
          Authorization: "Bearer {{token}}"

      - name: "Update User"
        method: "PUT"
        url: "{{base_url}}/api/users/{{user_id}}"
        headers:
          Content-Type: "application/json"
          Authorization: "Bearer {{token}}"
        body: |
          {
            "name": "updated user"
          }
```

---

## 5-8. 外部AIツール設定

```yaml
# ~/.gonesh/ai-tools.yaml

# デフォルトで使用するツール
default: "Claude Code"

tools:
  - name: "Claude Code"
    command: "claude"
    shortcut: "c"
    args: ["--print"]
    description: "Anthropic公式CLI"

  - name: "Cursor"
    command: "cursor"
    shortcut: "u"
    args: ["--goto"]
    description: "AI搭載エディタ"

  - name: "Cline"
    command: "code"
    shortcut: "l"
    args: ["--goto"]
    description: "VSCode + Cline拡張"

  - name: "Aider"
    command: "aider"
    shortcut: "d"
    args: ["--message"]
    description: "ターミナルAIペアプログラマー"

  - name: "Custom"
    command: "gpt"
    shortcut: "g"
    args: ["-m", "gpt-4o"]
    description: "カスタムGPT CLI"

# コンテキスト送信時のオプション
context:
  include_file_path: true      # ファイルパスを含める
  include_line_numbers: true   # 行番号を含める
  include_git_diff: false      # git diffを含める
  max_context_lines: 500       # 最大行数

# AIオーケストレーション設定
ai_orchestration:
  enabled: true                 # AIからのツール起動を有効化
  auto_launch: true             # AIの判断で自動起動
  confirm_before_launch: true   # 起動前に確認ダイアログ

  # タスクと推奨ツールのマッピング
  task_routing:
    refactor: "Claude Code"     # リファクタリング → Claude Code
    debug: "Claude Code"        # デバッグ → Claude Code
    test: "Aider"               # テスト作成 → Aider
    ui: "Cursor"                # UI修正 → Cursor
    explain: "internal"         # 説明 → 内蔵AIで処理
    review: "internal"          # レビュー → 内蔵AIで処理
```

---

## 5-9. Git Auto Commit設定

```yaml
# ~/.gonesh/git.yaml

auto_commit:
  enabled: true
  model: "gemini-1.5-flash"    # 高速モデルを使用
  language: "ja"               # メッセージの言語 (ja/en)
  emoji: true                  # 絵文字を付ける
  candidates: 5                # 生成する候補数
  max_diff_lines: 500          # 解析する最大行数
  auto_stage: false            # 自動ステージング

  # コミットメッセージのフォーマット
  format: "<type>: <description> <emoji>"

  # カスタム絵文字マッピング
  emoji_map:
    feat: "✨"
    fix: "🐛"
    update: "📝"
    refactor: "♻️"
    style: "💄"
    test: "✅"
    docs: "📚"
    chore: "📦"
    perf: "⚡"
    security: "🔒"
    wip: "🚧"
    hotfix: "🚑"

  # 除外パターン（これらのファイルは解析しない）
  exclude:
    - "*.lock"
    - "node_modules/*"
    - "vendor/*"
    - ".env*"

  # コミット前フック
  hooks:
    pre_commit:
      - "go fmt ./..."
      - "go vet ./..."
```
