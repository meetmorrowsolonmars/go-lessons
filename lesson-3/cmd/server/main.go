package main

import (
	"context"
	"log"
	"net"
	"net/http"
	"os/signal"
	"syscall"

	grpcprom "github.com/grpc-ecosystem/go-grpc-middleware/providers/prometheus"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	balanceimpl "github.com/meetmorrowsolonmars/go-lessons/lesson-3/internal/controllers/grpc/v1/balance"
	"github.com/meetmorrowsolonmars/go-lessons/lesson-3/internal/pb/api/v1/balance"
)

func main() {

	// metrics setup
	metrics := grpcprom.NewServerMetrics(grpcprom.WithServerHandlingTimeHistogram())
	registry := prometheus.NewRegistry()
	registry.MustRegister(metrics)

	// server implementation

	balanceServerImpl := balanceimpl.NewBalanceServerImplementation()

	// grpc server setup
	server := grpc.NewServer(grpc.ChainUnaryInterceptor(metrics.UnaryServerInterceptor()))

	balance.RegisterBalanceServiceServer(server, balanceServerImpl)

	reflection.Register(server)

	go func() {
		mux := runtime.NewServeMux()

		_ = balance.RegisterBalanceServiceHandlerServer(context.Background(), mux, balanceServerImpl)

		err := http.ListenAndServe(":9080", mux)
		if err != nil {
			log.Fatal("failed to start http server:", err)
		}
	}()

	go func() {
		listener, err := net.Listen("tcp", ":9082")
		if err != nil {
			log.Fatal("failed to listen:", err)
		}

		err = server.Serve(listener)
		if err != nil {
			log.Fatal("failed to serve:", err)
		}
	}()

	go func() {
		mux := http.NewServeMux()
		mux.Handle("/metrics", promhttp.HandlerFor(registry, promhttp.HandlerOpts{}))

		err := http.ListenAndServe(":9084", mux)
		if err != nil {
			log.Fatal("failed to start debug server:", err)
		}
	}()

	ctx, cancel := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer cancel()

	<-ctx.Done()

	server.GracefulStop()
}
