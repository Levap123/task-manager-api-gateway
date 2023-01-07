package main

import (
	"os"
	"os/signal"
	"syscall"

	utils "github.com/Levap123/task-manager-api-gateway/pkg"
	"github.com/Levap123/task-manager-api-gateway/servers"
	"github.com/Levap123/task-manager-api-gateway/transport/rest"
	"github.com/sirupsen/logrus"
)

// @title task-manager API documentation
// @version 1.0.0// @host localhost:8080
// @BasePath /
func main() {
	logger := utils.NewLogger()
	rest := rest.NewRest(logger)
	server := new(servers.Server)
	go func() {
		if err := server.Run(":8080", rest.InitRoutes()); err != nil {
			logrus.Fatalln(err)
		}
	}()
	logrus.Print("taks manager api gateway started")
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<-quit

	logrus.Print("taks manager api gateway shutting down")
}
