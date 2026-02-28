package service

import (
	"context"
	"fmt"
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

func (s *DataService) GetAllData(ctx context.Context, username, authtoken string, sections []*pb.DataInformation) (map[string][]*pb.DataItem, error) {
	result := make(map[string][]*pb.DataItem)

	for _, sec := range sections {
		name := sec.GetDatasection()
		total := sec.GetDatalength()
		minChunk := sec.GetMinchunk()
		chunkSize := minChunk

		if name == "stringsData" && total > 700 {
			total = 700
		}

		var items []*pb.DataItem
		start := int32(0)

		// TODO: Optimize the fetching, use time and chunk sizes to increase the request size until optimal request duration is reached.
		// Although for now it seems like the min chuck is the fastest

		for start < total {
			end := start + chunkSize - 1
			if end >= total {
				end = total - 1 // cap so we never request past the last valid index
			}

			r, err := s.client.GetData(ctx, &pb.DataRequest{
				Username:       username,
				Authtoken:      authtoken,
				Datarequested:  name,
				Datastartindex: start,
				Datachunksize:  end - start + 1, // number of items to fetch
			})

			if err != nil {
				return nil, fmt.Errorf("section %s stopped at index %d: %w", name, start, err)
			}

			data := r.GetData()
			if len(data) == 0 {
				break
			}
			items = append(items, data...)
			start += int32(len(data))
		}

		log.Printf("Section=%s fetched %d/%d items", name, len(items), total)
		result[name] = items
	}

	return result, nil
}
