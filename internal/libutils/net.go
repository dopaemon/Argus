package libutils

import (
	"fmt"

	"github.com/shirou/gopsutil/v4/net"
)

func formatNetBytes(bytes uint64) string {
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

func GetPrimaryInterface() (string, error) {
	interfaces, err := net.Interfaces()
	if err != nil {
		return "", fmt.Errorf("failed to get network interfaces: %v", err)
	}

	for _, iface := range interfaces {
		if len(iface.Addrs) > 0 {
			return iface.Name, nil
		}
	}

	return "", fmt.Errorf("no active network interface found")
}

func GetNetBytesSent(iface string) (string, error) {
	counters, err := net.IOCounters(true)
	if err != nil {
		return "", fmt.Errorf("failed to get network IO counters: %v", err)
	}

	for _, c := range counters {
		if c.Name == iface {
			return formatNetBytes(c.BytesSent), nil
		}
	}

	return "", fmt.Errorf("interface %s not found", iface)
}

func GetNetBytesRecv(iface string) (string, error) {
	counters, err := net.IOCounters(true)
	if err != nil {
		return "", fmt.Errorf("failed to get network IO counters: %v", err)
	}

	for _, c := range counters {
		if c.Name == iface {
			return formatNetBytes(c.BytesRecv), nil
		}
	}

	return "", fmt.Errorf("interface %s not found", iface)
}

func GetNetPacketsSent(iface string) (string, error) {
	counters, err := net.IOCounters(true)
	if err != nil {
		return "", fmt.Errorf("failed to get network IO counters: %v", err)
	}

	for _, c := range counters {
		if c.Name == iface {
			return fmt.Sprintf("%d packets", c.PacketsSent), nil
		}
	}

	return "", fmt.Errorf("interface %s not found", iface)
}

func GetNetPacketsRecv(iface string) (string, error) {
	counters, err := net.IOCounters(true)
	if err != nil {
		return "", fmt.Errorf("failed to get network IO counters: %v", err)
	}

	for _, c := range counters {
		if c.Name == iface {
			return fmt.Sprintf("%d packets", c.PacketsRecv), nil
		}
	}

	return "", fmt.Errorf("interface %s not found", iface)
}
