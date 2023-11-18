package server

import (
	"fmt"
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
	conn, err := net.Dial("tcp", producer.opts.Addr)
	if err != nil {
		fmt.Println(err)
	}

	for message := range producer.messageChannel {
		go producer.publish(message, conn)
	}
}

func (producer *GoatQProducer) PutToChannel(message []byte) error {
	go func(message []byte) {
		producer.messageChannel <- message
	}(message)
	return nil
}

func (producer *GoatQProducer) publish(msg []byte, conn net.Conn) {
	fmt.Println("publisher ", string(msg))
	_, err := conn.Write([]byte(fmt.Sprintf("WRITE %s", string(msg))))
	if err != nil {
		fmt.Println(err)
	}
	buff := make([]byte, 1024)
	for {
		n, err := conn.Read(buff)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(string(buff[:n]))
	}
}
