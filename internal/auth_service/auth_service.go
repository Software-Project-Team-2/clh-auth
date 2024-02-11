package auth_service

import (
	"context"
	"fmt"

	"github.com/Software-Project-Team-2/clh-auth/internal/entities"
	"github.com/Software-Project-Team-2/clh-auth/internal/jwt"
	clh_auth "github.com/Software-Project-Team-2/clh-auth/internal/pb/auth"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type AuthService struct {
	clh_auth.UnimplementedAuthServiceServer
}

func (s *AuthService) Login(ctx context.Context, req *clh_auth.LoginRequest) (*clh_auth.LoginResponse, error) {

	email, password := req.GetEmail(), req.GetPassword()

	if !validateUserCredentials(email, password) {
		return nil, status.Error(codes.Unauthenticated, "invalid username or password")
	}

	i, err := GetUserIdByEmail(email)

	if err != nil {
		fmt.Printf("Unable to parse it : %v", err)
	}

	token, err2 := jwt.GenerateJWT(i, email)

	if err2 != nil {
		fmt.Printf("Unable to generate JWT token: %v", err2)
	}

	return &clh_auth.LoginResponse{Token: token}, nil
}

func (s *AuthService) CreateUser(ctx context.Context, req *clh_auth.CreateUserRequest) (*clh_auth.CreateUserResponse, error) {
	if req.GetUsername() == "" || req.GetPassword() == "" || req.GetEmail() == "" {
		return nil, status.Errorf(codes.InvalidArgument, "username and password and email cannot be empty")
	}

	if len(req.GetUsername()) < 3 || len(req.GetPassword()) < 6 {
		return nil, status.Errorf(codes.InvalidArgument, "username must be at least 3 characters and password at least 6 characters long")
	}

	if validEmail(req.GetEmail()) == false {
		return nil, status.Errorf(codes.Canceled, "email does not fit the regex")
	}

	_, err := GetUserIdByEmail(req.GetEmail())
	if err == nil {
		return nil, status.Errorf(codes.AlreadyExists, "user with this email already exists")
	}

	user := entities.User{Password: req.GetPassword(), Name: req.GetUsername(), Email: "test@example.com"}
	userId := GenerateUserId()

	CreateUserHashRedis(userId, user)
	LinkUserEmailWithId(req.Email, userId)

	return &clh_auth.CreateUserResponse{Success: true, Message: "User has been created!"}, nil
}

func (s *AuthService) ValidateToken(ctx context.Context, req *clh_auth.ValidateRequest) (*clh_auth.ValidateResponse, error) {

	isValid := jwt.ValidateToken(req.GetToken())

	return &clh_auth.ValidateResponse{Valid: isValid}, nil
}

func validateUserCredentials(email, password string) bool {
	u, err := GetUserProfileByEmail(email)

	if err != nil {
		fmt.Printf("%v", err)
		return false
	}

	return u.Password == password
}
