// Package clipboard provides internal clipboard for GoNeSh.
// This clipboard is internal to GoNeSh and works consistently
// regardless of local or SSH environment.
package clipboard

import "sync"

var (
	buffer string
	mu     sync.RWMutex
)

// Copy copies text to the internal clipboard
func Copy(text string) error {
	mu.Lock()
	defer mu.Unlock()
	buffer = text
	return nil
}

// Paste returns text from the internal clipboard
func Paste() (string, error) {
	mu.RLock()
	defer mu.RUnlock()
	return buffer, nil
}

// Clear clears the internal clipboard
func Clear() {
	mu.Lock()
	defer mu.Unlock()
	buffer = ""
}

// IsEmpty returns whether the clipboard is empty
func IsEmpty() bool {
	mu.RLock()
	defer mu.RUnlock()
	return buffer == ""
}
