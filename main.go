package main

import (
	"log"
	"os"

	"github.com/raja-dettex/goatQ_producer/api"
	"github.com/raja-dettex/goatQ_producer/server"
)

func main() {
	addr := os.Getenv("ADDR")
	listenAddr := os.Getenv("LISTEN_ADDR")
	serverOpts := api.ServerOpts{ListenAddr: listenAddr}
	producerOpts := &server.ProducerOpts{Addr: addr}
	producer := server.NewGoatQProducer(*producerOpts)
	go producer.Start()
	server := api.NewAPIServer(serverOpts, producer)
	if err := server.Start(); err != nil {
		log.Fatal(err)
	}
}
