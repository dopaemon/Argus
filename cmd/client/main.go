package main

import (
	"context"
	"log"
	"time"

	"github.com/dopaemon/artus/internal/libutils"
	"github.com/dopaemon/artus/internal/gRPC/metrics"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

func getString(fn func() (string, error)) string {
	if val, err := fn(); err == nil {
		return val
	}
	return ""
}

func sendMetrics(client metrics.MetricsServiceClient, apiKey string) {
	iface, _ := libutils.GetPrimaryInterface()

	req := &metrics.MetricsRequest{
		ClientIp:      getString(libutils.GetHostIP),
		CpuName:       getString(libutils.GetCPUName),
		LogicalCore:   getString(libutils.GetLogicalCores),
		PhysicalCore:  getString(libutils.GetPhysicalCores),
		CpuUsage:      getString(libutils.GetCPUUsage),
		TotalRam:      getString(libutils.GetTotalRAM),
		UsedRam:       getString(libutils.GetUsedRAM),
		FreeRam:       getString(libutils.GetFreeRAM),
		RamUsage:      getString(libutils.GetRAMUsagePercent),
		DiskTotal:     getString(func() (string, error) { return libutils.GetDiskTotal("/") }),
		DiskUsed:      getString(func() (string, error) { return libutils.GetDiskUsed("/") }),
		DiskFree:      getString(func() (string, error) { return libutils.GetDiskFree("/") }),
		DiskUsage:     getString(func() (string, error) { return libutils.GetDiskUsagePercent("/") }),
		Inbound:       getString(func() (string, error) { return libutils.GetNetBytesRecv(iface) }),
		Outbound:      getString(func() (string, error) { return libutils.GetNetBytesSent(iface) }),
		PacketsIn:     getString(func() (string, error) { return libutils.GetNetPacketsRecv(iface) }),
		PacketsOut:    getString(func() (string, error) { return libutils.GetNetPacketsSent(iface) }),
		Hostname:      getString(libutils.GetHostName),
		Os:            getString(libutils.GetOS),
		Platform:      getString(libutils.GetPlatform),
		KernelVersion: getString(libutils.GetKernelVersion),
		Uptime:        getString(libutils.GetUptime),
		BootTime:      getString(libutils.GetBootTime),
	}

	ctx := metadata.AppendToOutgoingContext(context.Background(), "x-api-key", apiKey)

	go func() {
		resp, err := client.SendMetrics(ctx, req)
		if err != nil {
			log.Printf("failed to send metrics: %v", err)
		} else {
			log.Printf("metrics sent successfully: %v", resp.Message)
		}
	}()
}

func main() {
	apiKey := "2be97c1902efce4fcd93deea058102574f6a895f323762c9c88d94a0e6ad789b"

	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("failed to connect gRPC server: %v", err)
	}
	defer conn.Close()

	client := metrics.NewMetricsServiceClient(conn)

	ticker := time.NewTicker(time.Second)
	defer ticker.Stop()

	for range ticker.C {
		sendMetrics(client, apiKey)
	}
}
