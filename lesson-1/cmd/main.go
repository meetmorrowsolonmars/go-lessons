package main

import (
	"context"
	"errors"
	"log"
	"net/http"
	"os/signal"
	"syscall"

	"github.com/meetmorrowsolonmars/go-lessons/lesson-1/fixture"
	appcard "github.com/meetmorrowsolonmars/go-lessons/lesson-1/internal/app/card"
	appitem "github.com/meetmorrowsolonmars/go-lessons/lesson-1/internal/app/item"
	domaincard "github.com/meetmorrowsolonmars/go-lessons/lesson-1/internal/domain/card"
	domainitem "github.com/meetmorrowsolonmars/go-lessons/lesson-1/internal/domain/item"
	repositories "github.com/meetmorrowsolonmars/go-lessons/lesson-1/internal/repository"
)

func main() {

	// Service
	itemRepository := repositories.NewItemRepository(fixture.Items)
	itemService := domainitem.NewItemService(itemRepository)

	cardRepository := repositories.NewCardRepository()
	cardService := domaincard.NewCardService(cardRepository)

	// Application
	itemServer := appitem.NewItemServerImplementation(itemService)
	cardServer := appcard.NewCardServerImplementation(cardService, itemService)

	// Mux
	mux := http.NewServeMux()

	appitem.RegisterRoutes(mux, itemServer)
	appcard.RegisterRoutes(mux, cardServer)

	// Server
	server := http.Server{
		Addr:    ":8080",
		Handler: mux,
	}

	ctx, cancel := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer cancel()

	go func() {
		log.Printf("server listening at %s", server.Addr)

		err := server.ListenAndServe()
		if err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Fatalf("failed to start server: %v", err)
		}
	}()

	<-ctx.Done()
	_ = server.Shutdown(ctx)
}
