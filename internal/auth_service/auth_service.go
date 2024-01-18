package auth_service

import (
	"context"
	clh_auth "github.com/Software-Project-Team-2/clh-auth/internal/pb/auth"
	"log"
)

type AuthService struct {
	clh_auth.UnimplementedAuthServiceServer
}

// TODO: implement ../jwt/jwt_token.go

func (s *AuthService) Login(ctx context.Context, req *clh_auth.LoginRequest) (*clh_auth.LoginResponse, error) {
	log.Printf("Login request: %v", req)
	return &clh_auth.LoginResponse{Token: "sample_token"}, nil
}

func (s *AuthService) ValidateToken(ctx context.Context, req *clh_auth.ValidateRequest) (*clh_auth.ValidateResponse, error) {
	log.Printf("ValidateToken request: %v", req)
	return &clh_auth.ValidateResponse{Valid: true}, nil
}
