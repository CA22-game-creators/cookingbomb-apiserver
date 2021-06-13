package main

import (
	"fmt"
	"log"
	"net"
	"os"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	pb "github.com/CA22-game-creators/cookingbomb-proto/server/pb/api"

	accountDI "github.com/CA22-game-creators/cookingbomb-apiserver/di/account"
)

func main() {
	listen, err := net.Listen("tcp", fmt.Sprintf(":%s", os.Getenv("PORT")))
	if err != nil {
		log.Fatalf(err.Error())
	}
	log.Printf("Listening :%s", os.Getenv("PORT"))

	grpcServer := grpc.NewServer()
	pb.RegisterAccountServicesServer(grpcServer, accountDI.DI())

	reflection.Register(grpcServer)
	if err = grpcServer.Serve(listen); err != nil {
		log.Fatalf(err.Error())
	}
}
