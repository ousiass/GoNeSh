// Package core provides the main application logic for GoNeSh.
package core

import (
	"github.com/charmbracelet/bubbles/help"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/ousiass/GoNeSh/internal/history"
	"github.com/ousiass/GoNeSh/internal/ui/context"
	"github.com/ousiass/GoNeSh/internal/ui/organisms"
	"github.com/ousiass/GoNeSh/pkg/config"
)

// App represents the main application state
type App struct {
	config    *config.Config
	ui        *context.UI
	keys      KeyMap
	help      help.Model
	width     int
	height    int
	tabBar    *organisms.TabBar
	statusBar *organisms.StatusBar
	helpModal *organisms.HelpModal
	showHelp  bool

	// Terminal sessions per tab
	terminals         map[int]*organisms.Terminal
	terminalIDCounter int

	// Command history
	history       *history.History
	historySearch *organisms.HistorySearch
}

// NewApp creates a new application instance
func NewApp(cfg *config.Config) *App {
	// Create UI context
	ui := context.New()

	h := help.New()
	h.ShowAll = true
	h.Styles.ShortKey = lipgloss.NewStyle().Foreground(ui.Theme.Primary).Bold(true)
	h.Styles.ShortDesc = lipgloss.NewStyle().Foreground(ui.Theme.Text)
	h.Styles.ShortSeparator = lipgloss.NewStyle().Foreground(ui.Theme.BorderAlt)
	h.Styles.FullKey = lipgloss.NewStyle().Foreground(ui.Theme.Secondary).Bold(true)
	h.Styles.FullDesc = lipgloss.NewStyle().Foreground(ui.Theme.Text)
	h.Styles.FullSeparator = lipgloss.NewStyle().Foreground(ui.Theme.BorderAlt)

	// Initialize history
	hist := history.New(history.DefaultMaxEntries)
	_ = hist.Load() // Load existing history

	app := &App{
		config:            cfg,
		ui:                ui,
		keys:              DefaultKeyMap(),
		help:              h,
		tabBar:            organisms.NewTabBar(ui),
		statusBar:         organisms.NewStatusBar(ui),
		helpModal:         organisms.NewHelpModal(ui),
		terminals:         make(map[int]*organisms.Terminal),
		terminalIDCounter: 0,
		history:           hist,
		historySearch:     organisms.NewHistorySearch(ui, hist),
	}

	// Create initial terminal for the first tab
	app.terminals[0] = organisms.NewTerminal(ui, 0)

	return app
}

// Init initializes the application
func (a *App) Init() tea.Cmd {
	cmds := []tea.Cmd{
		a.statusBar.Init(),
		tea.SetWindowTitle("GoNeSh"),
	}

	// Initialize the first terminal
	if term, ok := a.terminals[0]; ok {
		cmds = append(cmds, term.Init())
	}

	return tea.Batch(cmds...)
}

// Update handles messages and updates the application state
func (a *App) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmds []tea.Cmd

	// Handle history search result
	if result, ok := msg.(organisms.HistorySearchResult); ok {
		if result.Selected && result.Entry != "" {
			// Send selected history entry to terminal
			if term := a.activeTerminal(); term != nil {
				term.SendInput(result.Entry)
			}
		}
		return a, nil
	}

	// If history search is visible, forward messages to it
	if a.historySearch.IsVisible() {
		var cmd tea.Cmd
		a.historySearch, cmd = a.historySearch.Update(msg)
		if cmd != nil {
			cmds = append(cmds, cmd)
		}
		return a, tea.Batch(cmds...)
	}

	switch msg := msg.(type) {
	case tea.KeyMsg:
		// ヘルプ表示中はショートカットキーを直接受け付ける
		if a.showHelp {
			a.showHelp = false
			key := msg.String()
			switch key {
			case "t":
				cmd := a.addNewTab()
				return a, cmd
			case "w":
				if a.closeCurrentTab() {
					return a, tea.Quit
				}
				return a, nil
			case "]":
				a.tabBar.NextTab()
				return a, nil
			case "[":
				a.tabBar.PrevTab()
				return a, nil
			case "a", "p", "c", "x", "f", "s", "r", "g":
				// TODO: 実装
				return a, nil
			case "esc", "?":
				return a, nil
			}
			return a, nil
		}

		// 通常時 - Alt キーのショートカット
		switch msg.String() {
		case "ctrl+c", "ctrl+q":
			// Save history before quitting
			_ = a.history.Save()
			a.closeAllTerminals()
			return a, tea.Quit
		case "ctrl+r":
			// Show history search
			a.historySearch.SetSize(a.width, a.calculateContentHeight())
			a.historySearch.Show()
			return a, nil
		case "?":
			a.showHelp = true
			return a, nil
		case "alt+t":
			cmd := a.addNewTab()
			return a, cmd
		case "alt+w":
			if a.closeCurrentTab() {
				return a, tea.Quit
			}
			return a, nil
		case "alt+]":
			a.tabBar.NextTab()
			return a, nil
		case "alt+[":
			a.tabBar.PrevTab()
			return a, nil
		case "alt+a": // AIパネル
		case "alt+p": // プリセット
		case "alt+c": // Claude
		case "alt+x": // 外部AI
		case "alt+f": // ファイル
		case "alt+s": // 転送
		case "alt+r": // API
		case "alt+g": // Git
		default:
			// Forward key to active terminal
			if term := a.activeTerminal(); term != nil {
				var cmd tea.Cmd
				term, cmd = term.Update(msg)
				a.terminals[a.tabBar.ActiveTabIndex()] = term
				cmds = append(cmds, cmd)
			}
		}

	case tea.WindowSizeMsg:
		a.width = msg.Width
		a.height = msg.Height
		a.ui.SetSize(msg.Width, msg.Height)
		a.tabBar.SetWidth(msg.Width)
		a.statusBar.SetWidth(msg.Width)
		a.help.Width = msg.Width

		// Update terminal sizes
		contentHeight := a.calculateContentHeight()
		for _, term := range a.terminals {
			term.SetSize(msg.Width, contentHeight)
		}
		a.historySearch.SetSize(msg.Width, contentHeight)

	default:
		// Forward other messages to active terminal
		if term := a.activeTerminal(); term != nil {
			var cmd tea.Cmd
			term, cmd = term.Update(msg)
			a.terminals[a.tabBar.ActiveTabIndex()] = term
			cmds = append(cmds, cmd)
		}
	}

	// ステータスバーを更新
	var cmd tea.Cmd
	a.statusBar, cmd = a.statusBar.Update(msg)
	cmds = append(cmds, cmd)

	return a, tea.Batch(cmds...)
}

