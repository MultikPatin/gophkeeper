package server

import (
	"fmt"
	l "main/internal/logger"
	"main/internal/server/app/proto"
	"main/internal/server/config"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	fmt.Printf("Build version: %s\n", buildVersion)
	fmt.Printf("Build date: %s\n", buildDate)
	fmt.Printf("Build commit: %s\n", buildCommit)

	logger := l.GetLogger()
	defer l.SyncLogger()

	logger.Info("Initializing server...")

	c := config.Parse(logger)

	a, err := proto.NewApp(c, logger)
	if err != nil {
		logger.Fatalw(err.Error(), "event", "initialize application")
		return
	}
	defer a.Close()

	doneCh := make(chan struct{})

	go func() {
		stopChan := make(chan os.Signal, 1)
		signal.Notify(stopChan, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)

		select {
		case <-stopChan:
			logger.Info("Received shutdown signal.")
			a.Close()
		case <-doneCh:
			logger.Info("Application closed normally.")
		}
		close(doneCh)
	}()

	if err := a.StartServer(); err != nil {
		logger.Fatalw(err.Error(), "event", "start server")
	}

	<-doneCh
}
