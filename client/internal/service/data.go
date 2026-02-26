package service

import (
	"context"
	"log"

	pb "github.com/gerhardotto/animated-telegram/client/backendservice"
)

type DataService struct {
	client pb.DataBackendClient
}

func NewDataService(client pb.DataBackendClient) *DataService {
	return &DataService{client: client}
}

func (s *DataService) GetTypesOfData(ctx context.Context, username, authtoken string) ([]*pb.DataInformation, error) {
	r, err := s.client.GetTypesOfData(ctx, &pb.DataInfoRequest{
		Username:  username,
		Authtoken: authtoken,
	})
	if err != nil {
		return nil, err
	}
	log.Printf("Data info: %s", r.GetInfo())
	for _, d := range r.GetAlldatainfo() {
		log.Printf("  Section=%s Length=%d MinChunk=%d MaxChunk=%d",
			d.GetDatasection(), d.GetDatalength(), d.GetMinchunk(), d.GetMaxchunk())
	}
	return r.GetAlldatainfo(), nil
}
