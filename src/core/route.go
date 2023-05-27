package core

import (
	"github.com/labstack/echo/v4"
	"gitlab.com/putuyaza/antrian/src/config"
	"gitlab.com/putuyaza/antrian/src/routes"
)

func Route() {
	route := echo.New()
	// load route for web
	routes.Web(route)
	// load app config
	app := config.App()
	// start server
	route.Logger.Fatal(route.Start(":" + app.PORT))
}
