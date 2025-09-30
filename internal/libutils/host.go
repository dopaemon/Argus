package libutils

import (
	"fmt"

	"github.com/shirou/gopsutil/v4/host"
)

func GetHostName() (string, error) {
	info, err := host.Info()
	if err != nil {
		return "", fmt.Errorf("failed to get hostname: %v", err)
	}
	return info.Hostname, nil
}

func GetOS() (string, error) {
	info, err := host.Info()
	if err != nil {
		return "", fmt.Errorf("failed to get OS info: %v", err)
	}
	return info.OS, nil
}

func GetPlatform() (string, error) {
	info, err := host.Info()
	if err != nil {
		return "", fmt.Errorf("failed to get platform info: %v", err)
	}
	return info.Platform, nil
}

func GetPlatformVersion() (string, error) {
	info, err := host.Info()
	if err != nil {
		return "", fmt.Errorf("failed to get platform version: %v", err)
	}
	return info.PlatformVersion, nil
}

func GetKernelVersion() (string, error) {
	info, err := host.Info()
	if err != nil {
		return "", fmt.Errorf("failed to get kernel version: %v", err)
	}
	return info.KernelVersion, nil
}

func GetUptime() (string, error) {
	info, err := host.Info()
	if err != nil {
		return "", fmt.Errorf("failed to get uptime: %v", err)
	}
	return fmt.Sprintf("%d seconds", info.Uptime), nil
}

func GetBootTime() (string, error) {
	info, err := host.Info()
	if err != nil {
		return "", fmt.Errorf("failed to get boot time: %v", err)
	}
	return fmt.Sprintf("%d", info.BootTime), nil
}
