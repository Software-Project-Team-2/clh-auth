package auth_service

import (
	"context"
	"fmt"
	"github.com/Software-Project-Team-2/clh-auth/internal/entities"
	"github.com/Software-Project-Team-2/clh-auth/internal/jwt"
	clh_auth "github.com/Software-Project-Team-2/clh-auth/internal/pb/auth"
	"golang.org/x/crypto/bcrypt"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"os"
)

type AuthService struct {
	clh_auth.UnimplementedAuthServiceServer
}

func (s *AuthService) Login(ctx context.Context, req *clh_auth.LoginRequest) (*clh_auth.LoginResponse, error) {
	email, password := req.GetEmail(), req.GetPassword()

	credentials, err := validateUserCredentials(email, password)
	if !credentials {
		return nil, status.Error(codes.Unauthenticated, "invalid username or password")
	}

	i, err := GetUserIdByEmail(email)
	if err != nil {
		return nil, status.Error(codes.Aborted, "Not able to get userid from email")
	}

	token, err := jwt.GenerateJWT(i, email)

	if err != nil {
		return nil, status.Error(codes.Aborted, "Unable to generate JWT token")
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

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.GetPassword()), bcrypt.DefaultCost)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "could not hash password")
	}

	user := entities.User{Password: string(hashedPassword), Name: req.GetUsername(), Email: "test@example.com", Permission: 0}
	userId := GenerateUserId()

	err = CreateUserHashRedis(userId, user)
	if err != nil {
		return nil, err

	}

	err = LinkUserEmailWithId(req.Email, userId)
	if err != nil {
		return nil, err
	}

	return &clh_auth.CreateUserResponse{Success: true, Message: "User has been created!"}, nil
}

func (s *AuthService) ValidateToken(_ context.Context, req *clh_auth.ValidateRequest) (*clh_auth.ValidateResponse, error) {
	if req.GetToken() == os.Getenv("ADMIN_TOKEN") {
		return &clh_auth.ValidateResponse{
			Valid:       true,
			Permissions: &clh_auth.UserPermissionsResponse{Permissions: 2}}, nil
	}

	claims, ok := jwt.ParseUserFromToken(req.GetToken())
	if !ok {
		return nil, fmt.Errorf("invalid or expired token")
	}

	userID, ok := (*claims)["id"].(float64)
	if !ok {
		return nil, fmt.Errorf("user ID not found in token")
	}

	userIDInt := int64(userID)
	userProfile, err := GetUserHashRedis(int(userIDInt))
	if err != nil {
		return nil, fmt.Errorf("failed to fetch user profile: %v", err)
	}

	permissions := clh_auth.UserPermissionsResponse{Permissions: int32(userProfile.Permission)}
	isValid := jwt.ValidateToken(req.GetToken())

	return &clh_auth.ValidateResponse{Valid: isValid, Permissions: &permissions}, nil
}

func (s AuthService) GetUserPermissions(ctx context.Context, req *clh_auth.UserPermissionsRequest) (*clh_auth.UserPermissionsResponse, error) {
	//check for Admin JWT token
	if req.GetToken() == os.Getenv("ADMIN_TOKEN") {
		return &clh_auth.UserPermissionsResponse{
			Permissions: int32(2),
		}, nil
	}

	// Parse the JWT token from the request
	claims, ok := jwt.ParseUserFromToken(req.GetToken())
	if !ok {
		return nil, fmt.Errorf("invalid or expired token")
	}

	userID, ok := (*claims)["id"].(float64)
	if !ok {
		return nil, fmt.Errorf("user ID not found in token")
	}

	userIDInt := int64(userID)
	userProfile, err := GetUserHashRedis(int(userIDInt))
	if err != nil {
		return nil, fmt.Errorf("failed to fetch user profile: %v", err)
	}

	permissions := userProfile.Permission

	response := &clh_auth.UserPermissionsResponse{
		Permissions: int32(permissions),
	}

	return response, nil
}

func validateUserCredentials(email, password string) (bool, error) {
	u, err := GetUserProfileByEmail(email)
	if err != nil {
		fmt.Printf("%v", err)
		return false, err
	}

	err = bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password))
	if err != nil {
		return false, nil
	}

	return true, nil
}
