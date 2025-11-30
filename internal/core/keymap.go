// Package core provides the main application logic for GoNeSh.
package core

import "github.com/charmbracelet/bubbles/key"

// KeyMap defines the keybindings for the application
type KeyMap struct {
	Quit          key.Binding
	NewTab        key.Binding
	CloseTab      key.Binding
	NextTab       key.Binding
	PrevTab       key.Binding
	ToggleAI      key.Binding
	SelectPreset  key.Binding
	FileBrowser   key.Binding
	QuickTransfer key.Binding
	APIClient     key.Binding
	GitCommit     key.Binding
	ClaudeCode    key.Binding
	ExternalAI    key.Binding
	ShowHelp      key.Binding
}

// DefaultKeyMap returns the default keybindings
func DefaultKeyMap() KeyMap {
	return KeyMap{
		Quit: key.NewBinding(
			key.WithKeys("ctrl+c", "ctrl+q"),
			key.WithHelp("ctrl+c", "終了"),
		),
		NewTab: key.NewBinding(
			key.WithKeys("alt+t"),
			key.WithHelp("alt+t", "新規タブ"),
		),
		CloseTab: key.NewBinding(
			key.WithKeys("alt+w"),
			key.WithHelp("alt+w", "閉じる"),
		),
		NextTab: key.NewBinding(
			key.WithKeys("alt+]"),
			key.WithHelp("alt+]", "次タブ"),
		),
		PrevTab: key.NewBinding(
			key.WithKeys("alt+["),
			key.WithHelp("alt+[", "前タブ"),
		),
		ToggleAI: key.NewBinding(
			key.WithKeys("alt+a"),
			key.WithHelp("alt+a", "AI"),
		),
		SelectPreset: key.NewBinding(
			key.WithKeys("alt+p"),
			key.WithHelp("alt+p", "プリセット"),
		),
		FileBrowser: key.NewBinding(
			key.WithKeys("alt+f"),
			key.WithHelp("alt+f", "ファイル"),
		),
		QuickTransfer: key.NewBinding(
			key.WithKeys("alt+s"),
			key.WithHelp("alt+s", "転送"),
		),
		APIClient: key.NewBinding(
			key.WithKeys("alt+r"),
			key.WithHelp("alt+r", "API"),
		),
		GitCommit: key.NewBinding(
			key.WithKeys("alt+g"),
			key.WithHelp("alt+g", "Git"),
		),
		ClaudeCode: key.NewBinding(
			key.WithKeys("alt+c"),
			key.WithHelp("alt+c", "Claude"),
		),
		ExternalAI: key.NewBinding(
			key.WithKeys("alt+x"),
			key.WithHelp("alt+x", "外部AI"),
		),
		ShowHelp: key.NewBinding(
			key.WithKeys("?"),
			key.WithHelp("?", "ヘルプ"),
		),
	}
}

// ShortHelp returns keybindings to be shown in the mini help view
func (k KeyMap) ShortHelp() []key.Binding {
	return []key.Binding{k.NewTab, k.CloseTab, k.ToggleAI, k.ShowHelp}
}

// FullHelp returns keybindings for the expanded help view (grouped by category)
func (k KeyMap) FullHelp() [][]key.Binding {
	return [][]key.Binding{
		{k.NewTab, k.CloseTab, k.NextTab, k.PrevTab},             // タブ操作
		{k.ToggleAI, k.SelectPreset, k.ClaudeCode, k.ExternalAI}, // AI
		{k.FileBrowser, k.QuickTransfer, k.APIClient, k.GitCommit}, // ファイル
	}
}
