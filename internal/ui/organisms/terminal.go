// Package organisms provides complex UI components for GoNeSh.
package organisms

import (
	"strings"
	"sync"
	"time"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/ousiass/GoNeSh/internal/terminal"
	"github.com/ousiass/GoNeSh/internal/ui/context"
)

const (
	// Buffer size for reading PTY output
	readBufferSize = 4096
	// Maximum lines to keep in scrollback
	maxScrollback = 10000
)

// ptyOutputMsg carries PTY output data
type ptyOutputMsg struct {
	data []byte
	id   int
}

// ptyErrorMsg signals a PTY error
type ptyErrorMsg struct {
	err error
	id  int
}

// Terminal represents a terminal emulator component
type Terminal struct {
	ctx    *context.UI
	pty    *terminal.PTY
	id     int
	width  int
	height int

	// Output buffer
	lines      []string
	scrollPos  int
	mu         sync.Mutex

	// State
	running bool
	err     error
}

// NewTerminal creates a new terminal component
func NewTerminal(ctx *context.UI, id int) *Terminal {
	return &Terminal{
		ctx:   ctx,
		id:    id,
		lines: []string{},
	}
}

// Init initializes the terminal and starts the shell
func (t *Terminal) Init() tea.Cmd {
	return t.startShell()
}

// startShell starts a new shell session
func (t *Terminal) startShell() tea.Cmd {
	return func() tea.Msg {
		pty, err := terminal.New()
		if err != nil {
			return ptyErrorMsg{err: err, id: t.id}
		}
		t.mu.Lock()
		t.pty = pty
		t.running = true
		t.mu.Unlock()

		// Set initial size
		if t.width > 0 && t.height > 0 {
			_ = pty.Resize(uint16(t.height), uint16(t.width))
		}

		// Start reading output
		go t.readLoop()

		return nil
	}
}

// readLoop continuously reads from the PTY
func (t *Terminal) readLoop() {
	buf := make([]byte, readBufferSize)
	for {
		t.mu.Lock()
		if !t.running || t.pty == nil {
			t.mu.Unlock()
			return
		}
		pty := t.pty
		t.mu.Unlock()

		n, err := pty.Read(buf)
		if err != nil {
			return
		}
		if n > 0 {
			t.mu.Lock()
			t.processOutput(buf[:n])
			t.mu.Unlock()
		}
	}
}

// processOutput processes raw PTY output and adds to buffer
func (t *Terminal) processOutput(data []byte) {
	// Simple line-based processing
	// TODO: Implement proper ANSI escape sequence handling
	text := string(data)

	// Split by newlines but keep partial lines
	parts := strings.Split(text, "\n")

	for i, part := range parts {
		if i == 0 && len(t.lines) > 0 {
			// Append to the last line
			t.lines[len(t.lines)-1] += part
		} else {
			t.lines = append(t.lines, part)
		}
	}

	// Trim scrollback if needed
	if len(t.lines) > maxScrollback {
		t.lines = t.lines[len(t.lines)-maxScrollback:]
	}

	// Auto-scroll to bottom
	t.scrollPos = len(t.lines)
}

// SetSize sets the terminal size
func (t *Terminal) SetSize(width, height int) {
	t.mu.Lock()
	defer t.mu.Unlock()

	t.width = width
	t.height = height

	if t.pty != nil && t.running {
		_ = t.pty.Resize(uint16(height), uint16(width))
	}
}

// Update handles messages for the terminal
func (t *Terminal) Update(msg tea.Msg) (*Terminal, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		return t.handleKeyInput(msg)
	case ptyOutputMsg:
		if msg.id == t.id {
			// Output is handled in readLoop
			return t, t.waitForOutput()
		}
	case ptyErrorMsg:
		if msg.id == t.id {
			t.err = msg.err
			t.running = false
		}
	}
	return t, nil
}

// handleKeyInput handles keyboard input
func (t *Terminal) handleKeyInput(msg tea.KeyMsg) (*Terminal, tea.Cmd) {
	t.mu.Lock()
	defer t.mu.Unlock()

	if !t.running || t.pty == nil {
		return t, nil
	}

	// Convert key to bytes and send to PTY
	var data []byte
	switch msg.Type {
	case tea.KeyEnter:
		data = []byte("\r")
	case tea.KeyBackspace:
		data = []byte{0x7f}
	case tea.KeyTab:
		data = []byte("\t")
	case tea.KeyEscape:
		data = []byte{0x1b}
	case tea.KeyUp:
		data = []byte("\x1b[A")
	case tea.KeyDown:
		data = []byte("\x1b[B")
	case tea.KeyRight:
		data = []byte("\x1b[C")
	case tea.KeyLeft:
		data = []byte("\x1b[D")
	case tea.KeyCtrlC:
		data = []byte{0x03}
	case tea.KeyCtrlD:
		data = []byte{0x04}
	case tea.KeyCtrlZ:
		data = []byte{0x1a}
	case tea.KeyCtrlL:
		data = []byte{0x0c}
	case tea.KeyRunes:
		data = []byte(string(msg.Runes))
	case tea.KeySpace:
		data = []byte(" ")
	default:
		// For other keys, try to use the string representation
		if msg.String() != "" && len(msg.String()) == 1 {
			data = []byte(msg.String())
		}
	}

	if len(data) > 0 {
		_, _ = t.pty.Write(data)
	}

	return t, nil
}

// waitForOutput returns a command that waits briefly for more output
func (t *Terminal) waitForOutput() tea.Cmd {
	return tea.Tick(10*time.Millisecond, func(time.Time) tea.Msg {
		return ptyOutputMsg{id: t.id}
	})
}

// View renders the terminal
func (t *Terminal) View() string {
	t.mu.Lock()
	defer t.mu.Unlock()

	if t.width == 0 || t.height == 0 {
		return ""
	}

	if t.err != nil {
		return lipgloss.NewStyle().
			Foreground(t.ctx.Theme.Error).
			Render("Error: " + t.err.Error())
	}

	// Calculate visible lines
	visibleLines := t.height
	startLine := 0
	if len(t.lines) > visibleLines {
		startLine = len(t.lines) - visibleLines
	}

	// Build output
	var output strings.Builder
	for i := startLine; i < len(t.lines) && i < startLine+visibleLines; i++ {
		line := t.lines[i]
		// Truncate long lines
		if len(line) > t.width {
			line = line[:t.width]
		}
		output.WriteString(line)
		if i < len(t.lines)-1 {
			output.WriteString("\n")
		}
	}

	// Pad remaining lines
	lineCount := len(t.lines) - startLine
	if lineCount < 0 {
		lineCount = 0
	}
	for i := lineCount; i < visibleLines; i++ {
		output.WriteString("\n")
	}

	return lipgloss.NewStyle().
		Width(t.width).
		Height(t.height).
		Background(t.ctx.Theme.Bg).
		Foreground(t.ctx.Theme.Text).
		Render(output.String())
}

// Close closes the terminal
func (t *Terminal) Close() error {
	t.mu.Lock()
	defer t.mu.Unlock()

	t.running = false
	if t.pty != nil {
		return t.pty.Close()
	}
	return nil
}

// IsRunning returns whether the terminal is running
func (t *Terminal) IsRunning() bool {
	t.mu.Lock()
	defer t.mu.Unlock()
	return t.running
}

// SendInput sends a string to the terminal input
func (t *Terminal) SendInput(input string) {
	t.mu.Lock()
	defer t.mu.Unlock()

	if !t.running || t.pty == nil {
		return
	}

	// Write the input to the PTY
	_, _ = t.pty.Write([]byte(input))
}
