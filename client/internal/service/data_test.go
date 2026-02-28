package service

import (
	"context"
	"errors"
	"fmt"
	"testing"

	pb "github.com/gerhardotto/animated-telegram/client/backendservice"
	"google.golang.org/grpc"
)

func TestDataService_GetTypesOfData(t *testing.T) {
	tests := []struct {
		name           string
		inputUser      string
		inputToken     string
		mockReply      *pb.DataInfoReply
		mockErr        error
		wantErr        bool
		wantErrMsg     string
		wantSectionLen int
	}{
		{
			name:       "success",
			inputUser:  "alice",
			inputToken: "tok-abc",
			mockReply: &pb.DataInfoReply{
				Info: "ok",
				Alldatainfo: []*pb.DataInformation{
					{Datasection: "intData", Datalength: 100, Minchunk: 10, Maxchunk: 50},
					{Datasection: "stringsData", Datalength: 200, Minchunk: 20, Maxchunk: 100},
				},
			},
			mockErr:        nil,
			wantErr:        false,
			wantSectionLen: 2,
		},
		{
			name:       "rpc error",
			inputUser:  "bob",
			inputToken: "tok-xyz",
			mockReply:  nil,
			mockErr:    errors.New("rpc failed"),
			wantErr:    true,
			wantErrMsg: "rpc failed",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var capturedUser, capturedToken string
			mock := &mockDataBackendClient{
				getTypesOfData: func(ctx context.Context, in *pb.DataInfoRequest, opts ...grpc.CallOption) (*pb.DataInfoReply, error) {
					capturedUser = in.GetUsername()
					capturedToken = in.GetAuthtoken()
					return tt.mockReply, tt.mockErr
				},
			}

			svc := NewDataService(mock)
			sections, err := svc.GetTypesOfData(context.Background(), tt.inputUser, tt.inputToken)

			if capturedUser != tt.inputUser {
				t.Errorf("forwarded username = %q, want %q", capturedUser, tt.inputUser)
			}
			if capturedToken != tt.inputToken {
				t.Errorf("forwarded authtoken = %q, want %q", capturedToken, tt.inputToken)
			}

			if tt.wantErr {
				if err == nil {
					t.Fatal("expected error, got nil")
				}
				if err.Error() != tt.wantErrMsg {
					t.Errorf("error = %q, want %q", err.Error(), tt.wantErrMsg)
				}
			} else {
				if err != nil {
					t.Errorf("unexpected error: %v", err)
				}
				if len(sections) != tt.wantSectionLen {
					t.Errorf("sections len = %d, want %d", len(sections), tt.wantSectionLen)
				}
			}
		})
	}
}

func TestDataService_GetAllData(t *testing.T) {
	t.Run("success single chunk", func(t *testing.T) {
		section := &pb.DataInformation{Datasection: "intData", Datalength: 3, Minchunk: 3, Maxchunk: 3}
		mock := &mockDataBackendClient{
			getDataFn: func(ctx context.Context, in *pb.DataRequest, opts ...grpc.CallOption) (*pb.DataReply, error) {
				return &pb.DataReply{Data: []*pb.DataItem{{IntVal: 1}, {IntVal: 2}, {IntVal: 3}}}, nil
			},
		}

		svc := NewDataService(mock)
		result, err := svc.GetAllData(context.Background(), "alice", "tok", []*pb.DataInformation{section})

		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if got := len(result["intData"]); got != 3 {
			t.Errorf("items = %d, want 3", got)
		}
	})

	t.Run("success multiple chunks", func(t *testing.T) {
		section := &pb.DataInformation{Datasection: "intData", Datalength: 6, Minchunk: 3, Maxchunk: 3}
		callCount := 0
		mock := &mockDataBackendClient{
			getDataFn: func(ctx context.Context, in *pb.DataRequest, opts ...grpc.CallOption) (*pb.DataReply, error) {
				callCount++
				return &pb.DataReply{Data: []*pb.DataItem{{IntVal: 1}, {IntVal: 2}, {IntVal: 3}}}, nil
			},
		}

		svc := NewDataService(mock)
		result, err := svc.GetAllData(context.Background(), "alice", "tok", []*pb.DataInformation{section})

		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if callCount != 2 {
			t.Errorf("GetData calls = %d, want 2", callCount)
		}
		if got := len(result["intData"]); got != 6 {
			t.Errorf("items = %d, want 6", got)
		}
	})

	t.Run("stringsData cap", func(t *testing.T) {
		section := &pb.DataInformation{Datasection: "stringsData", Datalength: 1000, Minchunk: 700, Maxchunk: 700}
		var capturedChunkSize int32
		mock := &mockDataBackendClient{
			getDataFn: func(ctx context.Context, in *pb.DataRequest, opts ...grpc.CallOption) (*pb.DataReply, error) {
				capturedChunkSize = in.GetDatachunksize()
				items := make([]*pb.DataItem, 700)
				return &pb.DataReply{Data: items}, nil
			},
		}

		svc := NewDataService(mock)
		result, err := svc.GetAllData(context.Background(), "alice", "tok", []*pb.DataInformation{section})

		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if capturedChunkSize != 700 {
			t.Errorf("requested chunk size = %d, want 700", capturedChunkSize)
		}
		if got := len(result["stringsData"]); got != 700 {
			t.Errorf("items = %d, want 700", got)
		}
	})

	t.Run("GetData rpc error", func(t *testing.T) {
		section := &pb.DataInformation{Datasection: "testSection", Datalength: 5, Minchunk: 2, Maxchunk: 2}
		mock := &mockDataBackendClient{
			getDataFn: func(ctx context.Context, in *pb.DataRequest, opts ...grpc.CallOption) (*pb.DataReply, error) {
				return nil, errors.New("rpc failed")
			},
		}

		svc := NewDataService(mock)
		_, err := svc.GetAllData(context.Background(), "alice", "tok", []*pb.DataInformation{section})

		if err == nil {
			t.Fatal("expected error, got nil")
		}
		wantMsg := fmt.Sprintf("section testSection stopped at index 0: rpc failed")
		if err.Error() != wantMsg {
			t.Errorf("error = %q, want %q", err.Error(), wantMsg)
		}
	})

	t.Run("empty data breaks loop", func(t *testing.T) {
		section := &pb.DataInformation{Datasection: "intData", Datalength: 10, Minchunk: 5, Maxchunk: 5}
		mock := &mockDataBackendClient{
			getDataFn: func(ctx context.Context, in *pb.DataRequest, opts ...grpc.CallOption) (*pb.DataReply, error) {
				return &pb.DataReply{Data: []*pb.DataItem{}}, nil
			},
		}

		svc := NewDataService(mock)
		result, err := svc.GetAllData(context.Background(), "alice", "tok", []*pb.DataInformation{section})

		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if got := len(result["intData"]); got != 0 {
			t.Errorf("items = %d, want 0", got)
		}
	})
}
