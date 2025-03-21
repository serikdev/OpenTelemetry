package tempo

import (
	"context"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracegrpc"
	"go.opentelemetry.io/otel/sdk/resource"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.21.0"
	"google.golang.org/grpc"
)

func Initracer() (*sdktrace.TracerProvider, error) {
	// Создаем OTLP Exporter для отправки данных в Tempo
	exporter, err := otlptracegrpc.New(
		context.Background(),
		otlptracegrpc.WithInsecure(),                   // Используем insecure соединение (без TLS)
		otlptracegrpc.WithEndpoint("localhost:4317"),   // Адрес Tempo
		otlptracegrpc.WithDialOption(grpc.WithBlock()), // Блокирующее соединение
	)
	if err != nil {
		return nil, err
	}

	// Создаем Tracer Provider
	tp := sdktrace.NewTracerProvider(
		sdktrace.WithBatcher(exporter),
		sdktrace.WithResource(resource.NewWithAttributes(
			semconv.SchemaURL,
			semconv.ServiceName("serdar-service"),          // Имя вашего сервиса
			semconv.ServiceVersion("1.0.0"),                // Версия сервиса
			attribute.String("environment", "development"), // Окружение
		)),
	)

	// Устанавливаем Tracer Provider
	otel.SetTracerProvider(tp)

	return tp, nil
}
