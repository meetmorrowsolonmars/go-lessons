package main

import (
	"context"
	"errors"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/meetmorrowsolonmars/go-lessons/testing/coverage/internal/api"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("POST /numbers:is_even", api.IsEvenHandleFunc)

	server := &http.Server{
		Addr:    ":8080",
		Handler: mux,
	}

	go func() {

		if err := server.ListenAndServe(); err != nil {
			if !errors.Is(err, http.ErrServerClosed) {
				slog.Error("failed to start server", "error", err)
				os.Exit(1)
			}
			slog.Info("server shutdown gracefully")
		}
	}()

	ctx, cancel := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer cancel()

	<-ctx.Done()
	_ = server.Shutdown(context.Background())
}
