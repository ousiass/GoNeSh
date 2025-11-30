---
title: "2. システムアーキテクチャ"
description: "GoNeShのシステム構成、UI設計、Atomic Design"
weight: 3
---

## 2-1. レイヤー構成

| レイヤー | コンポーネント | 役割 |
|---------|--------------|------|
| Presentation | Liquid UI Engine | Bubbletea (TUI) をベースに、テキストとグラフを統合描画。タブ/ペイン管理。ステータスバーは下部固定。 |
| Application | Core Logic | コマンド解析、ウィンドウ管理、ショートカット処理、履歴管理。 |
| Service | Neural Daemon | Gemini/OpenAI API通信、ローカルLLM推論（llama.cpp/ggml）、プリセット管理。 |
| Network | Portal Manager | SSH接続管理(crypto/ssh)、Tailscale連携、Agentlessリソース監視、SCP転送。 |
| System | Host Bridge | クロスプラットフォームAPI連携（ファイルシステム監視、クリップボード、通知）。 |

---

## 2-2. UIデザインコンセプト

**デザインリファレンス:** [Hyprland](https://hyprland.org/)

Hyprlandのモダンでミニマルな美学を参考に、以下の要素を取り入れる：

| 要素 | 説明 |
|------|------|
| **フローティング要素** | 角丸のフローティングメニュー、オーバーレイ |
| **グラデーション & 半透明** | ステータスバー、パネル背景に微細な透過効果 |
| **アクセントカラー** | 重要な情報（Prod警告、AI応答中）にビビッドなアクセント |
| **タイポグラフィ** | モノスペースフォント + クリーンなUI用フォントの組み合わせ |
| **アニメーション** | パネル開閉、タブ切り替えに滑らかなトランジション |
| **ボーダー** | 細いボーダーラインでセクションを区切り、余白を活かす |

---

## 2-3. UI設計パターン: Atomic Design

UIコンポーネントは **Atomic Design** パターンで構築する。

```
ui/
├── atoms/          # 最小単位のUI要素
│   ├── text.go         # テキスト表示
│   ├── icon.go         # アイコン
│   ├── badge.go        # ステータスバッジ
│   └── spinner.go      # ローディング
│
├── molecules/      # Atomsを組み合わせた機能単位
│   ├── tab_item.go     # タブ1つ分
│   ├── resource_meter.go   # CPU/GPU メーター
│   ├── preset_item.go  # プリセット1行
│   └── file_row.go     # ファイル一覧の1行
│
├── organisms/      # Moleculesを組み合わせた独立機能
│   ├── tab_bar.go      # タブバー全体
│   ├── status_bar.go   # ステータスバー全体
│   ├── ai_panel.go     # AIアシスタントパネル
│   ├── file_browser.go # ファイルブラウザ
│   └── floating_menu.go    # フローティングメニュー
│
├── templates/      # ページレイアウト
│   ├── main_layout.go  # メイン画面レイアウト
│   └── split_layout.go # スプリットビュー
│
└── pages/          # 実際の画面
    ├── shell.go        # シェル画面
    └── settings.go     # 設定画面
```

### 2-3-1. 設計原則

| 階層 | 役割 |
|------|------|
| **Atoms** | 単一責任、再利用可能な最小パーツ |
| **Molecules** | Atomsの組み合わせ、1つの機能を担う |
| **Organisms** | 独立して動作可能なUI領域 |
| **Templates** | レイアウト定義（コンテンツなし） |
| **Pages** | 実際のデータを流し込んだ画面 |

---

## 2-4. ディレクトリ構成

```
gonesh/
├── cmd/
│   └── gonesh/
│       └── main.go
├── internal/
│   ├── ui/              # Atomic Design構造
│   ├── core/            # Core Logic
│   ├── neural/          # AI連携
│   ├── portal/          # SSH/Network
│   ├── api/             # API Client（Postman風機能）
│   └── bridge/          # OS連携
├── pkg/
│   └── config/          # 設定読み込み
├── docs/                # ドキュメント
└── test/                # テスト
```
