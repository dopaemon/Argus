package metrics

import (
	"context"
	"log"

	sys "github.com/dopaemon/artus/internal/system"
	netw "github.com/dopaemon/artus/internal/network"
	utils "github.com/dopaemon/artus/internal/utils"
	libutils "github.com/dopaemon/artus/internal/libutils"
	grpcMetrics "github.com/dopaemon/artus/internal/gRPC/metrics"
	"google.golang.org/grpc/metadata"
)

func Send(client grpcMetrics.MetricsServiceClient, apiKey string) {
	iface := netw.PrimaryInterface()

	req := &grpcMetrics.MetricsRequest{
		ClientIp:      utils.GetString(libutils.GetHostIP),
		CpuName:       sys.Name(),
		LogicalCore:   sys.LogicalCores(),
		PhysicalCore:  sys.PhysicalCores(),
		CpuUsage:      sys.CPUUsage(),
		TotalRam:      sys.RAMTotal(),
		UsedRam:       sys.RAMUsed(),
		FreeRam:       sys.RAMFree(),
		RamUsage:      sys.RAMUsage(),
		DiskTotal:     sys.DiskTotal("/"),
		DiskUsed:      sys.DiskUsed("/"),
		DiskFree:      sys.DiskFree("/"),
		DiskUsage:     sys.DiskUsage("/"),
		Inbound:       netw.BytesRecv(iface),
		Outbound:      netw.BytesSent(iface),
		PacketsIn:     netw.PacketsRecv(iface),
		PacketsOut:    netw.PacketsSent(iface),
		Hostname:      utils.GetString(libutils.GetHostName),
		Os:            utils.GetString(libutils.GetOS),
		Platform:      utils.GetString(libutils.GetPlatform),
		KernelVersion: utils.GetString(libutils.GetKernelVersion),
		Uptime:        utils.GetString(libutils.GetUptime),
		BootTime:      utils.GetString(libutils.GetBootTime),
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
