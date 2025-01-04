package collectors

import (
    "context"
    "runtime"

    "go.opentelemetry.io/otel/metric"
)

// Register an ObservableGauge for process_goroutine_count
func RegisterGoroutinesMetricsCollector(meter metric.Meter) {
    meter.Int64ObservableGauge(
        "service_goroutine_count",
        metric.WithInt64Callback(
            func(ctx context.Context, obs metric.Int64Observer) error {
                goroutineCount := int64(runtime.NumGoroutine())
                obs.Observe(goroutineCount)
                return nil
            },
        ),
    )
}
