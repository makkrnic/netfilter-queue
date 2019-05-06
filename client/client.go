package main

import (
	"bufio"
	"bytes"
	"fmt"
	"net"

	clogger "bitbucket.org/sartura/netfilter-queue/common/lib/logger"
	"bitbucket.org/sartura/netfilter-queue/src/config"
	"go.uber.org/zap"
)

func main() {
	log := clogger.NewProduction()

	cfg, err := config.Load(log)
	if err != nil {
		log.Fatal("error loading config file", zap.Error(err))
	}

	p := make([]byte, 2048)
	conn, err := net.Dial("udp", fmt.Sprintf("%s:67", cfg.Address))
	if err != nil {
		log.Error("error connecting to udp server", zap.Error(err))
		return
	}

	fmt.Fprintf(conn, "Sending UDP message")
	_, err = bufio.NewReader(conn).Read(p)
	if err == nil {
		log.Info("meessage sent", zap.Binary("message", bytes.NewBuffer(p).Bytes()))
	} else {
		log.Error("error occured reading data", zap.Error(err))
	}
	conn.Close()
}
