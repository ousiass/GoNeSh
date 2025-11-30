// Package history provides command history management for GoNeSh.
package history

import (
	"bufio"
	"os"
	"path/filepath"
	"strings"
	"sync"
)

const (
	// DefaultMaxEntries is the default maximum number of history entries
	DefaultMaxEntries = 10000
	// HistoryFileName is the name of the history file
	HistoryFileName = "history"
)

// History manages command history
type History struct {
	entries    []string
	maxEntries int
	filePath   string
	mu         sync.RWMutex
	position   int // Current position for navigation
}

// New creates a new History instance
func New(maxEntries int) *History {
	if maxEntries <= 0 {
		maxEntries = DefaultMaxEntries
	}

	h := &History{
		entries:    make([]string, 0),
		maxEntries: maxEntries,
		position:   -1,
	}

	// Set default file path
	h.filePath = h.defaultFilePath()

	return h
}

// defaultFilePath returns the default history file path
func (h *History) defaultFilePath() string {
	home, err := os.UserHomeDir()
	if err != nil {
		return ""
	}
	return filepath.Join(home, ".gonesh", HistoryFileName)
}

// Load loads history from the file
func (h *History) Load() error {
	h.mu.Lock()
	defer h.mu.Unlock()

	if h.filePath == "" {
		return nil
	}

	file, err := os.Open(h.filePath)
	if err != nil {
		if os.IsNotExist(err) {
			return nil // No history file yet
		}
		return err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	h.entries = make([]string, 0)

	for scanner.Scan() {
		line := scanner.Text()
		if line != "" {
			h.entries = append(h.entries, line)
		}
	}

	// Trim to max entries
	if len(h.entries) > h.maxEntries {
		h.entries = h.entries[len(h.entries)-h.maxEntries:]
	}

	h.position = len(h.entries)
	return scanner.Err()
}

// Save saves history to the file
func (h *History) Save() error {
	h.mu.RLock()
	defer h.mu.RUnlock()

	if h.filePath == "" {
		return nil
	}

	// Ensure directory exists
	dir := filepath.Dir(h.filePath)
	if err := os.MkdirAll(dir, 0700); err != nil {
		return err
	}

	file, err := os.Create(h.filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	writer := bufio.NewWriter(file)
	for _, entry := range h.entries {
		if _, err := writer.WriteString(entry + "\n"); err != nil {
			return err
		}
	}

	return writer.Flush()
}

// Add adds a new entry to the history
func (h *History) Add(entry string) {
	h.mu.Lock()
	defer h.mu.Unlock()

	entry = strings.TrimSpace(entry)
	if entry == "" {
		return
	}

	// Don't add duplicates of the last entry
	if len(h.entries) > 0 && h.entries[len(h.entries)-1] == entry {
		h.position = len(h.entries)
		return
	}

	h.entries = append(h.entries, entry)

	// Trim to max entries
	if len(h.entries) > h.maxEntries {
		h.entries = h.entries[1:]
	}

	h.position = len(h.entries)
}

// Previous returns the previous entry in history
func (h *History) Previous() (string, bool) {
	h.mu.Lock()
	defer h.mu.Unlock()

	if len(h.entries) == 0 {
		return "", false
	}

	if h.position > 0 {
		h.position--
	}

	return h.entries[h.position], true
}

// Next returns the next entry in history
func (h *History) Next() (string, bool) {
	h.mu.Lock()
	defer h.mu.Unlock()

	if len(h.entries) == 0 || h.position >= len(h.entries)-1 {
		h.position = len(h.entries)
		return "", false
	}

	h.position++
	return h.entries[h.position], true
}

// ResetPosition resets the history navigation position
func (h *History) ResetPosition() {
	h.mu.Lock()
	defer h.mu.Unlock()
	h.position = len(h.entries)
}

// Search searches for entries matching the query
func (h *History) Search(query string) []string {
	h.mu.RLock()
	defer h.mu.RUnlock()

	if query == "" {
		return h.entries
	}

	query = strings.ToLower(query)
	var results []string

	// Search from newest to oldest
	for i := len(h.entries) - 1; i >= 0; i-- {
		if strings.Contains(strings.ToLower(h.entries[i]), query) {
			results = append(results, h.entries[i])
		}
	}

	return results
}

// SearchReverse searches for entries matching the query (reverse incremental search)
func (h *History) SearchReverse(query string, startPos int) (string, int, bool) {
	h.mu.RLock()
	defer h.mu.RUnlock()

	if len(h.entries) == 0 {
		return "", -1, false
	}

	query = strings.ToLower(query)

	// Start from current position and go backwards
	start := startPos
	if start < 0 || start >= len(h.entries) {
		start = len(h.entries) - 1
	}

	for i := start; i >= 0; i-- {
		if strings.Contains(strings.ToLower(h.entries[i]), query) {
			return h.entries[i], i, true
		}
	}

	return "", -1, false
}

// Entries returns all history entries
func (h *History) Entries() []string {
	h.mu.RLock()
	defer h.mu.RUnlock()

	result := make([]string, len(h.entries))
	copy(result, h.entries)
	return result
}

// Len returns the number of entries
func (h *History) Len() int {
	h.mu.RLock()
	defer h.mu.RUnlock()
	return len(h.entries)
}

// Clear clears all history
func (h *History) Clear() {
	h.mu.Lock()
	defer h.mu.Unlock()
	h.entries = make([]string, 0)
	h.position = 0
}
