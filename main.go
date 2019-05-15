package main

import (
	"os"
	"os/signal"

	clogger "github.com/sartura/netfilter-queue/common/lib/logger"
)

func main() {
	log := clogger.NewDevelopment()

	log.Info("starting netfilter")

	queue := NewQueue(0, log)

	go queue.Start()

	sig := make(chan os.Signal, 1)
	signal.Notify(sig, os.Interrupt)
	<-sig

}
