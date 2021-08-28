package repository

import (
	"izurl/model"
	"time"
)

type ItemRepo interface {
	Save(string, time.Time) (string, error)
	Load(string) (string, error)
	LoadInfo(string) (*model.Item, error)
	Close() error
}
