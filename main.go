package main

import (
	"io/ioutil"
	"net"
	"os"

	userspb "github.com/raymond-design/buffy/proto/users/v1"
	"github.com/raymond-design/buffy/users"
	"google.golang.org/grpc"
	"google.golang.org/grpc/grpclog"
	"google.golang.org/grpc/reflection"
)

func main() {
	log := grpclog.NewLoggerV2(os.Stdout, ioutil.Discard, ioutil.Discard)
	grpclog.SetLoggerV2(log)

	listener, err := net.Listen("tcp", "0.0.0.0:10000")
	if err != nil {
		log.Fatal(err)
	}

	server := grpc.NewServer()
	userspb.RegisterUserServiceServer(server, &users.Server{})
	reflection.Register(server)

	log.Infoln("Server listening on:", listener.Addr().String())

	err = server.Serve(listener)
	if err != nil {
		log.Fatal(err)
	}
}
