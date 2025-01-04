package collectors

import (
    "context"
    "time"

    "go.opentelemetry.io/otel/metric"
    "go.opentelemetry.io/otel/attribute"
)

var startTime = time.Now()

// Register an ObservableGauge for service_uptime_seconds
func RegisterUptimeMetricsCollector(meter metric.Meter, attr []attribute.KeyValue) {
    meter.Float64ObservableGauge(
        "service_uptime_seconds",
        metric.WithFloat64Callback(
            func(ctx context.Context, obs metric.Float64Observer) error {
                uptime := time.Since(startTime).Seconds()
                obs.Observe(uptime, metric.WithAttributes(attr...))
                return nil
            },
        ),
    )
}
