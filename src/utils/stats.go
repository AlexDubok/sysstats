package utils

import (
	"fmt"
	"log"

	"github.com/shirou/gopsutil/v4/cpu"
	"github.com/shirou/gopsutil/v4/host"
	"github.com/shirou/gopsutil/v4/mem"
)

// Stats holds system metrics
type Stats struct {
	CPUUsage    float64 `json:"cpu_usage"`
	MemoryUsage float64 `json:"memory_usage"`
	Temperature float64 `json:"temperature"`
	Uptime      string  `json:"uptime"`
}

func GetStats() (*Stats, error) {
	// CPU usage
	cpuUsage, err := cpu.Percent(0, false)
	if err != nil {
		return nil, err
	}
	// Memory usage
	memInfo, err := mem.VirtualMemory()
	if err != nil {
		return nil, err
	}

	// Host uptime
	uptime, err := host.Uptime()
	if err != nil {
		return nil, err
	}

	temperature, err := GetTemperature()
	if err != nil {
		log.Fatalf("Error reading temperature: %v", err)
	}

	return &Stats{
		CPUUsage:    cpuUsage[0],
		MemoryUsage: memInfo.UsedPercent,
		Temperature: temperature,
		Uptime:      fmt.Sprintf("%d hours, %d minutes", uptime/3600, (uptime%3600)/60),
	}, nil
}
