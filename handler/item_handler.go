package handler

import (
	"izurl/repository"

	"github.com/labstack/echo/v4"
	"go.uber.org/zap"
)

type ItemHandler struct {
	ItemRepo repository.ItemRepo
	Logger   *zap.Logger
}

func (item *ItemHandler) Encode(c echo.Context) error {
	return nil
}

func (item *ItemHandler) Decode(c echo.Context) error {
	return nil
}

func (item *ItemHandler) Redirect(c echo.Context) error {
	return nil
}
