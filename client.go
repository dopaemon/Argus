package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"time"

	pb "artus/pb"

	"github.com/shirou/gopsutil/v3/cpu"
	"github.com/shirou/gopsutil/v3/disk"
	_ "github.com/shirou/gopsutil/v3/host"
	"github.com/shirou/gopsutil/v3/mem"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func sampleMetrics() (*pb.Metrics, error) {
	hostname, _ := os.Hostname()

	// CPU percent per CPU (interval 0 -> instantaneous since last call is not available; use a small interval)
	percentages, err := cpu.Percent(500*time.Millisecond, true)
	if err != nil {
		return nil, err
	}
	cpuInfos := make([]*pb.CpuInfo, 0, len(percentages))
	models, _ := cpu.Info()
	cores := len(percentages)
	modelName := ""
	if len(models) > 0 {
		modelName = models[0].ModelName
	}
	for i := 0; i < cores; i++ {
		cpuInfos = append(cpuInfos, &pb.CpuInfo{
			Model:        modelName,
			UsagePercent: percentages[i],
			Cores:        int32(cores),
		})
	}

	vm, err := mem.VirtualMemory()
	if err != nil {
		return nil, err
	}

	du, err := disk.Usage("/")
	if err != nil {
		// try current working directory if root fails
		wd, _ := os.Getwd()
		du, _ = disk.Usage(wd)
	}

	metrics := &pb.Metrics{
		Hostname: hostname,
		Cpus:     cpuInfos,
		Memory: &pb.MemoryInfo{
			Total:       vm.Total,
			Used:        vm.Used,
			Free:        vm.Free,
			UsedPercent: vm.UsedPercent,
		},
		Disk: &pb.DiskInfo{
			Total:       du.Total,
			Used:        du.Used,
			Free:        du.Free,
			UsedPercent: du.UsedPercent,
			Path:        du.Path,
		},
		Ts: timestamppb.Now(),
	}
	return metrics, nil
}

func main() {
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("failed to dial: %v", err)
	}
	defer conn.Close()
	client := pb.NewMonitorServiceClient(conn)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	stream, err := client.Monitor(ctx)
	if err != nil {
		log.Fatalf("Monitor stream error: %v", err)
	}

	// listen for server messages
	go func() {
		for {
			msg, err := stream.Recv()
			if err != nil {
				log.Println("server Recv error (server closed?):", err)
				return
			}
			log.Println("Server message:", msg.GetText(), "command:", msg.GetCommand())
			// (optionally handle commands)
		}
	}()

	// send metrics periodically
	ticker := time.NewTicker(5 * time.Second)
	defer ticker.Stop()

	// Stop on Ctrl+C
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)

loop:
	for {
		select {
		case <-ticker.C:
			metrics, err := sampleMetrics()
			if err != nil {
				log.Println("sampleMetrics error:", err)
				continue
			}
			if err := stream.Send(metrics); err != nil {
				log.Println("stream.Send error:", err)
				break loop
			}
			log.Printf("Sent metrics at %s\n", metrics.GetTs().AsTime().Format(time.RFC3339))
		case <-c:
			log.Println("Interrupted, closing stream")
			// Close send direction; with bidi streams, CloseSend so server sees EOF on receive
			if err := stream.CloseSend(); err != nil {
				log.Println("CloseSend error:", err)
			}
			break loop
		}
	}

	// wait a moment to receive final server messages
	time.Sleep(1 * time.Second)
}
