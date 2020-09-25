package main

import (
	"github.com/yule-labs/ivy/internal/server"
	"github.com/yule-labs/ivy/internal/storage"
	"log"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	storage.InitStorage()

	server.InitServer()

	sig := make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM)
	interrupt := <-sig

	log.Printf("Received system signal: %s. Shutting down redis-puf\n", interrupt)
}
