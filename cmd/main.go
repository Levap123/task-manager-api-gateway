package main

import (
	"os"
	"os/signal"
	"syscall"

	"github.com/Levap123/task-manager-api-gateway/config"
	"github.com/Levap123/task-manager-api-gateway/internal/client/rpc"
	"github.com/Levap123/task-manager-api-gateway/internal/transport/rest"
	utils "github.com/Levap123/task-manager-api-gateway/pkg"
	"github.com/Levap123/task-manager-api-gateway/proto"
	"github.com/Levap123/task-manager-api-gateway/servers"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/grpclog"
)

// @title task-manager API documentation
// @version 1.0.0// @host localhost:8080
// @BasePath /
func main() {
	cfg := config.NewConfigs()

	logger := utils.NewLogger()

	server := new(servers.Server)

	opts := []grpc.DialOption{
		grpc.WithInsecure(),
	}
	connAuth, err := grpc.Dial(cfg.AUTH_ADDRESS, opts...)
	if err != nil {
		grpclog.Fatalf("failt to dial: %v\n", err)
	}

	defer connAuth.Close()

	clientAuth := proto.NewAuthClient(connAuth)

	client := rpc.NewClient(clientAuth)

	rest := rest.NewRest(logger, client)

	go func() {
		if err := server.Run(cfg.SERVER_ADDRESS, rest.InitRoutes()); err != nil {
			logrus.Fatalln(err)
		}
	}()
	logrus.Print("taks manager api gateway started")
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<-quit

	logrus.Print("taks manager api gateway shutting down")
}
