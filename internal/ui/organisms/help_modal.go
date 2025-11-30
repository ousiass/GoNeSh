// Package organisms provides complex UI components for GoNeSh.
package organisms

import (
	"github.com/charmbracelet/lipgloss"
	"github.com/ousiass/GoNeSh/internal/ui/atoms"
	"github.com/ousiass/GoNeSh/internal/ui/context"
	"github.com/ousiass/GoNeSh/internal/ui/molecules"
	"github.com/ousiass/GoNeSh/internal/ui/templates"
)

// HelpModal represents the keyboard shortcuts help modal
type HelpModal struct {
	ctx    *context.UI
	width  int
	height int
}

// NewHelpModal creates a new help modal
func NewHelpModal(ctx *context.UI) *HelpModal {
	return &HelpModal{ctx: ctx}
}

// SetSize sets the modal size
func (h *HelpModal) SetSize(width, height int) {
	h.width = width
	h.height = height
}

// View renders the help modal
func (h *HelpModal) View() string {
	if h.width == 0 || h.height == 0 {
		return ""
	}

	// Build sections using molecules
	tabsSection := molecules.Section(h.ctx, atoms.IconTabs, "TABS", []string{
		molecules.HelpItem(h.ctx, "t", "New"),
		molecules.HelpItem(h.ctx, "w", "Close"),
		molecules.HelpItem(h.ctx, "]", "Next"),
		molecules.HelpItem(h.ctx, "[", "Prev"),
	})

	aiSection := molecules.Section(h.ctx, atoms.IconAI, "AI", []string{
		molecules.HelpItem(h.ctx, "a", "Panel"),
		molecules.HelpItem(h.ctx, "c", "Claude"),
		molecules.HelpItem(h.ctx, "p", "Presets"),
		molecules.HelpItem(h.ctx, "x", "External"),
	})

	filesSection := molecules.Section(h.ctx, atoms.IconFiles, "FILES", []string{
		molecules.HelpItem(h.ctx, "f", "Browser"),
		molecules.HelpItem(h.ctx, "s", "Transfer"),
		molecules.HelpItem(h.ctx, "r", "API"),
		molecules.HelpItem(h.ctx, "g", "Git"),
	})

	// 3-column layout
	colStyle := lipgloss.NewStyle().
		Padding(0, 1).
		Background(h.ctx.Theme.Bg)

	columns := lipgloss.JoinHorizontal(
		lipgloss.Top,
		colStyle.Render(tabsSection),
		colStyle.Render(aiSection),
		colStyle.Render(filesSection),
	)

	contentWidth := lipgloss.Width(columns)

	// Title and footer
	title := atoms.CenteredText(h.ctx, atoms.IconKeyboard+"  Keyboard Shortcuts", contentWidth, h.ctx.Theme.Accent)
	emptyRow := atoms.Fill(h.ctx, contentWidth)
	footer := atoms.CenteredText(h.ctx, "Press key to execute • ESC/? to close", contentWidth, h.ctx.Theme.TextMuted)

	content := lipgloss.JoinVertical(
		lipgloss.Center,
		title,
		emptyRow,
		columns,
		emptyRow,
		footer,
	)

	return templates.Modal(h.ctx, content, h.width, h.height)
}
