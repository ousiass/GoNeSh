// Package errors provides error handling with error codes for GoNeSh.
// Each error has a unique code for easy identification and documentation.
package errors

import (
	"fmt"
)

// ErrorCode represents a unique error identifier
type ErrorCode string

// Error categories
const (
	// E1xxx: Configuration errors
	E1001 ErrorCode = "E1001" // Config file not found
	E1002 ErrorCode = "E1002" // Config parse error
	E1003 ErrorCode = "E1003" // Invalid config value
	E1004 ErrorCode = "E1004" // Config directory creation failed

	// E2xxx: UI errors
	E2001 ErrorCode = "E2001" // Terminal initialization failed
	E2002 ErrorCode = "E2002" // Screen size too small
	E2003 ErrorCode = "E2003" // Render error
	E2004 ErrorCode = "E2004" // Input handling error

	// E3xxx: SSH/Portal errors
	E3001 ErrorCode = "E3001" // SSH connection failed
	E3002 ErrorCode = "E3002" // Authentication failed
	E3003 ErrorCode = "E3003" // Host not found
	E3004 ErrorCode = "E3004" // Key file not found
	E3005 ErrorCode = "E3005" // SCP transfer failed
	E3006 ErrorCode = "E3006" // Session timeout

	// E4xxx: AI errors
	E4001 ErrorCode = "E4001" // API key not set
	E4002 ErrorCode = "E4002" // API request failed
	E4003 ErrorCode = "E4003" // API rate limit exceeded
	E4004 ErrorCode = "E4004" // Invalid model specified
	E4005 ErrorCode = "E4005" // Response parse error

	// E5xxx: API Client errors
	E5001 ErrorCode = "E5001" // OpenAPI spec parse error
	E5002 ErrorCode = "E5002" // HTTP request failed
	E5003 ErrorCode = "E5003" // Invalid URL
	E5004 ErrorCode = "E5004" // Environment variable not found

	// E6xxx: Git errors
	E6001 ErrorCode = "E6001" // Not a git repository
	E6002 ErrorCode = "E6002" // No changes to commit
	E6003 ErrorCode = "E6003" // Git command failed

	// E9xxx: System errors
	E9001 ErrorCode = "E9001" // File system error
	E9002 ErrorCode = "E9002" // Process execution error
	E9003 ErrorCode = "E9003" // Resource monitoring error
	E9999 ErrorCode = "E9999" // Unknown error
)

// errorMessages maps error codes to their descriptions
var errorMessages = map[ErrorCode]string{
	E1001: "設定ファイルが見つかりません",
	E1002: "設定ファイルの解析に失敗しました",
	E1003: "無効な設定値です",
	E1004: "設定ディレクトリの作成に失敗しました",

	E2001: "ターミナルの初期化に失敗しました",
	E2002: "画面サイズが小さすぎます",
	E2003: "描画エラーが発生しました",
	E2004: "入力処理エラーが発生しました",

	E3001: "SSH接続に失敗しました",
	E3002: "認証に失敗しました",
	E3003: "ホストが見つかりません",
	E3004: "秘密鍵ファイルが見つかりません",
	E3005: "SCPファイル転送に失敗しました",
	E3006: "セッションがタイムアウトしました",

	E4001: "APIキーが設定されていません",
	E4002: "APIリクエストに失敗しました",
	E4003: "APIレート制限を超過しました",
	E4004: "無効なモデルが指定されました",
	E4005: "レスポンスの解析に失敗しました",

	E5001: "OpenAPI仕様の解析に失敗しました",
	E5002: "HTTPリクエストに失敗しました",
	E5003: "無効なURLです",
	E5004: "環境変数が見つかりません",

	E6001: "Gitリポジトリではありません",
	E6002: "コミットする変更がありません",
	E6003: "Gitコマンドの実行に失敗しました",

	E9001: "ファイルシステムエラーが発生しました",
	E9002: "プロセス実行エラーが発生しました",
	E9003: "リソース監視エラーが発生しました",
	E9999: "不明なエラーが発生しました",
}

// GoNeShError represents an error with a code
type GoNeShError struct {
	Code    ErrorCode
	Message string
	Cause   error
}

// Error implements the error interface
func (e *GoNeShError) Error() string {
	if e.Cause != nil {
		return fmt.Sprintf("[%s] %s: %v", e.Code, e.Message, e.Cause)
	}
	return fmt.Sprintf("[%s] %s", e.Code, e.Message)
}

// Unwrap returns the underlying error
func (e *GoNeShError) Unwrap() error {
	return e.Cause
}

// New creates a new GoNeShError with the given code
func New(code ErrorCode) *GoNeShError {
	msg, ok := errorMessages[code]
	if !ok {
		msg = errorMessages[E9999]
	}
	return &GoNeShError{
		Code:    code,
		Message: msg,
	}
}

// Wrap wraps an existing error with an error code
func Wrap(code ErrorCode, cause error) *GoNeShError {
	err := New(code)
	err.Cause = cause
	return err
}

// WithMessage creates an error with a custom message
func WithMessage(code ErrorCode, message string) *GoNeShError {
	return &GoNeShError{
		Code:    code,
		Message: message,
	}
}

// DocURL returns the documentation URL for an error code
func (e *GoNeShError) DocURL() string {
	return fmt.Sprintf("https://gonesh.ousiass.com/docs/errors/%s", e.Code)
}

// GetMessage returns the default message for an error code
func GetMessage(code ErrorCode) string {
	if msg, ok := errorMessages[code]; ok {
		return msg
	}
	return errorMessages[E9999]
}
