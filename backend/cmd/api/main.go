package main

import (
	"fmt"
	"net"

	"backend/cmd/api/pb"
	"backend/cmd/api/services"
	"backend/cmd/model"
	"backend/pkg/db"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	session, err := db.ConnectDatabase()
	if err != nil {
		panic(err)
	}

	model.SetRepo(session)

	defer model.CloseDB()
	var opts []grpc.ServerOption
	grpcServer := grpc.NewServer(opts...)
	chatServer := services.NewChatServer()

	pb.RegisterChatServiceServer(grpcServer,chatServer)
	fmt.Println("--- SERVER APP ---")

	address := fmt.Sprintf("0.0.0.0:%d", 5400)
	listener, err := net.Listen("tcp", address)
	if err != nil {
		panic(err)
	}
	reflection.Register(grpcServer)
	grpcServer.Serve(listener)

}
