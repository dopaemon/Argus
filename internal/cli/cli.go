package cli

import (
	"fmt"
	"github.com/dopaemon/artus/internal/db"
)

func ShowMetricsCLI() {
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
