// Package organisms provides complex UI components for GoNeSh.
package organisms

import (
	"github.com/charmbracelet/lipgloss"
	"github.com/ousiass/GoNeSh/internal/ui/atoms"
	"github.com/ousiass/GoNeSh/internal/ui/context"
	"github.com/ousiass/GoNeSh/internal/ui/molecules"
	"github.com/ousiass/GoNeSh/internal/ui/templates"
)

// Tab represents a terminal tab
type Tab struct {
	Name string
	Type molecules.TabType
}

// TabBar represents the tab bar component
type TabBar struct {
	ctx       *context.UI
	tabs      []Tab
	activeTab int
	width     int
}

// NewTabBar creates a new tab bar
func NewTabBar(ctx *context.UI) *TabBar {
	return &TabBar{
		ctx: ctx,
		tabs: []Tab{
			{Name: "local", Type: molecules.TabTypeLocal},
		},
		activeTab: 0,
	}
}

// SetWidth sets the tab bar width
func (t *TabBar) SetWidth(width int) {
	t.width = width
}

// AddTab adds a new tab
func (t *TabBar) AddTab(name string, tabType string) {
	tt := molecules.TabTypeLocal
	if tabType == "ssh" {
		tt = molecules.TabTypeSSH
	}
	t.tabs = append(t.tabs, Tab{Name: name, Type: tt})
	t.activeTab = len(t.tabs) - 1
}

// CloseTab closes the current tab, returns true if app should quit
func (t *TabBar) CloseTab() bool {
	if len(t.tabs) <= 1 {
		return true // quit
	}
	t.tabs = append(t.tabs[:t.activeTab], t.tabs[t.activeTab+1:]...)
	if t.activeTab >= len(t.tabs) {
		t.activeTab = len(t.tabs) - 1
	}
	return false
}

// NextTab switches to the next tab
func (t *TabBar) NextTab() {
	t.activeTab = (t.activeTab + 1) % len(t.tabs)
}

// PrevTab switches to the previous tab
func (t *TabBar) PrevTab() {
	t.activeTab = (t.activeTab - 1 + len(t.tabs)) % len(t.tabs)
}

// ActiveTab returns the current active tab
func (t *TabBar) ActiveTab() Tab {
	return t.tabs[t.activeTab]
}

// Tabs returns all tabs
func (t *TabBar) Tabs() []Tab {
	return t.tabs
}

// View renders the tab bar
func (t *TabBar) View() string {
	if t.width == 0 {
		return ""
	}

	var tabs []string
	for i, tab := range t.tabs {
		tabs = append(tabs, molecules.Tab(t.ctx, tab.Name, tab.Type, i == t.activeTab))
	}

	// Join tabs horizontally
	tabsJoined := lipgloss.JoinHorizontal(lipgloss.Top, tabs...)

	// Fill remaining space
	tabsWidth := lipgloss.Width(tabsJoined)
	fill := atoms.Fill(t.ctx, t.width-tabsWidth)

	row := lipgloss.JoinHorizontal(lipgloss.Top, tabsJoined, fill)

	return templates.TopBar(t.ctx, row, t.width)
}
