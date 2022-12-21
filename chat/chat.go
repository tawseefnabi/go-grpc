package chat

import (
	"log"

	"golang.org/x/net/context"
)

type Server struct {
}

func (s *Server) SayHello(ctx context.Context, message *Message) (*Message, error) {
	log.Printf("Recieved message from client: %s", message.Body)

	return &Message{Body: "Hello from server!"}, nil
}

func (s *Server) BroadCastMessage(ctx context.Context, in *Message) (*Message, error) {
	log.Printf("Broadcasting new message from a client: %s", in.Body)
	return &Message{Body: "Broadcasted message!"}, nil
}
