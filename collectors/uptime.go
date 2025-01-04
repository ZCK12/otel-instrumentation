package collectors

import (
    "context"
    "time"

    "go.opentelemetry.io/otel/metric"
)

var startTime = time.Now()

// Register an ObservableGauge for service_uptime_seconds
func RegisterUptimeMetricsCollector(meter metric.Meter) {
    meter.Float64ObservableGauge(
        "service_uptime_seconds",
        metric.WithFloat64Callback(
            func(ctx context.Context, obs metric.Float64Observer) error {
                uptime := time.Since(startTime).Seconds()
                obs.Observe(uptime)
                return nil
            },
        ),
    )
}
