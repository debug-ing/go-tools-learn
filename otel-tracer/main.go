package main

import (
	"context"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/exporters/stdout/stdouttrace"
	"go.opentelemetry.io/otel/sdk/trace"
)

func initTracer(status bool) (func(context.Context) error, error) {
	if status {
		exporter, err := stdouttrace.New(stdouttrace.WithPrettyPrint())
		if err != nil {
			return nil, err
		}
		tp := trace.NewTracerProvider(
			trace.WithBatcher(exporter),
		)
		otel.SetTracerProvider(tp)
		return tp.Shutdown, nil
	} else {
		return nil, nil
	}
}

func main() {
	_, err := initTracer(false)
	if err != nil {
		panic("ERROR")
	}
	r := gin.Default()
	r.GET("/trace", func(c *gin.Context) {
		tr := otel.Tracer("example-tracer")
		_, span := tr.Start(c.Request.Context(), "GET /trace")
		defer span.End()
		span.SetAttributes(attribute.String("user_id", "12345"))
		span.AddEvent("Finished processing request")
		c.String(http.StatusOK, "Tracing example with OpenTelemetry!")
	})
	if err := r.Run(":8081"); err != nil {
		log.Fatalf("ERROR RUN APP IN PORT", err)
	}
}
