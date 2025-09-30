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
}
