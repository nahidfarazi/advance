package main

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"strings"
	"time"

	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/disk"
	"github.com/shirou/gopsutil/host"
	"github.com/shirou/gopsutil/mem"
)

func main() {
        if runtime.GOOS != "linux" {
                fmt.Println("This tool is designed for Linux systems.")
                os.Exit(1)
        }

        for {
                clearScreen()

                hostInfo, _ := host.Info()
                fmt.Printf("Hostname: %s\n", hostInfo.Hostname)
                fmt.Printf("Uptime: %s\n", formatUptime(hostInfo.Uptime))
                fmt.Printf("OS: %s %s\n", hostInfo.Platform, hostInfo.PlatformVersion)
                fmt.Printf("Kernel: %s\n", hostInfo.KernelVersion)
                fmt.Println("--------------------")

                cpuPercent, _ := cpu.Percent(time.Second, false)
                fmt.Printf("CPU Usage: %.2f%%\n", cpuPercent[0])
                cpuInfo, _ := cpu.Info()
                fmt.Printf("CPU Model: %s\n", cpuInfo[0].ModelName)
                fmt.Println("--------------------")

                memInfo, _ := mem.VirtualMemory()
                fmt.Printf("Total Memory: %s\n", formatBytes(memInfo.Total))
                fmt.Printf("Available Memory: %s\n", formatBytes(memInfo.Available))
                fmt.Printf("Memory Usage: %.2f%%\n", memInfo.UsedPercent)
                fmt.Println("--------------------")

                diskInfo, _ := disk.Usage("/") // Get usage of the root partition
                fmt.Printf("Total Disk Space: %s\n", formatBytes(diskInfo.Total))
                fmt.Printf("Free Disk Space: %s\n", formatBytes(diskInfo.Free))
                fmt.Printf("Disk Usage: %.2f%%\n", diskInfo.UsedPercent)
                fmt.Println("--------------------")

                // Add network info here if needed using net package

                time.Sleep(1 * time.Second)
        }
}

func clearScreen() {
        cmd := exec.Command("clear") // For Linux/macOS
        cmd.Stdout = os.Stdout
        cmd.Run()
}

func formatBytes(b uint64) string {
        const unit = 1024
        if b < unit {
                return fmt.Sprintf("%d B", b)
        }
        div, exp := int64(unit), 0
        for n := b / unit; n >= unit; n /= unit {
                div *= unit
                exp++
        }
        return fmt.Sprintf("%.1f %ciB", float64(b)/float64(div), "KMGTPE"[exp])
}

func formatUptime(seconds uint64) string {
        days := seconds / (60 * 60 * 24)
        seconds -= days * (60 * 60 * 24)
        hours := seconds / (60 * 60)
        seconds -= hours * (60 * 60)
        minutes := seconds / 60
        seconds -= minutes * 60
        var parts []string
        if days > 0 {
                parts = append(parts, fmt.Sprintf("%d days", days))
        }
        if hours > 0 {
                parts = append(parts, fmt.Sprintf("%d hours", hours))
        }
        if minutes > 0 {
                parts = append(parts, fmt.Sprintf("%d minutes", minutes))
        }
        parts = append(parts, fmt.Sprintf("%d seconds", seconds))

        return strings.Join(parts, ", ")
}