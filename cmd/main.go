package main

import (
	"os"
	"os/signal"
	"syscall"

	"github.com/Levap123/task-manager-api-gateway/config"
	_ "github.com/Levap123/task-manager-api-gateway/docs"
	"github.com/Levap123/task-manager-api-gateway/internal/client/rpc"
	"github.com/Levap123/task-manager-api-gateway/internal/transport/rest"
	utils "github.com/Levap123/task-manager-api-gateway/pkg"
	"github.com/Levap123/task-manager-api-gateway/proto"
	"github.com/Levap123/task-manager-api-gateway/servers"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/grpclog"
)

// @title           API gateway for microservice application
// @version         1.0
// @description     API gateway for microservice task-manager app, GRPC, protobuf
// @termsOfService  http://swagger.io/terms/

// @contact.name   Levap Mik
// @contact.url    https://t.me/kavelpim123
// @contact.email  pavlikkim200401@gmail.com

// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html

// @host      localhost:8080
// @BasePath  /
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
