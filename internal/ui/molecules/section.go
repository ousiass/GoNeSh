package molecules

import (
	"github.com/charmbracelet/lipgloss"
	"github.com/ousiass/GoNeSh/internal/ui/context"
)

const sectionWidth = 16

// SectionHeader renders a section header with icon and title
func SectionHeader(ctx *context.UI, icon, title string) string {
	return lipgloss.NewStyle().
		Width(sectionWidth).
		Foreground(ctx.Theme.Secondary).
		Background(ctx.Theme.Bg).
		Bold(true).
		Render(icon + "  " + title)
}

// Section renders a complete section with header and items
func Section(ctx *context.UI, icon, title string, items []string) string {
	header := SectionHeader(ctx, icon, title)

	content := []string{header}
	content = append(content, items...)

	return lipgloss.JoinVertical(lipgloss.Left, content...)
}
