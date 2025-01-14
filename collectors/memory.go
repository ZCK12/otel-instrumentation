package collectors

import (
    "context"
    "runtime"

    "go.opentelemetry.io/otel/metric"
    "go.opentelemetry.io/otel/attribute"
)

// Register an ObservableGauge for process_memory_usage_mb
func RegisterMemoryMetricsCollector(meter metric.Meter, attr []attribute.KeyValue) {
    const Mb = 1024 * 1024

    meter.Float64ObservableGauge(
        "service_memory_usage_mb",
        metric.WithFloat64Callback(
            func(ctx context.Context, obs metric.Float64Observer) error {
                var memStats runtime.MemStats
                runtime.ReadMemStats(&memStats)
                allocatedMB := float64(memStats.Alloc) / float64(Mb)
                obs.Observe(allocatedMB, metric.WithAttributes(attr...))
                return nil
            },
        ),
    )
}
