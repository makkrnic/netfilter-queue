package main

import (
	"fmt"
	"net"

	clogger "github.com/sartura/netfilter-queue/common/lib/logger"

	"github.com/sartura/netfilter-queue/src/config"
	"go.uber.org/zap"
)

func sendResponse(conn *net.UDPConn, addr *net.UDPAddr) {
	_, err := conn.WriteToUDP([]byte("From server: Message received "), addr)
	if err != nil {
		fmt.Printf("Couldn't send response %v", err)
	}
}

func main() {
	log := clogger.NewProduction()

	cfg, err := config.Load(log)
	if err != nil {
		log.Fatal("error loading config file", zap.Error(err))
	}

	p := make([]byte, 2048)
	addr := net.UDPAddr{
		Port: 67,
		IP:   net.ParseIP(cfg.Address),
	}
	ser, err := net.ListenUDP("udp", &addr)
	if err != nil {
		log.Error("error listening", zap.Error(err))
		return
	}
	for {
		_, remoteaddr, err := ser.ReadFromUDP(p)
		log.Info("message received", zap.Any("address", remoteaddr), zap.Binary("message", p))
		if err != nil {
			log.Error("error handling request", zap.Error(err))
			continue
		}
		go sendResponse(ser, remoteaddr)
	}
}
