package main

import (
	"github.com/mayerkv/go-notifications/domain"
	"github.com/mayerkv/go-notifications/grpc-service"
	"github.com/mayerkv/go-notifications/repository"
	"google.golang.org/grpc"
	"log"
	"net"
)

func main() {
	templateRepository := repository.NewInMemoryTemplateRepository()
	notificationRepository := repository.NewInMemoryNotificationRepository()
	templateService := domain.NewTemplateService(templateRepository)
	notificationService := domain.NewNotificationService(notificationRepository, templateService)

	server := grpc_service.NewNotificationsServiceImpl(templateService, notificationService)

	if err := runGrpcSever(server); err != nil {
		log.Fatal(err)
	}
}

func runGrpcSever(server grpc_service.NotificationsServiceServer) error {
	listener, err := net.Listen("tcp", ":9090")
	if err != nil {
		return err
	}

	var opts []grpc.ServerOption
	grpcServer := grpc.NewServer(opts...)
	grpc_service.RegisterNotificationsServiceServer(grpcServer, server)

	return grpcServer.Serve(listener)
}
