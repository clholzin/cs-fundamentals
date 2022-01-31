package repo

import "github.com/go-redis/redis/v8"

func NewRepo(url, password string, db int) *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr:     url,
		Password: password,
		DB:       db,
	})
}
