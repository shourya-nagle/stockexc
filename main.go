package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
	"golang.org/x/sync/errgroup"
)

func main() {

	if err := run(); err != nil {
		switch err {
		case context.Canceled:
			log.Fatal("context was canceled")
		case http.ErrServerClosed:
			log.Fatal("server close error")
		default:
			log.Fatalf("canot run service because: %v", err)
		}

	}

}

func run() error {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	stklist := stocklist.NewStockLister(1 * time.Second)
	go startServer()

	errorgroup, errorcontext := errgroup.WithContext(ctx)
	errorgroup.Go(func() error {

		return handleSignals(errorcontext, cancel)
	})
	return nil
}

func startServer() {
	router := gin.Default()
	router.Run("localhost:9090")
}

func handleSignals(ctx context.Context, cancel context.CancelFunc) error {
	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, os.Interrupt, syscall.SIGTERM)
	select {
	case s := <-sigCh:
		log.Printf("got signal %v, stopping", s)
		cancel()
		return nil
	case <-ctx.Done():
		log.Printf("context is done")
		return ctx.Err()
	}
}
