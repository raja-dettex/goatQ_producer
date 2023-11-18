package server

import (
	"fmt"
	"testing"
)

func TestProducer(t *testing.T) {
	opts := ProducerOpts{Addr: ":4000"}
	pro := NewGoatQProducer(opts)
	go pro.Start()
	for i := 0; i < 3; i++ {
		pro.PutToChannel([]byte(fmt.Sprintf("hello %d", i)))
	}
}
