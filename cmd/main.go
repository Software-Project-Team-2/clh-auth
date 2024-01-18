package main

import (
	"github.com/Software-Project-Team-2/clh-auth/internal/pb/auth"
	"google.golang.org/grpc"
	"log"
	"net"
)

type server struct {
	clh_auth.UnimplementedInventoryServer
}

func main() {
	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		panic(err)
	}

	s := grpc.NewServer()
	clh_auth.RegisterInventoryServer(s, &server{})
	if err := s.Serve(listener); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
