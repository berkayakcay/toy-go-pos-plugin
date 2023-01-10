package main

import (
	"fmt"
	"github.com/berkayakcay/toy-pos-plugin/foundation/logger"
	"go.uber.org/zap"
	"os"
	"os/signal"
	"runtime"
	"syscall"
)

func main() {
	// Construct the application logger.
	log, err := logger.New("SALES-API")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	defer log.Sync()

	// Perform the startup and shutdown sequence.
	if err := run(log); err != nil {
		log.Errorw("startup", "ERROR", err)
		log.Sync()
		os.Exit(1)
	}

}

func run(log *zap.SugaredLogger) error {
	log.Infow("startup", "GOMAXPROCS", runtime.GOMAXPROCS(0))

	// Make a channel to listen for an interrupt or terminate signal from the OS.
	// Use a buffered channel because the signal package required it.
	shutdown := make(chan os.Signal, 1)
	signal.Notify(shutdown, syscall.SIGINT, syscall.SIGTERM)

	sig := <-shutdown
	log.Infow("shutdown", "shutdown started", "signal", sig)
	defer log.Infow("shutdown", "status", "shutdown complete", "signal", sig)

	return nil
}
