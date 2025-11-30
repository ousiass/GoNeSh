package molecules

import (
	"fmt"

	"github.com/charmbracelet/lipgloss"
	"github.com/ousiass/GoNeSh/internal/ui/atoms"
	"github.com/ousiass/GoNeSh/internal/ui/context"
)

// ResourceType represents a system resource type
type ResourceType int

const (
	ResourceCPU ResourceType = iota
	ResourceMEM
	ResourceGPU
)

// Resource renders a resource meter with label and percentage
func Resource(ctx *context.UI, resourceType ResourceType, percent float64, hasError bool) string {
	var label string
	var color lipgloss.Color
	var meterFunc func(*context.UI, float64) string

	switch resourceType {
	case ResourceCPU:
		label = "CPU"
		color = ctx.Theme.CPU
		meterFunc = atoms.MeterCPU
	case ResourceMEM:
		label = "MEM"
		color = ctx.Theme.MEM
		meterFunc = atoms.MeterMEM
	case ResourceGPU:
		label = "GPU"
		color = ctx.Theme.GPU
		meterFunc = atoms.MeterGPU
	}

	if hasError {
		return atoms.ErrorText(ctx, label+" "+atoms.IconWarning)
	}

	labelText := atoms.Label(ctx, label, color)
	meter := meterFunc(ctx, percent)
	percentText := atoms.TextAlt(ctx, fmt.Sprintf(" %2.0f%%", percent))

	return labelText + " " + meter + percentText
}
