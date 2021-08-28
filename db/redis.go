package db

import (
	"context"
	"time"

	"github.com/go-redis/redis/v8"
	"go.uber.org/zap"
)

type Redis struct {
	Client *redis.Client
	Host   string
	Port   string
	Logger *zap.Logger
}

func (rd *Redis) NewRedis() {
	rd.Client = redis.NewClient(&redis.Options{
		Addr:        rd.Host + ":" + rd.Port,
		IdleTimeout: 240 * time.Second,
	})
	_, err := rd.Client.Ping(context.Background()).Result()
	if err != nil {
		rd.Logger.Error("Kết nối không thành công tới redis ", zap.Error(err))
	}

	rd.Logger.Info("Kết nối thành công tới redis")
}
