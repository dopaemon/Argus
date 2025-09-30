package main

import (
	"context"
	"fmt"
	"log"
	"net"

	"github.com/dopaemon/artus/internal/config"
	"github.com/dopaemon/artus/internal/db"
	"github.com/dopaemon/artus/internal/gRPC/metrics"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

type metricsServer struct {
	metrics.UnimplementedMetricsServiceServer
}

func (s *metricsServer) SendMetrics(ctx context.Context, req *metrics.MetricsRequest) (*metrics.MetricsResponse, error) {
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

func showMetricsCLI() {
	for {
		fmt.Println("\n=== MENU ===")
		fmt.Println("1. Xem tất cả client metrics")
		fmt.Println("2. Xem metrics theo IP")
		fmt.Println("3. Thoát")
		fmt.Print("Chọn: ")

		var choice int
		fmt.Scan(&choice)

		switch choice {
		case 1:
			metricsList, err := db.GetAllMetrics()
			if err != nil {
				fmt.Println("Lỗi khi truy vấn metrics:", err)
				continue
			}
			for _, m := range metricsList {
				fmt.Printf("IP: %s | CPU: %s | RAM Usage: %s | Disk Usage: %s | Hostname: %s | OS: %s\n",
					m.ClientIP, m.CPUName, m.RAMUsage, m.DiskUsage, m.Hostname, m.OS)
			}
		case 2:
			fmt.Print("Nhập IP của client: ")
			var ip string
			fmt.Scan(&ip)
			m, err := db.GetMetricsByIP(ip)
			if err != nil {
				fmt.Println("Lỗi hoặc không tìm thấy IP:", err)
				continue
			}
			fmt.Printf("Metrics của %s:\nCPU: %s | LogicalCore: %s | PhysicalCore: %s | CPUUsage: %s\n"+
				"RAM Total: %s | Used: %s | Free: %s | Usage: %s\n"+
				"Disk Total: %s | Used: %s | Free: %s | Usage: %s\n"+
				"Inbound: %s | Outbound: %s | PacketsIn: %s | PacketsOut: %s\n"+
				"Hostname: %s | OS: %s | Platform: %s | Kernel: %s\nUptime: %s | BootTime: %s | UpdatedAt: %s\n",
				m.ClientIP, m.CPUName, m.LogicalCore, m.PhysicalCore, m.CPUUsage,
				m.TotalRAM, m.UsedRAM, m.FreeRAM, m.RAMUsage,
				m.DiskTotal, m.DiskUsed, m.DiskFree, m.DiskUsage,
				m.Inbound, m.Outbound, m.PacketsIn, m.PacketsOut,
				m.Hostname, m.OS, m.Platform, m.KernelVersion,
				m.Uptime, m.BootTime, m.UpdatedAt.UTC().Format("2006-01-02 15:04:05 MST"))
		case 3:
			fmt.Println("Thoát CLI metrics")
			return
		default:
			fmt.Println("Lựa chọn không hợp lệ")
		}
	}
}

func main() {
	db.InitDB()

	var username, password string
	users, err := db.GetAllUsers()
	if err != nil {
		log.Fatal("Lỗi truy vấn user:", err)
	}

	if len(users) == 0 {
		fmt.Println("Chưa có tài khoản, vui lòng đăng ký")
		fmt.Print("Nhập username: ")
		fmt.Scan(&username)
		fmt.Print("Nhập password: ")
		fmt.Scan(&password)

		u, err := db.CreateUser(username, password)
		if err != nil {
			log.Fatal("Không tạo được user:", err)
		}
		fmt.Println("Tạo tài khoản thành công")
		fmt.Println("API Key:", u.APIKey)
	} else {
		fmt.Println("Đăng nhập")
		fmt.Print("Username: ")
		fmt.Scan(&username)
		fmt.Print("Password: ")
		fmt.Scan(&password)

		if db.Authenticate(username, password) {
			fmt.Println("Đăng nhập thành công")
			config.Login = true
			config.APIKey, _ = db.GetAPIKey()
		} else {
			fmt.Println("Sai username hoặc password")
			return
		}
	}

	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	metrics.RegisterMetricsServiceServer(grpcServer, &metricsServer{})

	go grpcServer.Serve(lis)

	showMetricsCLI()
}
