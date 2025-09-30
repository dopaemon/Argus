package libutils

import (
	"fmt"

	"github.com/shirou/gopsutil/v4/mem"
)

func formatBytes(bytes uint64) string {
	const (
		MB = 1024 * 1024
		GB = 1024 * MB
		TB = 1024 * GB
	)

	switch {
	case bytes >= TB:
		return fmt.Sprintf("%.2f TB", float64(bytes)/float64(TB))
	case bytes >= GB:
		return fmt.Sprintf("%.2f GB", float64(bytes)/float64(GB))
	case bytes >= MB:
		return fmt.Sprintf("%.2f MB", float64(bytes)/float64(MB))
	default:
		return fmt.Sprintf("%d Bytes", bytes)
	}
}

func GetTotalRAM() (string, error) {
	vm, err := mem.VirtualMemory()
	if err != nil {
		return "", err
	}
	return formatBytes(vm.Total), nil
}

func GetUsedRAM() (string, error) {
	vm, err := mem.VirtualMemory()
	if err != nil {
		return "", err
	}
	return formatBytes(vm.Used), nil
}

func GetFreeRAM() (string, error) {
	vm, err := mem.VirtualMemory()
	if err != nil {
		return "", err
	}
	return formatBytes(vm.Available), nil
}

func GetRAMUsagePercent() (string, error) {
	vm, err := mem.VirtualMemory()
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%.2f%%", vm.UsedPercent), nil
}
