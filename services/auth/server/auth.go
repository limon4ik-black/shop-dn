package server

import (
	"context"
	authv1 "shop-dn/services/auth/gen/go"
	"shop-dn/services/auth/internal/validate"

	"github.com/jackc/pgx/v5/pgxpool"
	"golang.org/x/crypto/bcrypt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
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

	err := validate.NameIsValid(req.Name) // ъуйня какаято надо переделать
	if err != nil {
		return nil, err
	}

	err = validate.PasswordIsValid(req.Password) // ъуйня какаято надо переделать
	if err != nil {
		return nil, err
	}

	var hashPass []byte

	hashPass, err = bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)

	if err != nil {
		return nil, status.Error(codes.Internal, "failed to hash password")
	}

	var id string
	err = s.conn.QueryRow(s.ctx, "SELECT insert_user($1, $2)", req.Name, string(hashPass)).Scan(&id)
	if err != nil {
		return nil, status.Error(codes.Internal, "failed to select intp db")
	}

	return &authv1.RegisterResponse{
		UserId: id,
	}, nil
}

func (s *serverApi) Login(ctx context.Context, req *authv1.LoginRequest) (*authv1.LoginResponse, error) {
	//TODO: compare name // это select nedded write
	//TODO: compare password func CompareHashAndPassword(hashedPassword, password []byte) error
	//TODO: сделать базу с jwt token
	//TODO: выдавать jwt token

	return &authv1.LoginResponse{
		Token: "token",
	}, nil
}
