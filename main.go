package main

import (
	"os"
	"os/signal"

	clogger "bitbucket.org/sartura/netfilter-queue/common/lib/logger"
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
