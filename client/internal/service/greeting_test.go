package service

import (
	"context"
	"errors"
	"testing"

	pb "github.com/gerhardotto/animated-telegram/client/backendservice"
	"google.golang.org/grpc"
)

func TestGreetingService_SayHello(t *testing.T) {
	tests := []struct {
		name       string
		inputName  string
		mockReply  *pb.HelloReply
		mockErr    error
		wantErr    bool
		wantErrMsg string
	}{
		{
			name:      "success",
			inputName: "Alice",
			mockReply: &pb.HelloReply{Message: "Hello, Alice"},
			mockErr:   nil,
			wantErr:   false,
		},
		{
			name:       "rpc error",
			inputName:  "Bob",
			mockReply:  nil,
			mockErr:    errors.New("rpc failed"),
			wantErr:    true,
			wantErrMsg: "rpc failed",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var capturedName string
			mock := &mockDataBackendClient{
				sayHelloFn: func(ctx context.Context, in *pb.HelloRequest, opts ...grpc.CallOption) (*pb.HelloReply, error) {
					capturedName = in.GetName()
					return tt.mockReply, tt.mockErr
				},
			}

			svc := NewGreetingService(mock)
			err := svc.SayHello(context.Background(), tt.inputName)

			if capturedName != tt.inputName {
				t.Errorf("forwarded name = %q, want %q", capturedName, tt.inputName)
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
			}
		})
	}
}
