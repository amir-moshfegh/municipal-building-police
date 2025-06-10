package router

import (
	"github.com/amir-moshfegh/municipal-building-police/internal/adapter/handler"
	"github.com/labstack/echo/v4"
)

func Health(e *echo.Group) {
	h := handler.NewHealthHandler()

	e.GET("/", h.Health)
}
