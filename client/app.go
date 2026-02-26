package main

import (
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	"github.com/gerhardotto/animated-telegram/client/internal/config"
)

// App wires together configuration and the gRPC connection.
type App struct {
	conn *grpc.ClientConn
	cfg  *config.Config
}

// NewApp parses configuration and dials the gRPC server.
func NewApp() (*App, error) {
	cfg := config.Parse()

	conn, err := grpc.NewClient(cfg.Addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}

	return &App{
		conn: conn,
		cfg:  cfg,
	}, nil
}

// Close releases the underlying gRPC connection.
func (a *App) Close() {
	a.conn.Close()
}

// Run executes the client workflow.
func (a *App) Run() error {
	return nil
}
