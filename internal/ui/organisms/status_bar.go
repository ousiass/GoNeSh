// Package organisms provides complex UI components for GoNeSh.
package organisms

import (
	"fmt"
	"time"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/ousiass/GoNeSh/internal/monitor"
	"github.com/ousiass/GoNeSh/internal/ui/atoms"
	"github.com/ousiass/GoNeSh/internal/ui/context"
	"github.com/ousiass/GoNeSh/internal/ui/molecules"
	"github.com/ousiass/GoNeSh/internal/ui/templates"
)

// tickMsg is sent periodically to update resource info
type tickMsg time.Time

// resourceMsg carries resource information
type resourceMsg monitor.Resources

// StatusBar represents the bottom status bar
type StatusBar struct {
	ctx       *context.UI
	width     int
	resources monitor.Resources
	mode      string // 現在のモード（normal, ai, etc.）
	preset    string // 現在のプリセット名
	env       string // 環境（local, dev, prod）
}

// NewStatusBar creates a new status bar
func NewStatusBar(ctx *context.UI) *StatusBar {
	return &StatusBar{
		ctx:  ctx,
		mode: "normal",
		env:  "local",
	}
}

// Init initializes the status bar
func (s *StatusBar) Init() tea.Cmd {
	return tea.Batch(
		s.tick(),
		s.fetchResources(),
	)
}

// SetWidth sets the status bar width
func (s *StatusBar) SetWidth(width int) {
	s.width = width
}

// SetEnv sets the environment indicator
func (s *StatusBar) SetEnv(env string) {
	s.env = env
}

// SetPreset sets the current preset name
func (s *StatusBar) SetPreset(preset string) {
	s.preset = preset
}

// Update handles messages for the status bar
func (s *StatusBar) Update(msg tea.Msg) (*StatusBar, tea.Cmd) {
	switch msg := msg.(type) {
	case tickMsg:
		return s, tea.Batch(s.tick(), s.fetchResources())
	case resourceMsg:
		s.resources = monitor.Resources(msg)
	}
	return s, nil
}

// View renders the status bar
func (s *StatusBar) View() string {
	if s.width == 0 {
		return ""
	}

	// Left side: Resource meters
	cpuMeter := molecules.Resource(s.ctx, molecules.ResourceCPU, s.resources.CPU, s.resources.CPUError != nil)
	memMeter := molecules.Resource(s.ctx, molecules.ResourceMEM, s.resources.MEM, s.resources.MEMError != nil)

	sep := atoms.Separator(s.ctx)
	left := lipgloss.JoinHorizontal(lipgloss.Top, cpuMeter, sep, memMeter)

	// GPUs (multiple)
	for i, gpu := range s.resources.GPUs {
		label := "GPU"
		if len(s.resources.GPUs) > 1 {
			label = fmt.Sprintf("GPU%d", i)
		}
		gpuMeter := s.renderGPU(label, gpu.Percent, s.resources.GPUError != nil)
		left = lipgloss.JoinHorizontal(lipgloss.Top, left, sep, gpuMeter)
	}

	// Right side: Environment and preset
	right := atoms.PresetBadge(s.ctx, s.preset) + atoms.EnvBadge(s.ctx, s.env)

	// Calculate middle fill
	leftWidth := lipgloss.Width(left)
	rightWidth := lipgloss.Width(right)
	availableWidth := s.width - 2 // Account for padding

	middleWidth := availableWidth - leftWidth - rightWidth
	if middleWidth < 0 {
		middleWidth = 0
	}

	middle := atoms.Fill(s.ctx, middleWidth)
	content := lipgloss.JoinHorizontal(lipgloss.Top, left, middle, right)

	return templates.BottomBar(s.ctx, content, s.width)
}

// renderGPU renders a GPU meter with custom label
func (s *StatusBar) renderGPU(label string, percent float64, hasError bool) string {
	if hasError {
		return atoms.ErrorText(s.ctx, label+" "+atoms.IconWarning)
	}

	labelText := atoms.Label(s.ctx, label, s.ctx.Theme.GPU)
	meter := atoms.MeterGPU(s.ctx, percent)
	percentText := atoms.TextAlt(s.ctx, fmt.Sprintf(" %2.0f%%", percent))

	return labelText + " " + meter + percentText
}

// tick returns a command that triggers a tick after 2 seconds
func (s *StatusBar) tick() tea.Cmd {
	return tea.Tick(2*time.Second, func(t time.Time) tea.Msg {
		return tickMsg(t)
	})
}

// fetchResources returns a command that fetches system resources
func (s *StatusBar) fetchResources() tea.Cmd {
	return func() tea.Msg {
		return resourceMsg(monitor.Fetch())
	}
}
