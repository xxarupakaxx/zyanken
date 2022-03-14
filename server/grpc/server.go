package main

import (
	"fmt"
	grpc_logrus "github.com/grpc-ecosystem/go-grpc-middleware/logging/logrus"
	"github.com/sirupsen/logrus"
	pb "github.com/xxarupakaxx/zyanken/gen/proto"
	"github.com/xxarupakaxx/zyanken/server/router"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
	"os"
	"os/signal"
)

func main() {
	port := 50051
	lis,err := net.Listen("tcp",fmt.Sprintf(":%d",port))
	if err != nil {
		log.Fatalf("failed to listen :%v",err)
	}

	logrusLogger := logrus.New()
	logrusEnty := logrus.NewEntry(logrusLogger)
	grpc_logrus.ReplaceGrpcLogger(logrusEnty)
	server := grpc.NewServer(grpc.UnaryInterceptor(grpc_logrus.UnaryServerInterceptor(logrusEnty)))

	pb.RegisterZyankenServiceServer(server,router.NewGameHandler())
	pb.RegisterMatchingServiceServer(server,router.NewMatchingHandler())

	reflection.Register(server)

	go func() {
		log.Println("start gRPC server port ", port)
		err = server.Serve(lis)
		if err != nil {
			log.Println("failed to start server :",err)
		}
	}()
	quit := make(chan os.Signal)
	signal.Notify(quit,os.Interrupt)
	<- quit
	log.Println("stopping gRPC server..")
	server.GracefulStop()
}