// Package molecules provides composite UI components built from atoms.
package molecules

import (
	"github.com/ousiass/GoNeSh/internal/ui/atoms"
	"github.com/ousiass/GoNeSh/internal/ui/context"
)

// TabType represents the type of tab
type TabType string

const (
	TabTypeLocal TabType = "local"
	TabTypeSSH   TabType = "ssh"
)

// Tab renders a single tab with icon and name
func Tab(ctx *context.UI, name string, tabType TabType, active bool) string {
	icon := atoms.IconTerminal
	if tabType == TabTypeSSH {
		icon = atoms.IconSSH
	}

	text := icon + " " + name

	if active {
		return atoms.ActiveBadge(ctx, text)
	}
	return atoms.InactiveBadge(ctx, text)
}
