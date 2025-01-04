// telemetry/collectors/memory.go
package collectors

import (
    "context"
    "runtime"
    "time"

    "go.opentelemetry.io/otel/attribute"
    "go.opentelemetry.io/otel/metric"
)

// RegisterMemoryCollector sets up a memory usage gauge that runs on a schedule
// via the OTel SDK’s internal ticker. The collector measures allocated memory in MB.
func RegisterMemoryCollector(meter metric.Meter) {
    const Mb = 1024 * 1024

    // Register an ObservableGauge for process.allocated_memory
    meter.Float64ObservableGauge(
        "process.allocated_memory",
        metric.WithFloat64Callback(
            func(ctx context.Context, obs metric.Float64Observer) error {
                var memStats runtime.MemStats
                runtime.ReadMemStats(&memStats)
                allocatedMB := float64(memStats.Alloc) / float64(Mb)

                obs.Observe(allocatedMB)
                return nil
            },
        ),
    )
}

// You can add more collectors below for CPU, disk usage, etc.
