package libutils

import (
	"fmt"
	"time"

	"github.com/shirou/gopsutil/v4/cpu"
)

func GetCPUName() (string, error) {
	info, err := cpu.Info()
	if err != nil {
		return "", err
	}
	if len(info) == 0 {
		return "", fmt.Errorf("Can't Find CPU Name")
	}
	return info[0].ModelName, nil
}

func GetLogicalCores() (string, error) {
	count, err := cpu.Counts(true)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%d", count), nil
}

func GetPhysicalCores() (string, error) {
	count, err := cpu.Counts(false)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%d", count), nil
}

func GetCPUUsage() (string, error) {
	percent, err := cpu.Percent(time.Second, false)
	if err != nil {
		return "", err
	}
	if len(percent) == 0 {
		return "", fmt.Errorf("Can't get CPU Usage")
	}
	return fmt.Sprintf("%.2f%%", percent[0]), nil
}
