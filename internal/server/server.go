package server

import (
	"chat_rooms/pkg/global/log"
	"chat_rooms/pkg/protocol"
	"github.com/gogo/protobuf/proto"
)

var MyServer = NewServer()

type Server struct {
	Clients    map[string]*Client
	Broadcast  chan []byte
	Register   chan *Client
	Unregister chan *Client
}

func NewServer() *Server {
	return &Server{
		Clients:    make(map[string]*Client),
		Broadcast:  make(chan []byte),
		Register:   make(chan *Client),
		Unregister: make(chan *Client),
	}
}

func (s *Server) Start() {
	log.Logger.Info("start server", log.Any("start server", "start server..."))
	for {
		select {
		case conn := <-s.Register:
			log.Logger.Info("login", log.Any("login", "new user login in"+conn.Name))
			s.Clients[conn.Name] = conn
			msg := &protocol.Message{
				From:    "System",
				To:      conn.Name,
				Content: "welcome",
			}
			protoMsg, _ := proto.Marshal(msg)
			conn.Send <- protoMsg

		case conn := <-s.Unregister:
			log.Logger.Info("loginout", log.Any("loginout", conn.Name))
			if _, ok := s.Clients[conn.Name]; ok {
				close(conn.Send)
				delete(s.Clients, conn.Name)
			}

		case message := <-s.Broadcast:
			msg := &protocol.Message{}
			proto.Unmarshal(message, msg)
		}
	}
}
