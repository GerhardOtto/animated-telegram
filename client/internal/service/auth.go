package service

import (
	"context"
	"log"

	pb "github.com/gerhardotto/animated-telegram/client/backendservice"
)

type AuthService struct {
	client pb.DataBackendClient
}

func NewAuthService(client pb.DataBackendClient) *AuthService {
	return &AuthService{client: client}
}

// GetToken requests an auth token for the given username.
func (s *AuthService) GetToken(ctx context.Context, username string) (string, error) {
	r, err := s.client.GetAuthToken(ctx, &pb.AuthTokenRequest{Username: username})
	if err != nil {
		return "", err
	}
	log.Printf("Auth token received for user: %s", username)
	return r.GetAuthtoken(), nil
}