// View renders the application
func (a *App) View() string {
	if a.width == 0 {
		return "Loading..."
	}

	// タブバー
	tabBar := a.tabBar.View()
	tabBarHeight := lipgloss.Height(tabBar)

	// ステータスバー
	statusBar := a.statusBar.View()
	statusBarHeight := lipgloss.Height(statusBar)

	// メインコンテンツエリアの高さ
	contentHeight := a.height - tabBarHeight - statusBarHeight
	if contentHeight < 1 {
		contentHeight = 1
	}

	// コンテンツ
	var content string
	if a.historySearch.IsVisible() {
		// Show history search overlay on top of terminal
		if term := a.activeTerminal(); term != nil {
			term.SetSize(a.width, contentHeight)
			termView := term.View()
			searchView := a.historySearch.View()
			// Overlay search on terminal
			content = a.overlayViews(termView, searchView, a.width, contentHeight)
		} else {
			content = a.historySearch.View()
		}
	} else if a.showHelp {
		a.helpModal.SetSize(a.width, contentHeight)
		content = a.helpModal.View()
	} else if term := a.activeTerminal(); term != nil {
		term.SetSize(a.width, contentHeight)
		content = term.View()
	} else {
		content = ""
	}

	// 全体を結合
	tabBarStyled := lipgloss.NewStyle().
		Width(a.width).
		Background(a.ui.Theme.Bg).
		Render(tabBar)

	contentStyled := lipgloss.NewStyle().
		Width(a.width).
		Height(contentHeight).
		Background(a.ui.Theme.Bg).
		Render(content)

	return lipgloss.JoinVertical(
		lipgloss.Left,
		tabBarStyled,
		contentStyled,
		statusBar,
	)
}

// overlayViews overlays the search view on top of the terminal view
func (a *App) overlayViews(base, overlay string, width, height int) string {
	// Simple overlay - show search box at top
	return lipgloss.JoinVertical(lipgloss.Left, overlay)
}

// activeTerminal returns the terminal for the active tab
func (a *App) activeTerminal() *organisms.Terminal {
	return a.terminals[a.tabBar.ActiveTabIndex()]
}

// addNewTab adds a new tab with a terminal
func (a *App) addNewTab() tea.Cmd {
	a.terminalIDCounter++
	id := a.terminalIDCounter

	a.tabBar.AddTab("new", "local")
	term := organisms.NewTerminal(a.ui, id)
	a.terminals[a.tabBar.ActiveTabIndex()] = term

	// Set size if known
	if a.width > 0 && a.height > 0 {
		contentHeight := a.calculateContentHeight()
		term.SetSize(a.width, contentHeight)
	}

	return term.Init()
}

// closeCurrentTab closes the current tab and its terminal
func (a *App) closeCurrentTab() bool {
	idx := a.tabBar.ActiveTabIndex()
	if term, ok := a.terminals[idx]; ok {
		_ = term.Close()
		delete(a.terminals, idx)
	}

	if a.tabBar.CloseTab() {
		// Save history before quitting
		_ = a.history.Save()
		return true // Should quit
	}

	// Re-index terminals after closing
	a.reindexTerminals()
	return false
}

// closeAllTerminals closes all terminal sessions
func (a *App) closeAllTerminals() {
	for _, term := range a.terminals {
		_ = term.Close()
	}
}

// reindexTerminals updates terminal indices after tab changes
func (a *App) reindexTerminals() {
	// This is a simplified approach - in production,
	// we'd need more sophisticated tab/terminal management
}

// calculateContentHeight calculates the content area height
func (a *App) calculateContentHeight() int {
	tabBarHeight := 1 // Approximate
	statusBarHeight := 1 // Approximate

	contentHeight := a.height - tabBarHeight - statusBarHeight
	if contentHeight < 1 {
		contentHeight = 1
	}
	return contentHeight
}

// AddToHistory adds a command to the history
func (a *App) AddToHistory(cmd string) {
	a.history.Add(cmd)
}
