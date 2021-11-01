package db

import (
	"context"

	"github.com/go-redis/redis/v8"
	"go.uber.org/zap"
)

type Redis struct {
	Client *redis.Client
	Logger *zap.Logger

	// Dùng dưới local
	Host string
	Port string

	// Dùng trên server heroku
	// Url string
}

func (rd *Redis) NewRedis() {
	// Dùng dưới local
	rd.Client = redis.NewClient(&redis.Options{
		Addr: rd.Host + ":" + rd.Port,
	})

	// Dùng trên server heroku
	// opt, _ := redis.ParseURL(rd.Url)
	// rd.Client = redis.NewClient(opt)

	_, err := rd.Client.Ping(context.Background()).Result()
	if err != nil {
		rd.Logger.Error("Kết nối không thành công tới redis ", zap.Error(err))
	}

	rd.Logger.Info("Kết nối thành công tới redis")
}
