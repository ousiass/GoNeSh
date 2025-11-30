// Package terminal provides PTY-based terminal functionality for GoNeSh.
package terminal

import (
	"io"
	"os"
	"os/exec"
	"sync"
	"syscall"
	"time"

	"github.com/creack/pty"
)

// PTY represents a pseudo-terminal session
type PTY struct {
	cmd    *exec.Cmd
	pty    *os.File
	mu     sync.Mutex
	closed bool
}

// New creates a new PTY session with the default shell
func New() (*PTY, error) {
	shell := os.Getenv("SHELL")
	if shell == "" {
		shell = "/bin/sh"
	}
	return NewWithCommand(shell)
}

// NewWithCommand creates a new PTY session with a specific command
func NewWithCommand(command string, args ...string) (*PTY, error) {
	cmd := exec.Command(command, args...)
	cmd.Env = os.Environ()
	// Create a new process group so we can kill all child processes
	cmd.SysProcAttr = &syscall.SysProcAttr{Setpgid: true}

	// Start the command with a PTY
	ptmx, err := pty.Start(cmd)
	if err != nil {
		return nil, err
	}

	return &PTY{
		cmd: cmd,
		pty: ptmx,
	}, nil
}

// Read reads from the PTY output
func (p *PTY) Read(buf []byte) (int, error) {
	p.mu.Lock()
	if p.closed {
		p.mu.Unlock()
		return 0, io.EOF
	}
	p.mu.Unlock()

	return p.pty.Read(buf)
}

// Write writes to the PTY input
func (p *PTY) Write(data []byte) (int, error) {
	p.mu.Lock()
	if p.closed {
		p.mu.Unlock()
		return 0, io.EOF
	}
	p.mu.Unlock()

	return p.pty.Write(data)
}

// WriteString writes a string to the PTY input
func (p *PTY) WriteString(s string) (int, error) {
	return p.Write([]byte(s))
}

// Resize resizes the PTY to the given dimensions
func (p *PTY) Resize(rows, cols uint16) error {
	p.mu.Lock()
	defer p.mu.Unlock()

	if p.closed {
		return io.EOF
	}

	return pty.Setsize(p.pty, &pty.Winsize{
		Rows: rows,
		Cols: cols,
	})
}

// Close closes the PTY session
func (p *PTY) Close() error {
	p.mu.Lock()
	if p.closed {
		p.mu.Unlock()
		return nil
	}
	p.closed = true
	p.mu.Unlock()

	// Close PTY first to stop any blocking reads
	_ = p.pty.Close()

	// Send SIGTERM to the process group
	if p.cmd.Process != nil {
		// Kill the process group (negative PID)
		_ = syscall.Kill(-p.cmd.Process.Pid, syscall.SIGTERM)

		// Wait with timeout
		done := make(chan error, 1)
		go func() {
			done <- p.cmd.Wait()
		}()

		select {
		case <-done:
			// Process exited normally
		case <-time.After(100 * time.Millisecond):
			// Timeout - force kill
			_ = syscall.Kill(-p.cmd.Process.Pid, syscall.SIGKILL)
			<-done
		}
	}

	return nil
}

// File returns the underlying PTY file descriptor
func (p *PTY) File() *os.File {
	return p.pty
}

// Pid returns the process ID of the command
func (p *PTY) Pid() int {
	if p.cmd.Process == nil {
		return 0
	}
	return p.cmd.Process.Pid
}
