package router

import (
	"github.com/labstack/echo/v4"

	"github.com/saeedjhn/go-backend-clean-arch/internal/bootstrap"
	"github.com/saeedjhn/go-backend-clean-arch/internal/delivery/http/router/healthcheckrouter"
	"github.com/saeedjhn/go-backend-clean-arch/internal/delivery/http/router/taskrouter"
	"github.com/saeedjhn/go-backend-clean-arch/internal/delivery/http/router/userrouter"
)

func Setup(
	app *bootstrap.Application,
	echo *echo.Echo,
) {
	routerGroup := echo.Group("")

	userrouter.New(app, routerGroup)
	taskrouter.New(app, routerGroup)
	healthcheckrouter.New(app, routerGroup)
}