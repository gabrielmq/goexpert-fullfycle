package main

import (
	"database/sql"
	"log"
	"net"

	"github.com/gabrielmq/grpc/internal/database"
	"github.com/gabrielmq/grpc/internal/pb"
	"github.com/gabrielmq/grpc/internal/service"
	_ "github.com/mattn/go-sqlite3"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	db, err := sql.Open("sqlite3", "./data.db")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	categoryDB := database.NewCategory(db)
	categoryService := service.NewCategoryService(*categoryDB)

	// criando o servidor gRPC
	grpcServer := grpc.NewServer()
	reflection.Register(grpcServer)

	// registrando o serviço no servidor
	pb.RegisterCategoryServiceServer(grpcServer, categoryService)

	// abrindo uma conexão tpc para se comunicar com o servidor gRPC
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		panic(err)
	}

	// iniciando o servidor gRPC
	log.Println("starting server on port 50051")
	if err := grpcServer.Serve(lis); err != nil {
		panic(err)
	}
}
