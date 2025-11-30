// Package organisms provides complex UI components for GoNeSh.
package organisms

import (
	"strings"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/ousiass/GoNeSh/internal/history"
	"github.com/ousiass/GoNeSh/internal/ui/context"
)

// HistorySearchResult is sent when a history entry is selected
type HistorySearchResult struct {
	Entry    string
	Selected bool
}

// HistorySearch represents the Ctrl+R history search UI
type HistorySearch struct {
	ctx      *context.UI
	history  *history.History
	query    string
	results  []string
	selected int
	width    int
	height   int
	visible  bool
	searchPos int // Current search position in history
}

// NewHistorySearch creates a new history search component
func NewHistorySearch(ctx *context.UI, h *history.History) *HistorySearch {
	return &HistorySearch{
		ctx:       ctx,
		history:   h,
		results:   []string{},
		selected:  0,
		searchPos: -1,
	}
}

// Show shows the history search UI
func (hs *HistorySearch) Show() {
	hs.visible = true
	hs.query = ""
	hs.selected = 0
	hs.searchPos = hs.history.Len() - 1
	hs.updateResults()
}

// Hide hides the history search UI
func (hs *HistorySearch) Hide() {
	hs.visible = false
	hs.query = ""
	hs.results = []string{}
}

// IsVisible returns whether the search UI is visible
func (hs *HistorySearch) IsVisible() bool {
	return hs.visible
}

// SetSize sets the component size
func (hs *HistorySearch) SetSize(width, height int) {
	hs.width = width
	hs.height = height
}

// Update handles messages for the history search
func (hs *HistorySearch) Update(msg tea.Msg) (*HistorySearch, tea.Cmd) {
	if !hs.visible {
		return hs, nil
	}

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.Type {
		case tea.KeyEscape, tea.KeyCtrlC, tea.KeyCtrlG:
			hs.Hide()
			return hs, func() tea.Msg {
				return HistorySearchResult{Selected: false}
			}

		case tea.KeyEnter:
			if len(hs.results) > 0 && hs.selected < len(hs.results) {
				entry := hs.results[hs.selected]
				hs.Hide()
				return hs, func() tea.Msg {
					return HistorySearchResult{Entry: entry, Selected: true}
				}
			}
			hs.Hide()
			return hs, func() tea.Msg {
				return HistorySearchResult{Selected: false}
			}

		case tea.KeyCtrlR:
			// Search for next match (going backwards)
			if len(hs.results) > 0 {
				hs.selected = (hs.selected + 1) % len(hs.results)
			}
			return hs, nil

		case tea.KeyUp:
			if hs.selected > 0 {
				hs.selected--
			}
			return hs, nil

		case tea.KeyDown:
			if hs.selected < len(hs.results)-1 {
				hs.selected++
			}
			return hs, nil

		case tea.KeyBackspace:
			if len(hs.query) > 0 {
				hs.query = hs.query[:len(hs.query)-1]
				hs.updateResults()
			}
			return hs, nil

		case tea.KeyRunes:
			hs.query += string(msg.Runes)
			hs.updateResults()
			return hs, nil
		}
	}

	return hs, nil
}

// updateResults updates the search results based on the current query
func (hs *HistorySearch) updateResults() {
	hs.results = hs.history.Search(hs.query)
	hs.selected = 0

	// Limit results to fit in view
	maxResults := hs.height - 4
	if maxResults < 1 {
		maxResults = 5
	}
	if len(hs.results) > maxResults {
		hs.results = hs.results[:maxResults]
	}
}

// View renders the history search UI
func (hs *HistorySearch) View() string {
	if !hs.visible || hs.width == 0 {
		return ""
	}

	// Build the search box
	var sb strings.Builder

	// Header
	headerStyle := lipgloss.NewStyle().
		Foreground(hs.ctx.Theme.Primary).
		Bold(true)
	sb.WriteString(headerStyle.Render("(reverse-i-search)"))
	sb.WriteString(": ")

	// Query
	queryStyle := lipgloss.NewStyle().
		Foreground(hs.ctx.Theme.Text)
	sb.WriteString(queryStyle.Render(hs.query))
	sb.WriteString("_")
	sb.WriteString("\n")

	// Results
	for i, result := range hs.results {
		var line string
		if i == hs.selected {
			// Selected item
			selectedStyle := lipgloss.NewStyle().
				Background(hs.ctx.Theme.Primary).
				Foreground(hs.ctx.Theme.Bg).
				Width(hs.width - 4)
			line = selectedStyle.Render(truncate(result, hs.width-6))
		} else {
			// Normal item
			normalStyle := lipgloss.NewStyle().
				Foreground(hs.ctx.Theme.TextAlt).
				Width(hs.width - 4)
			line = normalStyle.Render(truncate(result, hs.width-6))
		}
		sb.WriteString("  " + line + "\n")
	}

	// Footer hint
	hintStyle := lipgloss.NewStyle().
		Foreground(hs.ctx.Theme.TextAlt).
		Italic(true)
	sb.WriteString(hintStyle.Render("Ctrl+R: next | Enter: select | Esc: cancel"))

	// Box style
	boxStyle := lipgloss.NewStyle().
		Border(lipgloss.RoundedBorder()).
		BorderForeground(hs.ctx.Theme.Border).
		Padding(0, 1).
		Width(hs.width - 2)

	return boxStyle.Render(sb.String())
}

// truncate truncates a string to the given length
func truncate(s string, maxLen int) string {
	if maxLen <= 0 {
		return ""
	}
	if len(s) <= maxLen {
		return s
	}
	if maxLen <= 3 {
		return s[:maxLen]
	}
	return s[:maxLen-3] + "..."
}
