package main

import (
	"context"
	"github.com/saratchandra13/sampleProject/config"
	"github.com/saratchandra13/sampleProject/pkg/application/httpserver"
	"github.com/saratchandra13/sampleProject/pkg/domain/services"
	"github.com/saratchandra13/sampleProject/pkg/infrastructure/memory"
	"github.com/saratchandra13/sampleProject/third_party/assetmnger"
	"github.com/saratchandra13/sampleProject/third_party/platlogger"
	"os"
	"os/signal"
	"syscall"
	"time"
)

const (
	serviceName = "sampleProject"
)

func main() {
	assetMng := assetmnger.Initialize()
	var config = config.NewConfig(assetMng)
	logger, _ := platlogger.NewLogger(serviceName, config, platlogger.ConsoleOutput(true), platlogger.StackDriverOutput(true))
	var memStore = memory.NewMemoryStore()

	var appLogic = services.NewAppLogic(memStore)

	idleConnsClosed := make(chan struct{})
	go func() {
		done := make(chan os.Signal, 1)
		signal.Notify(done, os.Interrupt, syscall.SIGTERM, syscall.SIGINT)
		<-done

		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		httpserver.Shutdown(ctx)
		defer cancel()
		close(idleConnsClosed)
	}()

	httpserver.NewServer(appLogic, logger)
	<-idleConnsClosed
}
