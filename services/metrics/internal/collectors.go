package internal

import (
    "time"

    "github.com/prometheus/client_golang/prometheus"
    "github.com/prometheus/client_golang/prometheus/promauto"
    "github.com/shirou/gopsutil/v3/cpu"
    "github.com/shirou/gopsutil/v3/disk"
    "github.com/shirou/gopsutil/v3/mem"
)

var (
    cpuUsage = promauto.NewGauge(prometheus.GaugeOpts{
        Name: "system_cpu_usage",
        Help: "CPU Usage Percentage",
    })
    memoryUsage = promauto.NewGauge(prometheus.GaugeOpts{
        Name: "system_memory_usage",
        Help: "Memory Usage Percentage",
    })
    diskUsage = promauto.NewGauge(prometheus.GaugeOpts{
        Name: "system_disk_usage",
        Help: "Disk Usage Percentage",
    })
    systemUptime = promauto.NewGauge(prometheus.GaugeOpts{
        Name: "system_uptime_seconds",
        Help: "System Uptime in Seconds",
    })
    serviceUptime = promauto.NewGaugeVec(
        prometheus.GaugeOpts{
            Name: "service_uptime_seconds",
            Help: "Service Uptime per Component",
        },
        []string{"service"},
    )
)

type MetricsCollector struct {
    startTime time.Time
}

type MetricsData struct {
    System struct {
        CPUUsage    float64 `json:"cpu_usage"`
        MemoryUsage float64 `json:"memory_usage"`
        DiskUsage   float64 `json:"disk_usage"`
        Uptime      float64 `json:"uptime"`
    } `json:"system"`
    Services map[string]float64 `json:"services"`
}

type HealthStatus struct {
    Status    string            `json:"status"`
    Timestamp int64             `json:"timestamp"`
    Checks    map[string]string `json:"checks"`
}

func NewMetricsCollector() *MetricsCollector {
    collector := &MetricsCollector{
        startTime: time.Now(),
    }
    collector.UpdateMetrics()
    return collector
}

func (mc *MetricsCollector) UpdateMetrics() {
    // CPU usage
    cpuPercent, err := cpu.Percent(0, false)
    if err == nil && len(cpuPercent) > 0 {
        cpuUsage.Set(cpuPercent[0])
    }

    // Memory usage
    memInfo, err := mem.VirtualMemory()
    if err == nil {
        memoryUsage.Set(memInfo.UsedPercent)
    }

    // Disk usage
    diskInfo, err := disk.Usage("/")
    if err == nil {
        diskUsage.Set(diskInfo.UsedPercent)
    }

    // System uptime
    uptime := time.Since(mc.startTime).Seconds()
    systemUptime.Set(uptime)

    // Service uptimes
    serviceUptime.WithLabelValues("gpio").Set(uptime)
    serviceUptime.WithLabelValues("auth").Set(uptime)
    serviceUptime.WithLabelValues("metrics").Set(uptime)
}

func (mc *MetricsCollector) GetMetrics() MetricsData {
    var data MetricsData
    data.System.CPUUsage = cpuUsage.Get()
    data.System.MemoryUsage = memoryUsage.Get()
    data.System.DiskUsage = diskUsage.Get()
    data.System.Uptime = systemUptime.Get()

    data.Services = make(map[string]float64)
    data.Services["gpio"] = serviceUptime.WithLabelValues("gpio").Get()
    data.Services["auth"] = serviceUptime.WithLabelValues("auth").Get()
    data.Services["metrics"] = serviceUptime.WithLabelValues("metrics").Get()

    return data
}

func (mc *MetricsCollector) GetHealth() HealthStatus {
    memInfo, _ := mem.VirtualMemory()
    cpuPercent, _ := cpu.Percent(0, false)
    diskInfo, _ := disk.Usage("/")

    status := "healthy"
    checks := make(map[string]string)

    checks["memory"] = "ok"
    checks["cpu"] = "ok"
    checks["disk"] = "ok"

    if memInfo != nil && memInfo.UsedPercent > 90 {
        status = "degraded"
        checks["memory"] = "warning"
    }

    if len(cpuPercent) > 0 && cpuPercent[0] > 90 {
        status = "degraded"
        checks["cpu"] = "warning"
    }

    if diskInfo != nil && diskInfo.UsedPercent > 90 {
        status = "degraded"
        checks["disk"] = "warning"
    }

    return HealthStatus{
        Status:    status,
        Timestamp: time.Now().Unix(),
        Checks:    checks,
    }
}
