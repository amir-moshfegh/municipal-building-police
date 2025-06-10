package http

import (
	"github.com/amir-moshfegh/municipal-building-police/config"
	"github.com/amir-moshfegh/municipal-building-police/internal/adapter/middlewares"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func NewServer(cfg *config.Config) *echo.Echo {
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middlewares.Cors(cfg.Cors))

	return e
}
