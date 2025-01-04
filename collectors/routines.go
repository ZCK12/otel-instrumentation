package collectors

import (
    "context"
    "runtime"

    "go.opentelemetry.io/otel/metric"
    "go.opentelemetry.io/otel/attribute"
)

// Register an ObservableGauge for process_goroutine_count
func RegisterGoroutinesMetricsCollector(meter metric.Meter, attr []attribute.KeyValue) {
    meter.Int64ObservableGauge(
        "service_goroutine_count",
        metric.WithInt64Callback(
            func(ctx context.Context, obs metric.Int64Observer) error {
                goroutineCount := int64(runtime.NumGoroutine())
                obs.Observe(goroutineCount, metric.WithAttributes(attr...))
                return nil
            },
        ),
    )
}
