// Package context provides the UI context and theme for GoNeSh.
package context

import "github.com/charmbracelet/lipgloss"

// Theme defines the color scheme for the UI
type Theme struct {
	// Background colors
	Bg      lipgloss.Color
	BgDark  lipgloss.Color
	BgLight lipgloss.Color

	// Border colors
	Border    lipgloss.Color
	BorderAlt lipgloss.Color

	// Accent colors
	Accent    lipgloss.Color
	Primary   lipgloss.Color
	Secondary lipgloss.Color

	// Text colors
	Text      lipgloss.Color
	TextMuted lipgloss.Color
	TextAlt   lipgloss.Color

	// Semantic colors
	Success lipgloss.Color
	Warning lipgloss.Color
	Error   lipgloss.Color
	Info    lipgloss.Color

	// Resource colors
	CPU lipgloss.Color
	MEM lipgloss.Color
	GPU lipgloss.Color
}

// TokyoNight returns the Tokyo Night theme
func TokyoNight() *Theme {
	return &Theme{
		Bg:      lipgloss.Color("#1a1b26"),
		BgDark:  lipgloss.Color("#0d0e14"),
		BgLight: lipgloss.Color("#24283b"),

		Border:    lipgloss.Color("#3e4451"),
		BorderAlt: lipgloss.Color("#565f89"),

		Accent:    lipgloss.Color("#58e2e8"),
		Primary:   lipgloss.Color("#7aa2f7"),
		Secondary: lipgloss.Color("#bb9af7"),

		Text:      lipgloss.Color("#c0caf5"),
		TextMuted: lipgloss.Color("#5c6370"),
		TextAlt:   lipgloss.Color("#abb2bf"),

		Success: lipgloss.Color("#98c379"),
		Warning: lipgloss.Color("#e5c07b"),
		Error:   lipgloss.Color("#e06c75"),
		Info:    lipgloss.Color("#61afef"),

		CPU: lipgloss.Color("#61afef"),
		MEM: lipgloss.Color("#c678dd"),
		GPU: lipgloss.Color("#98c379"),
	}
}

// UI holds the UI context for rendering components
type UI struct {
	Theme  *Theme
	Width  int
	Height int
}

// New creates a new UI context with the default theme
func New() *UI {
	return &UI{
		Theme: TokyoNight(),
	}
}

// SetSize updates the context dimensions
func (u *UI) SetSize(width, height int) {
	u.Width = width
	u.Height = height
}

// SetTheme updates the theme
func (u *UI) SetTheme(theme *Theme) {
	u.Theme = theme
}
