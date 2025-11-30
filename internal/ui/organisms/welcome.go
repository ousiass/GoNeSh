// Package organisms provides complex UI components for GoNeSh.
package organisms

import (
	"time"

	"github.com/charmbracelet/bubbles/spinner"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	figure "github.com/common-nighthawk/go-figure"
	"github.com/ousiass/GoNeSh/internal/ui/atoms"
	"github.com/ousiass/GoNeSh/internal/ui/context"
	"github.com/ousiass/GoNeSh/internal/ui/templates"
)

// Welcome represents the welcome screen component
type Welcome struct {
	ctx     *context.UI
	spinner spinner.Model
	width   int
	height  int
}

// NewWelcome creates a new welcome screen
func NewWelcome(ctx *context.UI) *Welcome {
	s := spinner.New()
	s.Spinner = spinner.Spinner{
		Frames: []string{"◜", "◠", "◝", "◞", "◡", "◟"},
		FPS:    time.Second / 10,
	}
	s.Style = lipgloss.NewStyle().
		Foreground(ctx.Theme.Accent).
		Background(ctx.Theme.Bg)

	return &Welcome{
		ctx:     ctx,
		spinner: s,
	}
}

// Init returns the initial command
func (w *Welcome) Init() tea.Cmd {
	return w.spinner.Tick
}

// SetSize sets the welcome screen size
func (w *Welcome) SetSize(width, height int) {
	w.width = width
	w.height = height
}

// Update handles spinner updates
func (w *Welcome) Update(msg tea.Msg) (*Welcome, tea.Cmd) {
	if tickMsg, ok := msg.(spinner.TickMsg); ok {
		var cmd tea.Cmd
		w.spinner, cmd = w.spinner.Update(tickMsg)
		return w, cmd
	}
	return w, nil
}

// View renders the welcome screen
func (w *Welcome) View() string {
	if w.width == 0 || w.height == 0 {
		return ""
	}

	// Generate ASCII art
	fig := figure.NewFigure("GoNeSh", "slant", true)
	asciiLogo := fig.String()

	// Get logo width for centering
	logoRendered := atoms.Title(w.ctx, asciiLogo)
	contentWidth := lipgloss.Width(logoRendered)

	// Build content rows
	logo := atoms.CenteredText(w.ctx, asciiLogo, contentWidth, w.ctx.Theme.Accent)
	version := atoms.CenteredText(w.ctx, "v0.1.0", contentWidth, w.ctx.Theme.Secondary)
	subtitle := lipgloss.NewStyle().
		Width(contentWidth).
		Align(lipgloss.Center).
		Foreground(w.ctx.Theme.Info).
		Background(w.ctx.Theme.Bg).
		Italic(true).
		Render(atoms.IconStar + " Prompt is the Desktop " + atoms.IconStar)
	emptyRow := atoms.Fill(w.ctx, contentWidth)
	hint := atoms.CenteredText(w.ctx, w.spinner.View()+" ? for shortcuts", contentWidth, w.ctx.Theme.TextMuted)

	innerContent := lipgloss.JoinVertical(
		lipgloss.Center,
		logo,
		version,
		subtitle,
		emptyRow,
		hint,
	)

	return templates.CenteredBox(w.ctx, innerContent, w.width, w.height)
}
