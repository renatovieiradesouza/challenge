package main

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracehttp"
	"go.opentelemetry.io/otel/sdk/resource"
	semconv "go.opentelemetry.io/otel/semconv/v1.24.0"
	"go.opentelemetry.io/otel/sdk/trace"
)

// resposta é só um alias para deixar JSON rápido.
type resposta map[string]string

func main() {
	// ---- 1) iniciar OpenTelemetry ----
	ctx := context.Background()
	shutdown := initTracer(ctx)
	defer shutdown(ctx)

	// ---- 2) rotas ----
	mux := http.NewServeMux()
	mux.Handle("/texto", otelhttp.NewHandler(http.HandlerFunc(texto), "texto"))
	mux.Handle("/horario", otelhttp.NewHandler(http.HandlerFunc(horario), "horario"))

	// ---- 3) servidor HTTP ----
	server := &http.Server{
		Addr:              ":5000",
		Handler:           cors(mux), // CORS + traces
		ReadHeaderTimeout: 5 * time.Second,
	}
	go func() {
		log.Println("Servidor ouvindo em :5000")
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("erro servidor: %v", err)
		}
	}()

	// ---- 4) graceful shutdown (Ctrl-C) ----
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	_ = server.Shutdown(ctx)
	log.Println("bye")
}

// ---------- handlers ----------

func texto(w http.ResponseWriter, r *http.Request) {
	out := resposta{"mensagem": "Aplicação 1 - Texto fixo"}
	writeJSON(w, out)
}

func horario(w http.ResponseWriter, r *http.Request) {
	now := time.Now().Format("15:04:05 - 02/01/2006")
	out := resposta{"horario": now}
	writeJSON(w, out)
}

func writeJSON(w http.ResponseWriter, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(data)
}

// ---------- CORS (simple) ----------

func cors(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusNoContent)
			return
		}
		next.ServeHTTP(w, r)
	})
}

// ---------- OpenTelemetry ----------

func initTracer(ctx context.Context) func(context.Context) error {
	// endpoint padrão: http://localhost:4318
	endpoint := os.Getenv("OTEL_EXPORTER_OTLP_ENDPOINT")
	if endpoint == "" {
		endpoint = "http://localhost:4318"
	}

	exp, err := otlptracehttp.New(ctx,
		otlptracehttp.WithEndpoint(endpoint),
		otlptracehttp.WithInsecure(), // collector sem TLS no dev
	)
	if err != nil {
		log.Fatalf("erro OTLP exporter: %v", err)
	}

	res, _ := resource.New(ctx,
		resource.WithAttributes(
			semconv.ServiceName("app1-golang"),
			semconv.ServiceVersion("1.0.0"),
			attribute.String("env", "local"),
		),
	)

	tp := trace.NewTracerProvider(
		trace.WithBatcher(exp),
		trace.WithResource(res),
	)
	otel.SetTracerProvider(tp)
	return tp.Shutdown
}
