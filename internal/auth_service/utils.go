package auth_service

import (
	"context"
	"encoding/json"
	"errors"
	"net/mail"
	"strconv"

	"github.com/Software-Project-Team-2/clh-auth/internal/entities"
	"github.com/Software-Project-Team-2/clh-auth/internal/redis_client"
)

func CreateUserHashRedis(id int64, user entities.User) error {
	var ctx = context.Background()
	rdb := redis_client.GetClient()

	userData, err := json.Marshal(user)
	if err != nil {
		return err
	}

	return rdb.HSet(ctx, "user_profile:"+strconv.FormatInt(id, 10), "data", userData).Err()
}

func LinkUserEmailWithId(email string, userId int64) error {
	rdb := redis_client.GetClient()
	ctx := context.Background()

	userIdStr := strconv.FormatInt(userId, 10)

	err := rdb.Set(ctx, "user_profile:"+email, userIdStr, 0).Err()
	return err
}

func GetUserIdByEmail(email string) (int64, error) {
	rdb := redis_client.GetClient()
	ctx := context.Background()

	result, err := rdb.Get(ctx, "user_profile:"+email).Result()
	if err != nil {
		return 0, err
	}

	userID, err := strconv.ParseInt(result, 10, 64)
	if err != nil {
		return 0, err
	}

	return userID, nil
}

func GetUserHashRedis(id int) (*entities.User, error) {
	var ctx = context.Background()

	rdb := redis_client.GetClient()
	result, err := rdb.HGetAll(ctx, "user_profile:"+strconv.Itoa(id)).Result()

	if err != nil {
		return nil, err
	}

	userData, exists := result["data"]
	if !exists {
		return nil, errors.New("user not found")
	}

	var user entities.User
	err = json.Unmarshal([]byte(userData), &user)

	if err != nil {
		return nil, err
	}

	return &user, nil
}

func GetUserProfileByEmail(email string) (*entities.User, error) {
	userId, err := GetUserIdByEmail(email)
	if err != nil {
		return nil, err
	}

	userProfile, err := GetUserHashRedis(int(userId))
	if err != nil {
		return nil, err
	}

	return userProfile, nil
}

func validEmail(email string) bool {
	_, err := mail.ParseAddress(email)
	return err == nil
}
