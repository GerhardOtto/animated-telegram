package service

import (
	"context"

	pb "github.com/gerhardotto/animated-telegram/client/backendservice"
	"google.golang.org/grpc"
)

type mockDataBackendClient struct {
	sayHelloFn     func(ctx context.Context, in *pb.HelloRequest, opts ...grpc.CallOption) (*pb.HelloReply, error)
	getAuthTokenFn func(ctx context.Context, in *pb.AuthTokenRequest, opts ...grpc.CallOption) (*pb.AuthTokenReply, error)
	getTypesOfData func(ctx context.Context, in *pb.DataInfoRequest, opts ...grpc.CallOption) (*pb.DataInfoReply, error)
	getDataFn      func(ctx context.Context, in *pb.DataRequest, opts ...grpc.CallOption) (*pb.DataReply, error)
}

func (m *mockDataBackendClient) SayHello(ctx context.Context, in *pb.HelloRequest, opts ...grpc.CallOption) (*pb.HelloReply, error) {
	if m.sayHelloFn != nil {
		return m.sayHelloFn(ctx, in, opts...)
	}
	return nil, nil
}

func (m *mockDataBackendClient) GetAuthToken(ctx context.Context, in *pb.AuthTokenRequest, opts ...grpc.CallOption) (*pb.AuthTokenReply, error) {
	if m.getAuthTokenFn != nil {
		return m.getAuthTokenFn(ctx, in, opts...)
	}
	return nil, nil
}

func (m *mockDataBackendClient) GetTypesOfData(ctx context.Context, in *pb.DataInfoRequest, opts ...grpc.CallOption) (*pb.DataInfoReply, error) {
	if m.getTypesOfData != nil {
		return m.getTypesOfData(ctx, in, opts...)
	}
	return nil, nil
}

func (m *mockDataBackendClient) GetData(ctx context.Context, in *pb.DataRequest, opts ...grpc.CallOption) (*pb.DataReply, error) {
	if m.getDataFn != nil {
		return m.getDataFn(ctx, in, opts...)
	}
	return nil, nil
}
