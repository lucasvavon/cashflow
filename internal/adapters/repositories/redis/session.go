package redis

import (
	"context"
	"fmt"
	"github.com/redis/go-redis/v9"
	"strconv"
	"time"
)

type RedisSessionRepository struct {
	client *redis.Client
}

func NewRedisSessionRepository() *RedisSessionRepository {
	client := redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
	})
	return &RedisSessionRepository{client: client}
}

func (r *RedisSessionRepository) SaveSession(ctx context.Context, sessionID string, userID uint, ttl time.Duration) error {
	return r.client.Set(ctx, sessionID, userID, ttl).Err()
}

func (r *RedisSessionRepository) GetSession(ctx context.Context, sessionID string) (uint, error) {

	result, err := r.client.Get(ctx, sessionID).Result()
	if err != nil {
		return 0, err
	}

	sessionIDUint, err := strconv.ParseUint(result, 10, 32)
	if err != nil {
		return 0, fmt.Errorf("invalid session ID format: %v", err)
	}

	return uint(sessionIDUint), nil
}

func (r *RedisSessionRepository) DeleteSession(ctx context.Context, sessionID string) error {
	return r.client.Del(ctx, sessionID).Err()
}
