package collectors

import (
    "context"
    "runtime"
    "time"

    "go.opentelemetry.io/otel/attribute"
    "go.opentelemetry.io/otel/metric"
)

// Register an ObservableGauge for process_goroutine_count
func RegisterGoroutinesMetricsCollector(meter metric.Meter) {
    meter.Int64ObservableGauge(
        "service_goroutine_count",
        metric.WithFloat64Callback(
            func(ctx context.Context, obs metric.Int64Observer) error {
                goroutineCount := int64(memStats.NumGoroutine())
                obs.Observe(allocatedMB)
                return nil
            },
        ),
    )
}
