---
title: "3. 機能要件仕様"
description: "GoNeShの機能要件（UI、キーバインド、SSH、AI）"
weight: 4
---

## 3-1. インターフェース (The Bottom-Heavy Cockpit)

視線移動を最小限にするため、ステータス情報は全て画面下部に集約する。

### 3-1-1. 画面構成

```
+-----------------------------------------------------------------------+
| [Tab1: local] [Tab2: server-a] [Tab3: server-b]                       |
+--------------------------------------------------+--------------------+
| user@local ~ $ ssh ai-training-server            | [AI Assistant]     |
| > Connected to 192.168.1.100                     |                    |
|                                                  | /review を実行中... |
| user@remote ~ $ python train.py --batch-size 64  |                    |
| [Epoch 1/10] Loss: 0.452 ...                     | このコードには以下の |
|                                                  | 問題があります：    |
+--------------------------------------------------+--------------------+
| [LOCAL: CPU 12% | GPU 2%] | [REMOTE: CPU 88% | GPU 98%] | /review    |
+-----------------------------------------------------------------------+
```

### 3-1-2. 領域説明

| 領域 | 説明 |
|------|------|
| **Tab Bar (最上部)** | 複数セッションをタブで管理 |
| **Main Terminal (左側)** | 通常のコマンドライン領域 |
| **AI Assistant Panel (右側)** | スプリット表示。プリセット実行結果を表示 |
| **Unified Status Bar (最下部)** | ローカルリソース / リモートリソース / AIプリセット名 |

---

## 3-2. キーバインド

**GoNeSh独自ショートカット:** `g` + キー（同時押し）

> `g` キーを押しながら他のキーを押すことで、GoNeSh特有のショートカットを実行。
> Ctrl+C のような感覚で `g+a` のように使用する。

### 3-2-1. ショートカット一覧

| 操作 | キーバインド | 覚え方 |
|------|------------|--------|
| 新規タブ | `g + t` | **t**ab |
| タブ切り替え（次） | `g + ]` | 右へ |
| タブ切り替え（前） | `g + [` | 左へ |
| タブを閉じる | `g + w` | close **w**indow |
| AIパネル表示/非表示 | `g + a` | **a**i |
| プリセット選択 | `g + p` | **p**reset |
| ファイルブラウザ | `g + f` | **f**ile |
| Quick Transfer | `g + s` | **s**end |
| API Client | `g + r` | **r**equest |
| Claude Code送信 | `g + c` | **c**laude |
| 外部AIツール選択 | `g + x` | e**x**ternal |
| Git Auto Commit | `g + g` | **g**it |
| 履歴検索 | `Ctrl + R` | (標準) |

### 3-2-2. プリセット選択UI

フローティングメニュー（ターミナルの上に重ねて表示）

---

## 3-3. SSH & リモート管理機能 (Portal Mode)

### 3-3-1. 認証方式

- 秘密鍵認証（~/.ssh/id_rsa, id_ed25519 等）
- Tailscale連携（~/.ssh/config 風の設定スキーム）

### 3-3-2. Agentless Dual Monitor

- SSH接続中、ローカル側から定期的に軽量コマンド（nvidia-smi等）を送り、その結果を下部バーにリアルタイム表示。
- リモートサーバーへのエージェントインストールは不要。

### 3-3-3. Enhanced SCP

- ターミナルウィンドウへのファイルドラッグ&ドロップでアップロード（ドロップ時にファイルブラウザ自動起動）。
- ビジュアルファイルブラウザ（2ペインでローカル↔リモート対照表示）。
- `g + f` でいつでも起動可能。
- 転送進捗のリアルタイム表示。

### 3-3-4. Quick Transfer（独自コマンド）

パス・接続先・ユーザー名を保存し、ローカルターミナルから簡単に転送実行。
**接続先は必ず明示する（引数なしはエラー）。**

```bash
# 使用例
push myfile.txt @dev      # → dev:/home/ubuntu/projects/myfile.txt
pull output.csv @dev      # ← dev:/home/ubuntu/projects/output.csv
pull app.log @logs        # ← prod:/var/log/app/app.log

# 引数なしはエラー
push myfile.txt           # → Error: 接続先を指定してください（例: @dev）
```

