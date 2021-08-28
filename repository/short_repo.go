package repository

import (
	"azurl/model"

	"time"
)

type ShortRepo interface {
	Save(string, time.Time) (string, error)
	Load(string) (string, error)
	LoadInfo(string) (*model.Url, error)
	Close() error
}
