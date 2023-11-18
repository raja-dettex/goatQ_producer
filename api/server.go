package api

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/raja-dettex/goatQ_producer/server"
)

type DTO struct {
	Value string
}

type ServerOpts struct {
	ListenAddr string
}

type APIServer struct {
	producer server.Producer
	opts     ServerOpts
}

func NewAPIServer(opts ServerOpts, producer *server.GoatQProducer) *APIServer {
	return &APIServer{
		producer: producer,
		opts:     opts,
	}
}

func (server *APIServer) Start() error {
	server.RegisterHandlers()
	return http.ListenAndServe(server.opts.ListenAddr, nil)
}

func (server *APIServer) RegisterHandlers() {
	http.HandleFunc("/", server.handlePublishMessage)
}

func (server *APIServer) handlePublishMessage(w http.ResponseWriter, r *http.Request) {
	dto := &DTO{}
	if err := json.NewDecoder(r.Body).Decode(dto); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(fmt.Sprintf("error %v", err)))
	}
	if err := server.producer.PutToChannel([]byte(dto.Value)); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(fmt.Sprintf("error %v", err)))
	}
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("buffered to put to queue"))
}