### 3-3-5. Smart Context Switching

- 接続先（Prod/Dev）に応じて、下部バーの背景色を微細に変更し、誤操作を防ぐ。
- Prod接続時は赤みのラインを表示。

### 3-3-6. Multi-Session Management

- タブ/ペイン分割で複数SSHセッションを並行管理。
- セッション間でのコピー&ペースト対応。

### 3-3-7. 将来拡張

- 多段SSH（ProxyJump）対応

---

## 3-4. AI・自動補完機能

### 3-4-1. 対応AIプロバイダー

| プロバイダー | 実装タイミング |
|-------------|--------------|
| Gemini API（Google） | Step 4（初期実装の主軸） |
| OpenAI API（GPT-4等） | Step 4 |
| ローカルLLM（llama.cpp/ggml） | Step 5 |

### 3-4-2. APIキー管理

```bash
# 環境変数で管理
export GEMINI_API_KEY="your-key"
export OPENAI_API_KEY="your-key"
```

### 3-4-3. AI選択方式

- プリセットベースで選択（モデルはプリセットに紐づく）
- `g + p` でフローティングメニューからプリセット選択

### 3-4-4. Native AI Integration

- ターミナル内のテキストを選択して `g + a` でAIパネルへコンテキスト転送。
- 過去のコマンド履歴と現在のディレクトリ構成を考慮した「予知補完」。

### 3-4-5. Preset System

- 複雑な指示を事前定義し、スラッシュコマンド（例: /refactor）で呼び出し。
- 各プリセットにモデル（Gemini/OpenAI等）を指定可能。

### 3-4-6. 外部AIツール連携

Claude Code、Cursor、Cline等の外部AIコーディングツールとシームレスに連携。

#### Claude Code連携

```bash
# ショートカットで即座に送信
g + c  # → 選択テキスト or 現在のコンテキストをClaude Codeに送信

# コマンドで送信
cc "このエラーを修正して"           # テキスト指示を送信
cc -f main.go "リファクタリングして"  # ファイルを指定して送信
cc -d ./src "テスト書いて"          # ディレクトリを指定して送信
```

#### 外部ツール選択メニュー

`g + x` でフローティングメニューを表示し、送信先を選択。

| ツール | 説明 |
|--------|------|
| Claude Code | Anthropic公式CLI |
| Cursor | AI搭載エディタ |
| Cline | VSCode AI拡張 |
| Aider | ターミナルAIペアプログラマー |
| カスタム | ユーザー定義のコマンド |

#### AIからのツール起動

GoNeSh内蔵AIが状況を判断し、適切な外部ツールを自動起動。

```
ユーザー: このファイルリファクタリングして
AI: コードの複雑さを考慮し、Claude Codeに依頼します。
    [Claude Code を起動中...]
```

**AIが起動を判断するケース:**

| 状況 | 起動ツール | 理由 |
|------|-----------|------|
| 大規模リファクタリング | Claude Code | ファイル横断の変更が必要 |
| バグ修正依頼 | Aider | git連携が強い |
| UI修正 | Cursor | ビジュアル確認が必要 |
| 簡単な質問 | 内蔵AI | 外部ツール不要 |

**AIコマンドでの明示的指定:**

```bash
# AIパネルから指示
> Claude Codeでこのエラー直して
> Aiderでテスト書いて
> Cursorでこのコンポーネント開いて
```

**自動起動の設定:**

```yaml
# ~/.gonesh/ai-tools.yaml

ai_orchestration:
  enabled: true
  auto_launch: true          # AIの判断で自動起動
  confirm_before_launch: true # 起動前に確認ダイアログ

  # タスクと推奨ツールのマッピング
  task_routing:
    refactor: "Claude Code"
    debug: "Claude Code"
    test: "Aider"
    ui: "Cursor"
    explain: "internal"      # 内蔵AIで処理
```

#### 設定ファイル

```yaml
# ~/.gonesh/ai-tools.yaml

tools:
  - name: "Claude Code"
    command: "claude"
    shortcut: "c"
    args: ["--print"]

  - name: "Cursor"
    command: "cursor"
    args: ["--goto"]

  - name: "Aider"
    command: "aider"
    args: ["--message"]

  - name: "Custom GPT CLI"
    command: "gpt"
    args: ["-m", "gpt-4o"]

default: "Claude Code"
```

