package telemetry

import (
    "context"
    "fmt"
    "time"

    "go.opentelemetry.io/otel/exporters/otlp/otlpmetric/otlpmetricgrpc"
    "go.opentelemetry.io/otel/metric"
    "go.opentelemetry.io/otel/sdk/resource"
    metricsdk "go.opentelemetry.io/otel/sdk/metric"
)

// SCHEMA_URL holds the schema definition for attributes and resources.
const SCHEMA_URL = "https://opentelemetry.io/schemas/1.26.0"

// MeterConfig includes any parameters you want to allow callers
// to override when creating a MeterProvider.
type MeterConfig struct {
    Endpoint       string
    Insecure       bool
    ExportInterval time.Duration
}

// NewMeterProvider creates and returns a MeterProvider based on the provided config.
// This is the main entry point to set up telemetry in your services.
func NewMeterProvider(ctx context.Context, cfg MeterConfig) (metric.MeterProvider, error) {
    // Create a resource describing this service/application.
    res, err := newResource(ctx)
    if err != nil {
        return nil, fmt.Errorf("could not create resource: %w", err)
    }

    // Create the exporter that sends metrics to your desired endpoint.
    exporter, err := newExporter(ctx, cfg)
    if err != nil {
        return nil, fmt.Errorf("could not create exporter: %w", err)
    }

    // Create a reader that periodically exports metrics.
    reader := metricsdk.NewPeriodicReader(
        exporter,
        metricsdk.WithInterval(cfg.ExportInterval),
    )

    // Build the MeterProvider with the resource and reader.
    provider := metricsdk.NewMeterProvider(
        metricsdk.WithReader(reader),
        metricsdk.WithResource(res),
    )

    return provider, nil
}

// newResource sets important attributes for the telemetry data.
func newResource(ctx context.Context) (*resource.Resource, error) {
    return resource.New(
        ctx,
        resource.WithSchemaURL(SCHEMA_URL),
    )
}

// newExporter creates a gRPC exporter for sending metrics to an OTel Collector.
func newExporter(ctx context.Context, cfg MeterConfig) (metricsdk.Exporter, error) {
    // Graceful time limit for establishing a connection.
    cctx, cancel := context.WithTimeout(ctx, 5*time.Second)
    defer cancel()

    opts := []otlpmetricgrpc.Option{
        otlpmetricgrpc.WithEndpoint(cfg.Endpoint),
        otlpmetricgrpc.WithCompressor("gzip"),
    }
    if cfg.Insecure {
        opts = append(opts, otlpmetricgrpc.WithInsecure())
    }

    exporter, err := otlpmetricgrpc.New(cctx, opts...)
    if err != nil {
        return nil, fmt.Errorf("could not create OTLP metric exporter: %w", err)
    }
    return exporter, nil
}
