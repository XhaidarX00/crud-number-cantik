package main

import (
	"context"
	"log"
	"main/infra"
	"main/router"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"time"
)

func main() {
	// Test()
	ctx, err := infra.NewIntegrateContext()
	if err != nil {
		log.Fatal("can't init service context %w", err)
	}

	r := router.NewRoutes(ctx)

	srv := &http.Server{
		Addr:    ":8080",
		Handler: r,
	}

	go func() {
		// Start the server
		log.Printf("Server running on port 8080")
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	// Wait for interrupt signal to gracefully shutdown the server with a timeout of 5 seconds
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("Shutdown Server ...")

	// Create a timeout context for graceful shutdown
	shutdownCtx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	// Attempt to gracefully shutdown the server
	if err := srv.Shutdown(shutdownCtx); err != nil {
		log.Fatal("Server Shutdown:", err)
	}

	// Catching context timeout
	select {
	case <-shutdownCtx.Done():
		log.Println("Timeout of 5 seconds.")
	}

	log.Println("Server exiting")
}
