---
title: "8. 品質保証"
description: "GoNeShのテスト方針、CI/CD、コード品質"
weight: 9
---

## 8-1. テスト方針

### 8-1-1. ユニットテスト

各コンポーネントの単体テスト。

```go
// 例: ui/atoms/badge_test.go
func TestBadge_Render(t *testing.T) {
    badge := NewBadge("online", ColorGreen)
    result := badge.View()
    assert.Contains(t, result, "online")
}
```

**カバレッジ目標:** 80%以上

### 8-1-2. 統合テスト

複数コンポーネントの連携テスト。

- SSH接続テスト（モックサーバー使用）
- AI API連携テスト（モックレスポンス使用）
- ファイル転送テスト

### 8-1-3. E2Eテスト

実際のユーザー操作をシミュレート。

- シェル起動→コマンド実行→終了
- SSH接続→リモートコマンド実行→切断
- API Client起動→リクエスト送信→レスポンス確認

---

## 8-2. CI/CD

### 8-2-1. GitHub Actions ワークフロー

```yaml
# .github/workflows/ci.yml

name: CI

on:
  push:
    branches: [main, develop]
  pull_request:
    branches: [main]

jobs:
  test:
    runs-on: ubuntu-latest
    steps:
      - uses: actions/checkout@v4
      - uses: actions/setup-go@v5
        with:
          go-version: '1.22'
      - name: Run tests
        run: go test -v -race -coverprofile=coverage.out ./...
      - name: Upload coverage
        uses: codecov/codecov-action@v4

  lint:
    runs-on: ubuntu-latest
    steps:
      - uses: actions/checkout@v4
      - uses: golangci/golangci-lint-action@v4

  build:
    runs-on: ubuntu-latest
    needs: [test, lint]
    strategy:
      matrix:
        goos: [linux, windows, darwin]
        goarch: [amd64, arm64]
        exclude:
          - goos: windows
            goarch: arm64
    steps:
      - uses: actions/checkout@v4
      - uses: actions/setup-go@v5
        with:
          go-version: '1.22'
      - name: Build
        run: |
          GOOS=${{ matrix.goos }} GOARCH=${{ matrix.goarch }} \
          go build -o gonesh-${{ matrix.goos }}-${{ matrix.goarch }} ./cmd/gonesh
      - uses: actions/upload-artifact@v4
        with:
          name: gonesh-${{ matrix.goos }}-${{ matrix.goarch }}
          path: gonesh-*
```

### 8-2-2. リリースワークフロー

```yaml
# .github/workflows/release.yml

name: Release

on:
  push:
    tags:
      - 'v*'

jobs:
  release:
    runs-on: ubuntu-latest
    steps:
      - uses: actions/checkout@v4
      - uses: actions/setup-go@v5
        with:
          go-version: '1.22'
      - name: Build all platforms
        run: |
          for GOOS in linux darwin windows; do
            for GOARCH in amd64 arm64; do
              if [ "$GOOS" = "windows" ] && [ "$GOARCH" = "arm64" ]; then
                continue
              fi
              output="gonesh-${GOOS}-${GOARCH}"
              if [ "$GOOS" = "windows" ]; then
                output="${output}.exe"
              fi
              GOOS=$GOOS GOARCH=$GOARCH go build -o $output ./cmd/gonesh
            done
          done
      - name: Create Release
        uses: softprops/action-gh-release@v1
        with:
          files: gonesh-*
```

---

## 8-3. コード品質

### 8-3-1. Linter設定

```yaml
# .golangci.yml

linters:
  enable:
    - gofmt
    - goimports
    - govet
    - errcheck
    - staticcheck
    - gosimple
    - ineffassign
    - unused
    - misspell

linters-settings:
  gofmt:
    simplify: true
  goimports:
    local-prefixes: github.com/yourusername/gonesh
```

### 8-3-2. Pre-commit Hooks

```yaml
# .pre-commit-config.yaml

repos:
  - repo: local
    hooks:
      - id: go-fmt
        name: go fmt
        entry: gofmt -w
        language: system
        types: [go]
      - id: go-vet
        name: go vet
        entry: go vet ./...
        language: system
        types: [go]
      - id: go-test
        name: go test
        entry: go test ./...
        language: system
        types: [go]
```

---

## 8-4. ドキュメント

### 8-4-1. コードコメント

```go
// Package ui provides the user interface components following Atomic Design pattern.
package ui

// Badge represents a status indicator with color.
// It is an Atom-level component.
type Badge struct {
    text  string
    color Color
}

// NewBadge creates a new Badge with the given text and color.
func NewBadge(text string, color Color) *Badge {
    return &Badge{text: text, color: color}
}
```

### 8-4-2. CHANGELOG

```markdown
# Changelog

## [Unreleased]

## [0.1.0] - YYYY-MM-DD
### Added
- Initial release
- Basic shell functionality
- Tab management
- Status bar with local resources
```

---

## 8-5. セキュリティ

### 8-5-1. 機密情報の取り扱い

- APIキーは環境変数で管理（設定ファイルに直接記載しない）
- SSH秘密鍵のパスのみ設定ファイルに記載
- 履歴にパスワードを含むコマンドは保存しない

### 8-5-2. 依存関係の監査

```bash
# 脆弱性チェック
go list -m all | nancy sleuth

# または
govulncheck ./...
```
