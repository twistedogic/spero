package commands

import (
	"log"
	"os"
	"os/signal"
	"syscall"
)

type StopService interface {
	Stop()
}

func Graceful(stopFuncs ...StopService) error {
	defer os.Exit(0)
	sigc := make(chan os.Signal, 1)
	signal.Notify(sigc,
		syscall.SIGHUP,
		syscall.SIGINT,
		syscall.SIGTERM,
		syscall.SIGQUIT,
		os.Interrupt)
	log.Printf("recieved %v, initialize shutdown", <-sigc)
	for _, stopFunc := range stopFuncs {
		stopFunc.Stop()
	}
	return nil
}
