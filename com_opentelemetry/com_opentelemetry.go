package com_opentelemetry

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/exporters/jaeger"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracegrpc"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracehttp"
	stdout "go.opentelemetry.io/otel/exporters/stdout/stdouttrace"
	"go.opentelemetry.io/otel/propagation"
	"go.opentelemetry.io/otel/sdk/resource"
	tracesdk "go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.20.0"
)

type Opentelemetry struct {
	Service     string // 服务名称
	Environment string // 环境

	// 收集枚举类
	Collector Collector
	URI       string // 收集枚举类对应的URI
}

// Collector 没有默认值，必须指定
type Collector int

const (
	// OtlpGrpc accept OpenTelemetry Protocol (OTLP) over gRPC
	// localhost:4317
	// internal.span.format=proto
	OtlpGrpc Collector = iota + 1

	// OtlpHttp accept OpenTelemetry Protocol (OTLP) over HTTP
	// localhost:4318
	// internal.span.format=proto
	OtlpHttp

	// Jaeger accept jaeger.thrift directly from clients
	// http://localhost:14268/api/traces
	// internal.span.format=jaeger
	Jaeger

	// Stdout Create stdout exporter to be able to retrieve
	// 开发，测试
	Stdout
)

func Init(opentelemetry *Opentelemetry) (cancel func(), err error) {
	tp, err := tracerProvider(opentelemetry)
	if err != nil {
		return nil, err
	}
	// Register our TracerProvider as the global so any imported
	// instrumentation in the future will default to using it.
	otel.SetTracerProvider(tp)
	return func() {
		if err := tp.Shutdown(context.Background()); err != nil {
			log.Printf("Error shutting down tracer provider: %v", err)
		}

	}, nil
}

func tracerProvider(opentelemetry *Opentelemetry) (*tracesdk.TracerProvider, error) {
	var exp tracesdk.SpanExporter
	var err error
	switch opentelemetry.Collector {
	case OtlpGrpc:
		// Create the otlptracegrpc exporter
		ctx := context.Background()
		opts := []otlptracegrpc.Option{
			otlptracegrpc.WithInsecure(),
			otlptracegrpc.WithEndpoint(opentelemetry.URI),
			otlptracegrpc.WithReconnectionPeriod(50 * time.Millisecond),
		}
		client := otlptracegrpc.NewClient(opts...)
		exp, err = otlptrace.New(ctx, client)
		if err != nil {
			return nil, fmt.Errorf("creating OTLP gRPC trace exporter: %w", err)
		}
	case OtlpHttp:
		// Create the otlptracehttp exporter
		ctx := context.Background()
		opts := []otlptracehttp.Option{
			otlptracehttp.WithInsecure(),
			otlptracehttp.WithEndpoint(opentelemetry.URI),
		}
		client := otlptracehttp.NewClient(opts...)
		exp, err = otlptrace.New(ctx, client)
		if err != nil {
			return nil, fmt.Errorf("creating OTLP HTTP trace exporter: %w", err)
		}
	case Jaeger:
		// Create the Jaeger exporter
		exp, err = jaeger.New(jaeger.WithCollectorEndpoint(jaeger.WithEndpoint(opentelemetry.URI)))
		if err != nil {
			return nil, fmt.Errorf("creating Jaeger trace exporter: %w", err)
		}
	case Stdout:
		// Create stdout exporter to be able to retrieve
		// the collected spans.
		exp, err = stdout.New(stdout.WithPrettyPrint())
		if err != nil {
			return nil, fmt.Errorf("creating Stdout trace exporter: %w", err)
		}
	default:
		return nil, fmt.Errorf("collector must set")
	}

	tp := tracesdk.NewTracerProvider(
		// Always be sure to batch in production.
		tracesdk.WithBatcher(exp),
		// Record information about this application in a Resource.
		tracesdk.WithResource(resource.NewWithAttributes(
			semconv.SchemaURL,
			semconv.ServiceNameKey.String(opentelemetry.Service),
			attribute.String("environment", opentelemetry.Environment),
			attribute.Int("pid", os.Getpid()),
		)),
	)
	otel.SetTextMapPropagator(propagation.NewCompositeTextMapPropagator(propagation.TraceContext{}, propagation.Baggage{}))
	return tp, nil
}
