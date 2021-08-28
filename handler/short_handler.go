package handler

import (
	"azurl/repository"

	"github.com/labstack/echo/v4"
	"go.uber.org/zap"
)

type ShortHandler struct {
	ShortRepo repository.ShortRepo
	Logger    *zap.Logger
}

func (su *ShortHandler) Encode(c echo.Context) error {
	return nil
}

func (su *ShortHandler) Decode(c echo.Context) error {
	return nil
}

func (su *ShortHandler) Redirect(c echo.Context) error {
	return nil
}
