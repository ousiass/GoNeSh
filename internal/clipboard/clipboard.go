// Package clipboard provides clipboard operations for GoNeSh.
package clipboard

import (
	"github.com/atotto/clipboard"
)

// Copy copies text to the system clipboard
func Copy(text string) error {
	return clipboard.WriteAll(text)
}

// Paste returns text from the system clipboard
func Paste() (string, error) {
	return clipboard.ReadAll()
}

// IsSupported returns whether clipboard operations are supported
func IsSupported() bool {
	return !clipboard.Unsupported
}
