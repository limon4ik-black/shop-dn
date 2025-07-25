package server

import (
	"context"
	authv1 "shop-dn/services/auth/gen/go"

	"github.com/jackc/pgx/v5/pgxpool"
	"google.golang.org/grpc"
)

type serverApi struct {
	authv1.UnimplementedAuthServiceServer
	conn *pgxpool.Pool
	ctx  context.Context
}

func Register(gRPC *grpc.Server, conn *pgxpool.Pool, ctx context.Context) {
	authv1.RegisterAuthServiceServer(gRPC,
		&serverApi{
			conn: conn,
			ctx:  ctx,
		})
}

func (s *serverApi) Register(ctx context.Context, req *authv1.RegisterRequest) (*authv1.RegisterResponse, error) {
	// TODO: берем req.Name and req.Password
	var id string
	err := s.conn.QueryRow(s.ctx, "SELECT insert_user($1, $2)", req.Name, req.Password).Scan(&id)
	if err != nil {
		// либо лог сделать либо ошибка grpc/status
	}
	// INSERT INTO users (id name pass){
	//		}
	// return select id from users where name = req.Name
	//
	return &authv1.RegisterResponse{
		UserId: id, // временно фиксированный ID
	}, nil
}

func (s *serverApi) Login(ctx context.Context, req *authv1.LoginRequest) (*authv1.LoginResponse, error) {
	panic("implement me")
}
