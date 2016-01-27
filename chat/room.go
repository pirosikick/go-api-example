package main

import (
	"github.com/gorilla/websocket"
	"github.com/pirosikick/go-api-example/chat/trace"
	"log"
	"net/http"
	//"os"
)

const (
	socketBufferSize  = 1024
	messageBufferSize = 256
)

type room struct {
	forward chan []byte
	join    chan *client
	leave   chan *client
	clients map[*client]bool
	tracer  *trace.Tracer
}

func newRoom() *room {
	return &room{
		forward: make(chan []byte),
		join:    make(chan *client),
		leave:   make(chan *client),
		clients: make(map[*client]bool),
		//tracer:  trace.New(os.Stdout),
	}
}

func (r *room) run() {
	for {
		select {
		case client := <-r.join:
			r.clients[client] = true
			r.tracer.Trace("New client was joined")
		case client := <-r.leave:
			delete(r.clients, client)
			close(client.send)
			r.tracer.Trace("New client was leaved")
		case msg := <-r.forward:
			for client := range r.clients {
				select {
				case client.send <- msg:
					r.tracer.Trace("Sended to client")
				default:
					delete(r.clients, client)
					close(client.send)
					r.tracer.Trace("Failed to send. clean up client")
				}
			}
		}
	}
}

var upgrader = &websocket.Upgrader{
	ReadBufferSize:  socketBufferSize,
	WriteBufferSize: socketBufferSize,
}

func (r *room) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	socket, err := upgrader.Upgrade(w, req, nil)
	if err != nil {
		log.Fatal("ServeHTTP:", err)
	}
	client := &client{
		socket: socket,
		send:   make(chan []byte, messageBufferSize),
		room:   r,
	}
	r.join <- client
	defer func() { r.leave <- client }()
	go client.write()
	client.read()
}
