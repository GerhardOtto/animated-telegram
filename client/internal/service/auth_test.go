package service

import (
	"context"
	"errors"
	"testing"

	pb "github.com/gerhardotto/animated-telegram/client/backendservice"
	"google.golang.org/grpc"
)

func TestAuthService_GetToken(t *testing.T) {
	tests := []struct {
		name        string
		inputUser   string
		mockReply   *pb.AuthTokenReply
		mockErr     error
		wantErr     bool
		wantErrMsg  string
		wantToken   string
	}{
		{
			name:      "success",
			inputUser: "alice",
			mockReply: &pb.AuthTokenReply{Authtoken: "tok-abc123"},
			mockErr:   nil,
			wantErr:   false,
			wantToken: "tok-abc123",
		},
		{
			name:       "rpc error",
			inputUser:  "bob",
			mockReply:  nil,
			mockErr:    errors.New("rpc failed"),
			wantErr:    true,
			wantErrMsg: "rpc failed",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var capturedUsername string
			mock := &mockDataBackendClient{
				getAuthTokenFn: func(ctx context.Context, in *pb.AuthTokenRequest, opts ...grpc.CallOption) (*pb.AuthTokenReply, error) {
					capturedUsername = in.GetUsername()
					return tt.mockReply, tt.mockErr
				},
			}

			svc := NewAuthService(mock)
			token, err := svc.GetToken(context.Background(), tt.inputUser)

			if capturedUsername != tt.inputUser {
				t.Errorf("forwarded username = %q, want %q", capturedUsername, tt.inputUser)
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
				if token != tt.wantToken {
					t.Errorf("token = %q, want %q", token, tt.wantToken)
				}
			}
		})
	}
}
