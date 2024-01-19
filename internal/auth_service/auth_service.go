package auth_service

import (
	"context"
	"log"

	"github.com/Software-Project-Team-2/clh-auth/internal/entities"
	clh_auth "github.com/Software-Project-Team-2/clh-auth/internal/pb/auth"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type AuthService struct {
	clh_auth.UnimplementedAuthServiceServer
}

func (s *AuthService) Login(ctx context.Context, req *clh_auth.LoginRequest) (*clh_auth.LoginResponse, error) {

	username, password := req.GetUsername(), req.GetPassword()

	if !validateUserCredentials(username, password) {
		return nil, status.Error(codes.Unauthenticated, "invalid username or password")
	}

	return &clh_auth.LoginResponse{Token: "sample_token"}, nil
}

func (s *AuthService) CreateUser(ctx context.Context, req *clh_auth.CreateUserRequest) (*clh_auth.CreateUserResponse, error) {
	if req.GetUsername() == "" || req.GetPassword() == "" {
		return nil, status.Errorf(codes.InvalidArgument, "username and password cannot be empty")
	}

	if len(req.GetUsername()) < 3 || len(req.GetPassword()) < 6 {
		return nil, status.Errorf(codes.InvalidArgument, "username must be at least 3 characters and password at least 6 characters long")
	}

	user := entities.User{Password: req.GetPassword(), Name: req.GetUsername(), Email: "test@example.com"}
	userId := GenerateUserId()

	CreateUserHashRedis(userId, user)

	return &clh_auth.CreateUserResponse{Success: true, Message: "User has been created!"}, nil
}

func (s *AuthService) ValidateToken(ctx context.Context, req *clh_auth.ValidateRequest) (*clh_auth.ValidateResponse, error) {
	log.Printf("ValidateToken request: %v", req)
	return &clh_auth.ValidateResponse{Valid: true}, nil
}

func validateUserCredentials(username, password string) bool {
	// Implement actual credential validation logic here
	return username == "expectedUsername" && password == "expectedPassword"
}
