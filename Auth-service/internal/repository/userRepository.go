package repository

import (
	"context"
	"fmt"

	"github.com/DestWish/HackMate/Auth-service/internal/models"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

type UserRepository struct {
	db	*gorm.DB
	redisClient *redis.Client
}


func NewUserRepo(db *gorm.DB, redisClient *redis.Client) *UserRepository {
	return &UserRepository{db: db, redisClient:  redisClient}
}


func (r *UserRepository) userCaching(ctx context.Context, user *models.User) error {
	key := userCacheKey(User.login)
}


func userCacheKey(userID uint) string {
	return fmt.Sprintf("user:%v", userID)
}