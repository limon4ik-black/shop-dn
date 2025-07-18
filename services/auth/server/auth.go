package server

import (
	"context"
	authv1 "shop-dn/services/auth/gen/go"

	"google.golang.org/grpc"
)

type serverApi struct {
	authv1.UnimplementedAuthServiceServer
}

func Register(gRPC *grpc.Server) {
	authv1.RegisterAuthServiceServer(gRPC, &serverApi{})
}

func (s *serverApi) Register(ctx context.Context, req *authv1.RegisterRequest) (*authv1.RegisterResponse, error) {
	return &authv1.RegisterResponse{
		UserId: "user-123", // временно фиксированный ID
	}, nil
}

func (s *serverApi) Login(ctx context.Context, req *authv1.LoginRequest) (*authv1.LoginResponse, error) {
	panic("implement me")
}
