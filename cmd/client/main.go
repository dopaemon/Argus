package main

import (
	"fmt"

	"github.com/dopaemon/artus/internal/libutils"
)

func main() {
	if name, err := libutils.GetCPUName(); err == nil {
		fmt.Println("CPU Name:", name)
	} else {
		fmt.Println("Err:", err)
	}

	if cores, err := libutils.GetLogicalCores(); err == nil {
		fmt.Println("CPU Count Logical Core:", cores)
	} else {
		fmt.Println("Err:", err)
	}

	if cores, err := libutils.GetPhysicalCores(); err == nil {
		fmt.Println("CPU Count Physial Core:", cores)
	} else {
		fmt.Println("Err:", err)
	}

	if usage, err := libutils.GetCPUUsage(); err == nil {
		fmt.Println("CPU Usage:", usage)
	} else {
		fmt.Println("Err:", err)
	}

	if total, err := libutils.GetTotalRAM(); err == nil {
		fmt.Println("Ram Total:", total)
	} else {
		fmt.Println("Err:", err)
	}

	if used, err := libutils.GetUsedRAM(); err == nil {
		fmt.Println("Ram Usage:", used)
	} else {
		fmt.Println("Err:", err)
	}

	if free, err := libutils.GetFreeRAM(); err == nil {
		fmt.Println("Ram Free:", free)
	} else {
		fmt.Println("Err:", err)
	}

	if usage, err := libutils.GetRAMUsagePercent(); err == nil {
		fmt.Println("Ram Usage %:", usage)
	} else {
		fmt.Println("Err:", err)
	}

	if total, err := libutils.GetDiskTotal("/"); err == nil {
		fmt.Println("Disk Space:", total)
	} else {
		fmt.Println("Err:", err)
	}

	if usage, err := libutils.GetDiskUsed("/"); err == nil {
		fmt.Println("Disk Usage:", usage)
	} else {
		fmt.Println("Err:", err)
	}

	if free, err := libutils.GetDiskFree("/"); err == nil {
		fmt.Println("Disk Free:", free)
	} else {
		fmt.Println("Err:", err)
	}
}
