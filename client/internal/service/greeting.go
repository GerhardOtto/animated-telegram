package service

import (
	"context"
	"log"

	pb "github.com/gerhardotto/animated-telegram/client/backendservice"
)

// GreetingService handles the SayHello RPC.
type GreetingService struct {
	client pb.DataBackendClient
}

func NewGreetingService(client pb.DataBackendClient) *GreetingService {
	return &GreetingService{client: client}
}

func (s *GreetingService) SayHello(ctx context.Context, name string) error {
	r, err := s.client.SayHello(ctx, &pb.HelloRequest{Name: name})
	if err != nil {
		return err
	}
	log.Printf("Greeting: %s", r.GetMessage())
	return nil
}
