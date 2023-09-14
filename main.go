// go-template is a template for a Go project.
package main

import (
	"context"
	"errors"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	ctx, cancel := signal.NotifyContext(
		context.Background(),
		os.Interrupt,
		syscall.SIGTERM,
	)
	defer cancel()

	configFile := flag.String("config", "./config.yaml", "path to config file")
	flag.Parse()

	conf, err := ReadConfig(*configFile)
	if err != nil {
		log.Fatalf("Failed to read config: %v", err)
	}

	// Create HTTP server
	mux := http.NewServeMux()
	mux.HandleFunc("/", Handle)
	//nolint:gosec
	srv := http.Server{
		Addr:    conf.ListenAddr,
		Handler: mux,
	}

	// Run HTTP server
	go func() {
		log.Printf("Listening on %s...", conf.ListenAddr)
		err := srv.ListenAndServe()
		if err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Printf("HTTP server error: %v", err)
		}
	}()

	// Wait for SIGTERM/SIGINT
	<-ctx.Done()

	// Shutdown gracefully
	ctx, cancel = context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Printf("HTTP server shutdown error: %v", err)
	}
	fmt.Println("Shutdown gracefully")
}
