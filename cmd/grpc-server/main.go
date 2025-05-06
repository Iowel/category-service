package main

import (
	"log"
	"net"

	categoryservice "github.com/Iowel/category-service/internal/app/category-service"
	"github.com/Iowel/category-service/internal/service/category"
	cat_repository "github.com/Iowel/category-service/internal/service/category/repository"
	category_service "github.com/Iowel/category-service/pkg/api/category-service"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	grpcServer := grpc.NewServer()

	reflection.Register(grpcServer)

	categoryRepo := cat_repository.New()
	categoryService := category.New(categoryRepo)

	// создаём gRPC-имплементацию с нужным сервисом
	categoryServer := categoryservice.NewCategoryServiceServer(categoryService)

	// регстрация сервиса
	category_service.RegisterCategoryServiceServer(grpcServer, categoryServer)

	// запуск gRPC сервера
	listener, err := net.Listen("tcp", ":44044")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	log.Println("gRPC server started on :44044")
	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}

}

// type GrpcServer struct {
// 	categoryService category.Service
// }

// func NewGrpcServer(categoryService category.Service) *GrpcServer {
// 	return &GrpcServer{
// 		categoryService: categoryService,
// 	}
// }

// func (s *GrpcServer) Start (/* config */) error {
// 	ctx, cancel := context.WithCancel(context.Background())
// 	defer cancel()

// }
