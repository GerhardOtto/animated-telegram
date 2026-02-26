package main

import (
	"context"
	"log/slog"
	"time"

	"google.golang.org/grpc"
)

func timingInterceptor(
	ctx context.Context,
	method string,
	req, reply any,
	cc *grpc.ClientConn,
	invoker grpc.UnaryInvoker,
	opts ...grpc.CallOption,
) error {
	start := time.Now()
	err := invoker(ctx, method, req, reply, cc, opts...)
	duration := time.Since(start)
	if err != nil {
		slog.ErrorContext(ctx, "rpc failed", "method", method, "duration", duration, "error", err)
		return err
	}
	slog.InfoContext(ctx, "rpc completed", "method", method, "duration", duration)
	return nil
}
