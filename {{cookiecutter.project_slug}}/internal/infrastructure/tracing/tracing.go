package tracing

import (
	"{{ cookiecutter.project_slug }}/internal/infrastructure/config"

	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/exporters/jaeger"
	"go.opentelemetry.io/otel/sdk/resource"
	tracesdk "go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.17.0"
)

func InitTracerProvider(tracingConfig *config.TracingConfig) (*tracesdk.TracerProvider, error) {
	exp, err := jaeger.New(jaeger.WithCollectorEndpoint(jaeger.WithEndpoint(tracingConfig.ServiceURL)))
	if err != nil {
		return nil, err
	}
	tp := tracesdk.NewTracerProvider(
		tracesdk.WithBatcher(exp),
		tracesdk.WithResource(resource.NewWithAttributes(
			semconv.SchemaURL,
			semconv.ServiceName(tracingConfig.ServiceName),
			attribute.String("environment", tracingConfig.Environment),
			attribute.Int64("ID", tracingConfig.ServiceID),
		)),
	)

	return tp, nil
}
