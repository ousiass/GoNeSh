// Package templates provides layout templates for UI components.
package templates

import (
	"github.com/charmbracelet/lipgloss"
	"github.com/ousiass/GoNeSh/internal/ui/atoms"
	"github.com/ousiass/GoNeSh/internal/ui/context"
)

// Modal renders content in a centered bordered box
func Modal(ctx *context.UI, content string, width, height int) string {
	boxStyle := lipgloss.NewStyle().
		Border(lipgloss.RoundedBorder()).
		BorderForeground(ctx.Theme.Accent).
		Padding(1, 2).
		Background(ctx.Theme.Bg)

	box := boxStyle.Render(content)

	return lipgloss.Place(
		width,
		height,
		lipgloss.Center,
		lipgloss.Center,
		box,
		lipgloss.WithWhitespaceBackground(ctx.Theme.BgDark),
	)
}

// ModalWithPadding renders content in a centered bordered box with custom padding
func ModalWithPadding(ctx *context.UI, content string, width, height, padV, padH int) string {
	boxStyle := lipgloss.NewStyle().
		Border(lipgloss.RoundedBorder()).
		BorderForeground(ctx.Theme.Accent).
		Padding(padV, padH).
		Background(ctx.Theme.Bg)

	box := boxStyle.Render(content)

	return lipgloss.Place(
		width,
		height,
		lipgloss.Center,
		lipgloss.Center,
		box,
		lipgloss.WithWhitespaceBackground(ctx.Theme.BgDark),
	)
}

// CenteredBox renders content in a centered bordered box with standard background
func CenteredBox(ctx *context.UI, content string, width, height int) string {
	boxStyle := lipgloss.NewStyle().
		Border(lipgloss.RoundedBorder()).
		BorderForeground(ctx.Theme.Accent).
		Padding(1, 3).
		Background(ctx.Theme.Bg)

	box := boxStyle.Render(content)

	return lipgloss.Place(
		width,
		height,
		lipgloss.Center,
		lipgloss.Center,
		box,
		lipgloss.WithWhitespaceBackground(ctx.Theme.Bg),
	)
}

// Bar renders a horizontal bar with border
func Bar(ctx *context.UI, content string, width int, borderTop, borderBottom bool) string {
	style := lipgloss.NewStyle().
		Width(width).
		Background(ctx.Theme.Bg).
		Padding(0, 1)

	if borderTop {
		style = style.
			BorderTop(true).
			BorderStyle(atoms.HorizontalRule).
			BorderForeground(ctx.Theme.Border)
	}

	if borderBottom {
		style = style.
			BorderBottom(true).
			BorderStyle(atoms.HorizontalRule).
			BorderForeground(ctx.Theme.Border)
	}

	return style.Render(content)
}

// TopBar renders a bar with bottom border (for tab bar)
func TopBar(ctx *context.UI, content string, width int) string {
	return Bar(ctx, content, width, false, true)
}

// BottomBar renders a bar with top border (for status bar)
func BottomBar(ctx *context.UI, content string, width int) string {
	return Bar(ctx, content, width, true, false)
}
