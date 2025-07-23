package main

import (
	"net"
	"shop-dn/services/auth/internal/config"
	"shop-dn/services/auth/internal/db"
	"shop-dn/services/auth/internal/logger"
	"shop-dn/services/auth/server"
	"strconv"

	"google.golang.org/grpc"
)

func main() {

	cfg := config.LoadCfg()

	log := logger.InitLoggerSlogger(cfg.Env)

	// вот эту хуйню снизу переделать
	connDB, ctx := db.InitConnDb(log)
	defer connDB.Close(ctx)

	//TODO: инициализировать обьект приложения app??? думаю не надо посмотрим

	lis, err := net.Listen("tcp", ":"+strconv.Itoa(cfg.GRPC.Port))
	if err != nil {
		log.Error("failed to listen", "error", err)
	}

	grpcServer := grpc.NewServer()
	server.Register(grpcServer)

	log.Info("Auth gRPC server is running on port" + strconv.Itoa(cfg.GRPC.Port) + "...")
	if err := grpcServer.Serve(lis); err != nil {
		log.Error("failed to serve", "error", err)
	}
}
