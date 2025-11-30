// Package monitor provides system resource monitoring.
package monitor

import (
	"os/exec"
	"strconv"
	"strings"
	"time"

	"github.com/shirou/gopsutil/v4/cpu"
	"github.com/shirou/gopsutil/v4/mem"
)

// GPU holds information about a single GPU
type GPU struct {
	Index   int
	Name    string
	Percent float64 // utilization
	MemUsed int64   // MB
	MemTotal int64  // MB
}

// MemPercent returns memory usage percentage
func (g GPU) MemPercent() float64 {
	if g.MemTotal == 0 {
		return 0
	}
	return float64(g.MemUsed) / float64(g.MemTotal) * 100
}

// Resources holds system resource information
type Resources struct {
	CPU      float64
	MEM      float64
	GPUs     []GPU
	CPUError error
	MEMError error
	GPUError error
}

// HasGPU returns true if at least one GPU is available
func (r Resources) HasGPU() bool {
	return len(r.GPUs) > 0
}

// Fetch retrieves current system resource usage
func Fetch() Resources {
	r := Resources{}
	r.CPU, r.CPUError = fetchCPU()
	r.MEM, r.MEMError = fetchMEM()
	r.GPUs, r.GPUError = fetchGPUs()
	return r
}

// HasErrors returns true if any resource failed to fetch
func (r Resources) HasErrors() bool {
	return r.CPUError != nil || r.MEMError != nil || r.GPUError != nil
}

// ErrorString returns a summary of all errors
func (r Resources) ErrorString() string {
	var errs []string
	if r.CPUError != nil {
		errs = append(errs, "CPU: "+r.CPUError.Error())
	}
	if r.MEMError != nil {
		errs = append(errs, "MEM: "+r.MEMError.Error())
	}
	if r.GPUError != nil {
		errs = append(errs, "GPU: "+r.GPUError.Error())
	}
	return strings.Join(errs, "\n")
}

func fetchCPU() (float64, error) {
	percent, err := cpu.Percent(100*time.Millisecond, false)
	if err != nil {
		return 0, err
	}
	if len(percent) == 0 {
		return 0, nil
	}
	return percent[0], nil
}

func fetchMEM() (float64, error) {
	info, err := mem.VirtualMemory()
	if err != nil {
		return 0, err
	}
	if info == nil {
		return 0, nil
	}
	return info.UsedPercent, nil
}

func fetchGPUs() ([]GPU, error) {
	// NVIDIA GPU via nvidia-smi
	out, err := exec.Command("nvidia-smi",
		"--query-gpu=index,name,utilization.gpu,memory.used,memory.total",
		"--format=csv,noheader,nounits").Output()
	if err != nil {
		return nil, nil // GPUなしはエラーではない
	}

	lines := strings.Split(strings.TrimSpace(string(out)), "\n")
	var gpus []GPU

	for _, line := range lines {
		if line == "" {
			continue
		}
		parts := strings.Split(line, ", ")
		if len(parts) < 5 {
			continue
		}

		idx, _ := strconv.Atoi(strings.TrimSpace(parts[0]))
		name := strings.TrimSpace(parts[1])
		util, _ := strconv.ParseFloat(strings.TrimSpace(parts[2]), 64)
		memUsed, _ := strconv.ParseInt(strings.TrimSpace(parts[3]), 10, 64)
		memTotal, _ := strconv.ParseInt(strings.TrimSpace(parts[4]), 10, 64)

		gpus = append(gpus, GPU{
			Index:    idx,
			Name:     name,
			Percent:  util,
			MemUsed:  memUsed,
			MemTotal: memTotal,
		})
	}

	return gpus, nil
}
