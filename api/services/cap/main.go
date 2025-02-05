package main

import (
	"context"
	"os"
	"os/signal"
	"runtime"
	"syscall"

	"github.com/ekediala/chat/foundation/logger"
)

func main() {
	var log *logger.Logger

	traceIDFn := func(ctx context.Context) string {
		return ""
	}

	log = logger.New(os.Stdout, logger.LevelInfo, "CAP", traceIDFn)

	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGINT, syscall.SIGKILL, syscall.SIGTERM)
	defer stop()

	err := run(ctx, log)
	if err != nil {
		log.Error(ctx, "startup", "err", err)
		os.Exit(1)
	}

}

func run(ctx context.Context, log *logger.Logger) error {
	log.Info(ctx, "startup", "GOMAXPROCS", runtime.GOMAXPROCS(0))
	log.Info(ctx, "startup", "status", "started")
	defer log.Info(ctx, "startup", "status", "shutting down")
	<-ctx.Done()
	return nil
}
