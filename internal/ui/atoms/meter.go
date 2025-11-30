package atoms

import (
	"github.com/charmbracelet/lipgloss"
	"github.com/ousiass/GoNeSh/internal/ui/context"
)

const (
	MeterFilled = "▰"
	MeterEmpty  = "▱"
	MeterWidth  = 5
)

// Meter renders a progress bar with filled and empty segments
func Meter(ctx *context.UI, percent float64, color lipgloss.Color) string {
	filled := int(percent / 100 * float64(MeterWidth))
	if filled > MeterWidth {
		filled = MeterWidth
	}
	if filled < 0 {
		filled = 0
	}

	filledStyle := lipgloss.NewStyle().
		Foreground(color).
		Background(ctx.Theme.Bg)

	emptyStyle := lipgloss.NewStyle().
		Foreground(ctx.Theme.Border).
		Background(ctx.Theme.Bg)

	var filledBar, emptyBar string
	for i := 0; i < filled; i++ {
		filledBar += MeterFilled
	}
	for i := filled; i < MeterWidth; i++ {
		emptyBar += MeterEmpty
	}

	return filledStyle.Render(filledBar) + emptyStyle.Render(emptyBar)
}

// MeterCPU renders a CPU meter
func MeterCPU(ctx *context.UI, percent float64) string {
	return Meter(ctx, percent, ctx.Theme.CPU)
}

// MeterMEM renders a memory meter
func MeterMEM(ctx *context.UI, percent float64) string {
	return Meter(ctx, percent, ctx.Theme.MEM)
}

// MeterGPU renders a GPU meter
func MeterGPU(ctx *context.UI, percent float64) string {
	return Meter(ctx, percent, ctx.Theme.GPU)
}
