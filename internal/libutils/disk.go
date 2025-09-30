package libutils

import (
	"fmt"

	"github.com/shirou/gopsutil/v4/disk"
)

func formatDiskBytes(bytes uint64) string {
	const (
		KB = 1024
		MB = 1024 * KB
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
	case bytes >= KB:
		return fmt.Sprintf("%.2f KB", float64(bytes)/float64(KB))
	default:
		return fmt.Sprintf("%d Bytes", bytes)
	}
}

func GetDiskTotal(path string) (string, error) {
	usage, err := disk.Usage(path)
	if err != nil {
		return "", err
	}
	return formatDiskBytes(usage.Total), nil
}

func GetDiskUsed(path string) (string, error) {
	usage, err := disk.Usage(path)
	if err != nil {
		return "", err
	}
	return formatDiskBytes(usage.Used), nil
}

func GetDiskFree(path string) (string, error) {
	usage, err := disk.Usage(path)
	if err != nil {
		return "", err
	}
	return formatDiskBytes(usage.Free), nil
}

func GetDiskUsagePercent(path string) (string, error) {
	usage, err := disk.Usage(path)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%.2f%%", usage.UsedPercent), nil
}
