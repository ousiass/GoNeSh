package atoms

import (
	"github.com/charmbracelet/lipgloss"
	"github.com/ousiass/GoNeSh/internal/ui/context"
)

// Text renders plain text with default styling
func Text(ctx *context.UI, s string) string {
	return lipgloss.NewStyle().
		Foreground(ctx.Theme.Text).
		Background(ctx.Theme.Bg).
		Render(s)
}

// TextMuted renders muted/dimmed text
func TextMuted(ctx *context.UI, s string) string {
	return lipgloss.NewStyle().
		Foreground(ctx.Theme.TextMuted).
		Background(ctx.Theme.Bg).
		Render(s)
}

// TextAlt renders alternative colored text
func TextAlt(ctx *context.UI, s string) string {
	return lipgloss.NewStyle().
		Foreground(ctx.Theme.TextAlt).
		Background(ctx.Theme.Bg).
		Render(s)
}

// Title renders a bold title with accent color
func Title(ctx *context.UI, s string) string {
	return lipgloss.NewStyle().
		Foreground(ctx.Theme.Accent).
		Background(ctx.Theme.Bg).
		Bold(true).
		Render(s)
}

// Subtitle renders italic text with info color
func Subtitle(ctx *context.UI, s string) string {
	return lipgloss.NewStyle().
		Foreground(ctx.Theme.Info).
		Background(ctx.Theme.Bg).
		Italic(true).
		Render(s)
}

// Label renders a bold label with specified color
func Label(ctx *context.UI, s string, color lipgloss.Color) string {
	return lipgloss.NewStyle().
		Foreground(color).
		Background(ctx.Theme.Bg).
		Bold(true).
		Render(s)
}

// Version renders version text with secondary color
func Version(ctx *context.UI, s string) string {
	return lipgloss.NewStyle().
		Foreground(ctx.Theme.Secondary).
		Background(ctx.Theme.Bg).
		Render(s)
}

// ErrorText renders error text
func ErrorText(ctx *context.UI, s string) string {
	return lipgloss.NewStyle().
		Foreground(ctx.Theme.Error).
		Background(ctx.Theme.Bg).
		Render(s)
}

// SuccessText renders success text
func SuccessText(ctx *context.UI, s string) string {
	return lipgloss.NewStyle().
		Foreground(ctx.Theme.Success).
		Background(ctx.Theme.Bg).
		Render(s)
}

// WarningText renders warning text
func WarningText(ctx *context.UI, s string) string {
	return lipgloss.NewStyle().
		Foreground(ctx.Theme.Warning).
		Background(ctx.Theme.Bg).
		Render(s)
}

// CenteredText renders centered text with specified width and color
func CenteredText(ctx *context.UI, s string, width int, color lipgloss.Color) string {
	return lipgloss.NewStyle().
		Width(width).
		Align(lipgloss.Center).
		Foreground(color).
		Background(ctx.Theme.Bg).
		Render(s)
}

// Separator renders a text separator
func Separator(ctx *context.UI) string {
	return lipgloss.NewStyle().
		Foreground(ctx.Theme.Border).
		Background(ctx.Theme.Bg).
		Render("  ")
}