---

## 3-5. Git連携機能

### 3-5-1. Git Auto Commit

AIが変更内容を解析し、コミットメッセージを自動生成。

```bash
# ショートカットで実行
g + g  # → git diffを解析してコミットメッセージ生成

# コマンドで実行
gcommit                      # ステージング済みの変更をコミット
gcommit -a                   # 全変更をステージング＆コミット
gcommit --dry-run            # メッセージのみ生成（コミットしない）

# ディレクトリ・ファイル指定
gcommit src/auth/            # 特定ディレクトリのみ
gcommit src/auth/login.go    # 特定ファイルのみ
gcommit src/auth/ src/api/   # 複数指定
gcommit --exclude="*.test.go" # パターン除外
```

**生成されるメッセージ形式:**

```
<type>: <description> <emoji>

例:
feat: ユーザー認証機能を追加 ✨
fix: ログイン時のNullPointerExceptionを修正 🐛
update: READMEにインストール手順を追記 📝
refactor: 認証ロジックをAuthServiceに分離 ♻️
style: コードフォーマットを統一 💄
test: ユーザーAPIのユニットテストを追加 ✅
chore: 依存関係を更新 📦
```

**コミットタイプと絵文字:**

| タイプ | 説明 | 絵文字 |
|--------|------|--------|
| feat | 新機能 | ✨ |
| fix | バグ修正 | 🐛 |
| update | 機能改善・更新 | 📝 |
| refactor | リファクタリング | ♻️ |
| style | フォーマット・スタイル | 💄 |
| test | テスト追加・修正 | ✅ |
| docs | ドキュメント | 📚 |
| chore | ビルド・設定変更 | 📦 |
| perf | パフォーマンス改善 | ⚡ |
| security | セキュリティ修正 | 🔒 |

### 3-5-2. コミットフロー

AIが5つの候補を生成し、選択する形式。

```
┌─────────────────────────────────────────────────────────────────┐
│  user@local ~/projects $ gcommit src/auth/                      │
│                                                                 │
│  📊 変更を解析中...                                             │
│                                                                 │
│  ┌─ 変更サマリー ────────────────────────────────────────────┐ │
│  │  modified:  src/auth/login.go (+45, -12)                  │ │
│  │  modified:  src/auth/session.go (+23, -5)                 │ │
│  │  new file:  src/auth/oauth.go (+120)                      │ │
│  └───────────────────────────────────────────────────────────┘ │
│                                                                 │
│  💡 コミットメッセージ候補（5件）:                              │
│  ┌───────────────────────────────────────────────────────────┐ │
│  │  ▶ 1. feat: OAuth認証機能を追加 ✨                        │ │
│  │    2. feat: Google/GitHub OAuth対応を実装 ✨               │ │
│  │    3. update: 認証システムにOAuthプロバイダーを追加 📝     │ │
│  │    4. feat: ソーシャルログイン機能を実装 ✨                │ │
│  │    5. update: 認証フローにOAuth2.0を統合 📝                │ │
│  └───────────────────────────────────────────────────────────┘ │
│                                                                 │
│  [1-5] 選択  [e] カスタム編集  [r] 再生成  [Esc] キャンセル     │
└─────────────────────────────────────────────────────────────────┘
```

### 3-5-3. 設定

```yaml
# ~/.gonesh/git.yaml

auto_commit:
  enabled: true
  model: "gemini-1.5-flash"    # 高速なモデルを使用
  language: "ja"               # メッセージの言語
  emoji: true                  # 絵文字を付ける
  max_diff_lines: 500          # 解析する最大行数

  # カスタム絵文字マッピング
  emoji_map:
    feat: "✨"
    fix: "🐛"
    update: "📝"
    refactor: "♻️"
    test: "✅"
```

---

## 3-6. コマンド履歴

- シェル終了後も履歴を永続保存。
- ローカル・リモート共通で `~/.gonesh/history` に保存。
- 検索・フィルタ機能（`Ctrl+R`）。
