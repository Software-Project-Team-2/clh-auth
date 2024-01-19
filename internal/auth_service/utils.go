package auth_service

import (
	"context"
	"encoding/json"
	"errors"
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
