package main

import (
	"io/ioutil"
	"net"
	"os"

	"google.golang.org/grpc"
	"google.golang.org/grpc/grpclog"
	userspb "github.com/raymond-design/buffy/proto/users/v1;users"
)

func main() {
	log := grpclog.NewLoggerV2(os.Stdout, ioutil.Discard, ioutil.Discard)
	grpclog.SetLoggerV2(log)

	listener, err := net.Listen("tcp", "0.0.0.0:10000")
	if err != nil {
		log.Fatal(err)
	}

	server := grpc.NewServer()
	userspb.RegisterUserServiceServer(server, nil)
	reflection.Register(server)

	log.Infoln("Server listening on:", listender.Addr().String()
	err = server.serve(listener)
	if err != nil {
		log.Fatal(err)
	}
}
