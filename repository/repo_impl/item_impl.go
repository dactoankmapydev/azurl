package repo_impl

import (
	"izurl/db"
	"izurl/model"
	"izurl/repository"
	"time"
)

type RedisRepoImpl struct {
	client *db.Redis
}

func NewRedisRepo(client *db.Redis) repository.ItemRepo {
	return &RedisRepoImpl{
		client: client,
	}
}

func (rp *RedisRepoImpl) Save(string, time.Time) (string, error) {
	return "", nil
}

func (rp *RedisRepoImpl) Load(string) (string, error) {
	return "", nil
}

func (rp *RedisRepoImpl) LoadInfo(string) (*model.Item, error) {
	return nil, nil
}

func (rp *RedisRepoImpl) Close() error {
	return nil
}
