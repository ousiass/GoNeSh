package molecules

import (
	"github.com/charmbracelet/lipgloss"
	"github.com/ousiass/GoNeSh/internal/ui/atoms"
	"github.com/ousiass/GoNeSh/internal/ui/context"
)

const helpItemWidth = 16

// HelpItem renders a keyboard shortcut help item (key + description)
func HelpItem(ctx *context.UI, key, description string) string {
	keyBadge := atoms.KeyBadge(ctx, key)
	desc := atoms.TextAlt(ctx, " "+description)

	item := lipgloss.JoinHorizontal(lipgloss.Center, keyBadge, desc)

	return lipgloss.NewStyle().
		Width(helpItemWidth).
		Background(ctx.Theme.Bg).
		Render(item)
}
