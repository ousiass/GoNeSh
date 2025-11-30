package atoms

import (
	"github.com/charmbracelet/lipgloss"
	"github.com/ousiass/GoNeSh/internal/ui/context"
)

// KeyBadge renders a keyboard key badge
func KeyBadge(ctx *context.UI, key string) string {
	return lipgloss.NewStyle().
		Foreground(ctx.Theme.Accent).
		Background(ctx.Theme.Bg).
		Bold(true).
		Width(3).
		Align(lipgloss.Center).
		Render(key)
}

// EnvBadge renders an environment badge with appropriate color and icon
func EnvBadge(ctx *context.UI, env string) string {
	var color lipgloss.Color
	var icon string

	switch env {
	case "prod":
		color = ctx.Theme.Error
		icon = IconProd
	case "staging":
		color = ctx.Theme.Warning
		icon = IconStaging
	default: // dev, local
		color = ctx.Theme.Success
		icon = IconDev
	}

	return lipgloss.NewStyle().
		Foreground(color).
		Background(ctx.Theme.Bg).
		Bold(true).
		Render(icon + " " + env)
}

// PresetBadge renders a preset name badge
func PresetBadge(ctx *context.UI, preset string) string {
	if preset == "" {
		return ""
	}
	return lipgloss.NewStyle().
		Foreground(ctx.Theme.Secondary).
		Background(ctx.Theme.Bg).
		Render(preset + " ● ")
}

// ActiveBadge renders an active/selected state badge
func ActiveBadge(ctx *context.UI, text string) string {
	return lipgloss.NewStyle().
		Padding(0, 2).
		Bold(true).
		Foreground(ctx.Theme.Bg).
		Background(ctx.Theme.Accent).
		Render(text)
}

// InactiveBadge renders an inactive state badge
func InactiveBadge(ctx *context.UI, text string) string {
	return lipgloss.NewStyle().
		Padding(0, 2).
		Foreground(ctx.Theme.TextMuted).
		Background(ctx.Theme.Bg).
		Render(text)
}
