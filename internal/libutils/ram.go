package libutils

import (
	"fmt"

	"github.com/shirou/gopsutil/v4/mem"
)

func GetTotalRAM() (string, error) {
	vm, err := mem.VirtualMemory()
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%.2f GB", float64(vm.Total)/(1024*1024*1024)), nil
}

func GetUsedRAM() (string, error) {
	vm, err := mem.VirtualMemory()
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%.2f GB", float64(vm.Used)/(1024*1024*1024)), nil
}

func GetFreeRAM() (string, error) {
	vm, err := mem.VirtualMemory()
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%.2f GB", float64(vm.Available)/(1024*1024*1024)), nil
}

func GetRAMUsagePercent() (string, error) {
	vm, err := mem.VirtualMemory()
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%.2f%%", vm.UsedPercent), nil
}
