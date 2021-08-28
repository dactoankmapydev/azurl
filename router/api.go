package router

import (
	"izurl/handler"

	"github.com/labstack/echo/v4"
)

type API struct {
	Echo        *echo.Echo
	ItemHandler handler.ItemHandler
}

func (api *API) SetupRouter() {

	// item
	user := api.Echo.Group("/")
	user.POST("encode", api.ItemHandler.Encode)
	user.GET("{shortLink}", api.ItemHandler.Decode)
	user.GET("{shortLink}/info", api.ItemHandler.Redirect)
}
