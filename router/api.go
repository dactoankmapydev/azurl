package router

import (
	"azurl/handler"

	"github.com/labstack/echo/v4"
)

type API struct {
	Echo         *echo.Echo
	ShortHandler handler.ShortHandler
}

func (api *API) SetupRouter() {

	// item
	user := api.Echo.Group("/")
	user.POST("encode", api.ShortHandler.Encode)
	user.GET("{shortUrl}", api.ShortHandler.Decode)
	user.GET("{shortUrl}/info", api.ShortHandler.Redirect)
}
