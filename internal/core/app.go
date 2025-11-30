// Package core provides the main application logic for GoNeSh.
package core

import (
	"github.com/charmbracelet/bubbles/help"
	"github.com/charmbracelet/bubbles/spinner"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
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
	welcome   *organisms.Welcome
	helpModal *organisms.HelpModal
	showHelp  bool
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

	return &App{
		config:    cfg,
		ui:        ui,
		keys:      DefaultKeyMap(),
		help:      h,
		tabBar:    organisms.NewTabBar(ui),
		statusBar: organisms.NewStatusBar(ui),
		welcome:   organisms.NewWelcome(ui),
		helpModal: organisms.NewHelpModal(ui),
	}
}

// Init initializes the application
func (a *App) Init() tea.Cmd {
	return tea.Batch(
		a.statusBar.Init(),
		a.welcome.Init(),
		tea.SetWindowTitle("GoNeSh"),
	)
}

// Update handles messages and updates the application state
func (a *App) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmds []tea.Cmd

	switch msg := msg.(type) {
	case tea.KeyMsg:
		// ヘルプ表示中はショートカットキーを直接受け付ける
		if a.showHelp {
			a.showHelp = false
			key := msg.String()
			switch key {
			case "t":
				a.tabBar.AddTab("new", "local")
				return a, nil
			case "w":
				if a.tabBar.CloseTab() {
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

		// 通常時
		switch msg.String() {
		case "ctrl+c", "ctrl+q":
			return a, tea.Quit
		case "?":
			a.showHelp = true
			return a, nil
		case "alt+t":
			a.tabBar.AddTab("new", "local")
		case "alt+w":
			if a.tabBar.CloseTab() {
				return a, tea.Quit
			}
		case "alt+]":
			a.tabBar.NextTab()
		case "alt+[":
			a.tabBar.PrevTab()
		case "alt+a": // AIパネル
		case "alt+p": // プリセット
		case "alt+c": // Claude
		case "alt+x": // 外部AI
		case "alt+f": // ファイル
		case "alt+s": // 転送
		case "alt+r": // API
		case "alt+g": // Git
		}

	case tea.WindowSizeMsg:
		a.width = msg.Width
		a.height = msg.Height
		a.ui.SetSize(msg.Width, msg.Height)
		a.tabBar.SetWidth(msg.Width)
		a.statusBar.SetWidth(msg.Width)
		a.help.Width = msg.Width

	case spinner.TickMsg:
		var cmd tea.Cmd
		a.welcome, cmd = a.welcome.Update(msg)
		cmds = append(cmds, cmd)
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
	a.welcome.SetSize(a.width, contentHeight)
	content := a.welcome.View()

	// ヘルプモーダル
	if a.showHelp {
		a.helpModal.SetSize(a.width, contentHeight)
		content = a.helpModal.View()
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
