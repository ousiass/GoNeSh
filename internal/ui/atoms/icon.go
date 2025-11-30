package atoms

import (
	"github.com/charmbracelet/lipgloss"
	"github.com/ousiass/GoNeSh/internal/ui/context"
)

// Icon constants (Nerd Fonts)
const (
	// Tab icons
	IconTerminal = ""
	IconSSH      = "󰣀"

	// Section icons
	IconTabs   = "󰓩"
	IconAI     = ""
	IconFiles  = ""
	IconGit    = ""
	IconAPI    = ""

	// Environment icons
	IconDev     = ""
	IconProd    = ""
	IconStaging = ""

	// Status icons
	IconWarning  = "⚠"
	IconError    = ""
	IconSuccess  = ""
	IconInfo     = ""
	IconKeyboard = "⌨"
	IconStar     = "✦"
)

// Icon renders an icon with specified color
func Icon(ctx *context.UI, icon string, color lipgloss.Color) string {
	return lipgloss.NewStyle().
		Foreground(color).
		Background(ctx.Theme.Bg).
		Render(icon)
}

// IconWithText renders an icon followed by text
func IconWithText(ctx *context.UI, icon, text string, color lipgloss.Color) string {
	return lipgloss.NewStyle().
		Foreground(color).
		Background(ctx.Theme.Bg).
		Render(icon + " " + text)
}

// IconAccent renders an icon with accent color
func IconAccent(ctx *context.UI, icon string) string {
	return Icon(ctx, icon, ctx.Theme.Accent)
}

// IconMuted renders an icon with muted color
func IconMuted(ctx *context.UI, icon string) string {
	return Icon(ctx, icon, ctx.Theme.TextMuted)
}
