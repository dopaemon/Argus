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

	if usage, err := libutils.GetDiskUsagePercent("/"); err == nil {
		fmt.Println("Disk Usage %:", usage)
	} else {
		fmt.Println("Err:", err)
	}

	iface, err := libutils.GetPrimaryInterface()
	if err != nil {
		fmt.Println("Err:", err)
	}

	if sent, err := libutils.GetNetBytesSent(iface); err == nil {
		fmt.Println("Bytes sent:", sent)
	} else {
		fmt.Println("Err:", err)
	}

	if recv, err := libutils.GetNetBytesRecv(iface); err == nil {
		fmt.Println("Bytes received:", recv)
	} else {
		fmt.Println("Err:", err)
	}

	if psent, err := libutils.GetNetPacketsSent(iface); err == nil {
		fmt.Println("Packets sent:", psent)
	} else {
		fmt.Println("Err:", err)
	}

	if precv, err := libutils.GetNetPacketsRecv(iface); err == nil {
		fmt.Println("Packets received:", precv)
	} else {
		fmt.Println("Err:", err)
	}
}
