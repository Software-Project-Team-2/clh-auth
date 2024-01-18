package main

import (
	"github.com/Software-Project-Team-2/clh-auth/internal/auth_service"
	clh_auth "github.com/Software-Project-Team-2/clh-auth/internal/pb/auth"
	"google.golang.org/grpc"
	"log"
	"net"
	"os"
)

func main() {
	_, jwtTokenIsPresent := os.LookupEnv("JWT_SECRET_TOKEN")
	_, isProduction := os.LookupEnv("IS_PRODUCTION")

	if jwtTokenIsPresent == false && isProduction == true {
		log.Panic("Please set up a env variable for JWT token: \"JWT_SECRET_TOKEN\"")
	}

	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()

	clh_auth.RegisterAuthServiceServer(grpcServer, &auth_service.AuthService{})

	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
