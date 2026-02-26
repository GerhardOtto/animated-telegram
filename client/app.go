package main

import (
	"context"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	pb "github.com/gerhardotto/animated-telegram/client/backendservice"
	"github.com/gerhardotto/animated-telegram/client/internal/config"
	"github.com/gerhardotto/animated-telegram/client/internal/service"
)

// App wires together configuration and the gRPC connection.
type App struct {
	conn     *grpc.ClientConn
	cfg      *config.Config
	greeting *service.GreetingService
	auth     *service.AuthService
}

// NewApp parses configuration and dials the gRPC server.
func NewApp() (*App, error) {
	cfg := config.Parse()

	conn, err := grpc.NewClient(
		cfg.Addr,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithUnaryInterceptor(timingInterceptor),
	)
	if err != nil {
		return nil, err
	}

	client := pb.NewDataBackendClient(conn)

	return &App{
		conn:     conn,
		cfg:      cfg,
		greeting: service.NewGreetingService(client),
		auth:     service.NewAuthService(client),
	}, nil
}

// Close releases the underlying gRPC connection.
func (a *App) Close() {
	a.conn.Close()
}

// Run executes the client workflow.
func (a *App) Run() error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := a.greeting.SayHello(ctx, a.cfg.Name); err != nil {
		return err
	}

	_, err := a.auth.GetToken(ctx, a.cfg.Name)
	return err
}
