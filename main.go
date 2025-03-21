package main

import (
	"context"
	"log"
	"otel/tempo"
	"time"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/codes"
	"go.opentelemetry.io/otel/trace"
)

func someFunction(ctx context.Context) {
	tracer := otel.Tracer("some-function-tracer")
	_, span := tracer.Start(ctx, "some-function")
	defer span.End()

	// Добавляем атрибуты
	span.SetAttributes(attribute.String("param1", "value1"))

	// Имитация работы
	time.Sleep(50 * time.Millisecond)

	// Добавляем событие
	span.AddEvent("Processing completed", trace.WithAttributes(attribute.Int("items_processed", 10)))

	// Имитация ошибки
	span.SetStatus(codes.Error, "Something went wrong")
}

func main() {
	// Инициализация Tracer Provider
	tp, err := tempo.Initracer()
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		if err := tp.Shutdown(context.Background()); err != nil {
			log.Fatal(err)
		}
	}()

	// Создаем трассировку
	tracer := otel.Tracer("main-tracer")
	ctx, span := tracer.Start(context.Background(), "main-span")
	defer span.End()

	// Добавляем атрибуты
	span.SetAttributes(attribute.String("user_id", "12345"))

	// Имитация работы
	time.Sleep(100 * time.Millisecond)

	// Вызов вложенной функции
	someFunction(ctx)

	// Логируем информацию
	log.Println("Tracing started")
}
