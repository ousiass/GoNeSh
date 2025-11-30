package atoms

import (
	"github.com/charmbracelet/lipgloss"
	"github.com/ousiass/GoNeSh/internal/ui/context"
)

// HorizontalRule style for borders
var HorizontalRule = lipgloss.Border{
	Top:    "─",
	Bottom: "─",
}

// Fill creates a filler block with background
func Fill(ctx *context.UI, width int) string {
	return lipgloss.NewStyle().
		Width(width).
		Background(ctx.Theme.Bg).
		Render("")
}

// BaseStyle returns a base style with background
func BaseStyle(ctx *context.UI) lipgloss.Style {
	return lipgloss.NewStyle().Background(ctx.Theme.Bg)
}

// BorderedBox returns a style with rounded border
func BorderedBox(ctx *context.UI) lipgloss.Style {
	return lipgloss.NewStyle().
		Border(lipgloss.RoundedBorder()).
		BorderForeground(ctx.Theme.Accent).
		Background(ctx.Theme.Bg)
}
