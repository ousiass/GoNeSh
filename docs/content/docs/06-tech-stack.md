---
title: "6. 推奨技術スタック"
description: "GoNeShで使用するGoライブラリとビルド設定"
weight: 7
---

## 6-1. コアライブラリ

| カテゴリ | ライブラリ | 役割 |
|---------|----------|------|
| CLI/TUI | `github.com/charmbracelet/bubbletea` | UI構築、イベントループ管理 |
| Layout/Style | `github.com/charmbracelet/lipgloss` | スタイリング |
| Input | `github.com/charmbracelet/bubbles` | 入力コンポーネント |

---

## 6-2. ネットワーク

| カテゴリ | ライブラリ | 役割 |
|---------|----------|------|
| SSH | `golang.org/x/crypto/ssh` | SSHクライアント |
| HTTP | `net/http` (標準) | HTTPクライアント |
| HTTP Client | `github.com/go-resty/resty/v2` | REST APIクライアント |

---

## 6-3. AI連携

| カテゴリ | ライブラリ | 役割 |
|---------|----------|------|
| Gemini | `github.com/google/generative-ai-go` | Gemini API |
| OpenAI | `github.com/sashabaranov/go-openai` | OpenAI API |
| Local LLM | `github.com/go-skynet/go-llama.cpp` | ローカルLLM |

---

## 6-4. システム

| カテゴリ | ライブラリ | 役割 |
|---------|----------|------|
| System Info | `github.com/shirou/gopsutil` | リソース情報取得 |
| Process | `github.com/creack/pty` | PTY制御 |
| File Watch | `github.com/fsnotify/fsnotify` | ファイル監視 |

---

## 6-5. 設定・データ

| カテゴリ | ライブラリ | 役割 |
|---------|----------|------|
| Config | `github.com/spf13/viper` | YAML読み込み |
| CLI Flags | `github.com/spf13/cobra` | CLIフラグ解析 |
| JSON | `encoding/json` (標準) | JSON処理 |

---

## 6-6. API Client機能

| カテゴリ | ライブラリ | 役割 |
|---------|----------|------|
| OpenAPI Parser | `github.com/getkin/kin-openapi` | Swagger/OpenAPI解析 |
| HTTP Mock | `github.com/jarcoal/httpmock` | モックサーバー |
| Template | `text/template` (標準) | 環境変数展開 |

---

## 6-7. テスト

| カテゴリ | ライブラリ | 役割 |
|---------|----------|------|
| Testing | `testing` (標準) | ユニットテスト |
| Assertion | `github.com/stretchr/testify` | アサーション |
| Mock | `github.com/golang/mock` | モック生成 |

---

## 6-8. Go バージョン

```
go 1.22+
```

---

## 6-9. ビルド設定

```go
// go.mod
module github.com/yourusername/gonesh

go 1.22

require (
    github.com/charmbracelet/bubbletea v0.25.0
    github.com/charmbracelet/lipgloss v0.9.1
    github.com/charmbracelet/bubbles v0.18.0
    golang.org/x/crypto v0.21.0
    github.com/google/generative-ai-go v0.10.0
    github.com/sashabaranov/go-openai v1.20.0
    github.com/shirou/gopsutil/v3 v3.24.0
    github.com/spf13/viper v1.18.0
    github.com/spf13/cobra v1.8.0
    github.com/creack/pty v1.1.21
    github.com/getkin/kin-openapi v0.123.0
    github.com/go-resty/resty/v2 v2.11.0
)
```

---

## 6-10. クロスコンパイル

```bash
# Windows
GOOS=windows GOARCH=amd64 go build -o gonesh.exe ./cmd/gonesh

# Linux
GOOS=linux GOARCH=amd64 go build -o gonesh ./cmd/gonesh

# macOS (Intel)
GOOS=darwin GOARCH=amd64 go build -o gonesh-darwin ./cmd/gonesh

# macOS (Apple Silicon)
GOOS=darwin GOARCH=arm64 go build -o gonesh-darwin-arm64 ./cmd/gonesh
```
