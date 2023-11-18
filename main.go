package main

import (
	"log"

	"github.com/raja-dettex/goatQ_producer/api"
	"github.com/raja-dettex/goatQ_producer/server"
)

func main() {
	serverOpts := api.ServerOpts{ListenAddr: ":4000"}
	producerOpts := &server.ProducerOpts{Addr: ":3000"}
	producer := server.NewGoatQProducer(*producerOpts)
	go producer.Start()
	server := api.NewAPIServer(serverOpts, producer)
	if err := server.Start(); err != nil {
		log.Fatal(err)
	}
}
