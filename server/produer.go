package server

import (
	"fmt"
	"io"
	"net"
)

type Producer interface {
	PutToChannel([]byte) error
}

type ProducerOpts struct {
	Addr string
}

type GoatQProducer struct {
	opts           ProducerOpts
	messageChannel chan []byte
}

func NewGoatQProducer(opts ProducerOpts) *GoatQProducer {
	return &GoatQProducer{
		opts:           opts,
		messageChannel: make(chan []byte),
	}
}

func (producer *GoatQProducer) Start() {

	for message := range producer.messageChannel {
		go producer.publish(message)
	}
}

func (producer *GoatQProducer) PutToChannel(message []byte) error {
	go func(message []byte) {
		producer.messageChannel <- message
	}(message)
	return nil
}

func (producer *GoatQProducer) publish(msg []byte) {
	fmt.Println("publisher ", string(msg))
	conn, err := net.Dial("tcp", producer.opts.Addr)
	defer conn.Close()
	_, err = conn.Write([]byte(fmt.Sprintf("WRITE %s", string(msg))))
	if err != nil {
		fmt.Println("in publish error", err)
	}
	buff := make([]byte, 1024)
	for {
		n, err := conn.Read(buff)
		if err == io.EOF {
			return
		}
		if _, ok := err.(*net.OpError); ok {
			return
		}
		if err != nil {
			fmt.Println("read from storage node", err)
		}
		fmt.Println(string(buff[:n]))
	}
}
