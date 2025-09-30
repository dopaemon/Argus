package server

import (
	"context"
	"github.com/dopaemon/artus/internal/db"
	"github.com/dopaemon/artus/internal/gRPC/metrics"
	"google.golang.org/grpc/metadata"
)

type MetricsServer struct {
	metrics.UnimplementedMetricsServiceServer
}

func (s *MetricsServer) SendMetrics(ctx context.Context, req *metrics.MetricsRequest) (*metrics.MetricsResponse, error) {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return &metrics.MetricsResponse{Message: "No metadata received"}, nil
	}

	apiKeys := md.Get("x-api-key")
	if len(apiKeys) == 0 {
		return &metrics.MetricsResponse{Message: "Missing API Key"}, nil
	}
	apiKey := apiKeys[0]

	users, _ := db.GetAllUsers()
	if len(users) == 0 || users[0].APIKey != apiKey {
		return &metrics.MetricsResponse{Message: "Invalid API Key"}, nil
	}

	clientIP := req.ClientIp
	if clientIP == "" {
		clientIP = req.Hostname
	}

	metric := &db.Metrics{
		ClientIP:      clientIP,
		CPUName:       req.CpuName,
		LogicalCore:   req.LogicalCore,
		PhysicalCore:  req.PhysicalCore,
		CPUUsage:      req.CpuUsage,
		TotalRAM:      req.TotalRam,
		UsedRAM:       req.UsedRam,
		FreeRAM:       req.FreeRam,
		RAMUsage:      req.RamUsage,
		DiskTotal:     req.DiskTotal,
		DiskUsed:      req.DiskUsed,
		DiskFree:      req.DiskFree,
		DiskUsage:     req.DiskUsage,
		Inbound:       req.Inbound,
		Outbound:      req.Outbound,
		PacketsIn:     req.PacketsIn,
		PacketsOut:    req.PacketsOut,
		Hostname:      req.Hostname,
		OS:            req.Os,
		Platform:      req.Platform,
		KernelVersion: req.KernelVersion,
		Uptime:        req.Uptime,
		BootTime:      req.BootTime,
	}

	_ = db.SaveMetrics(metric)

	return &metrics.MetricsResponse{Message: "Metrics saved successfully"}, nil
}
