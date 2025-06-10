package middlewares

import (
	"github.com/amir-moshfegh/municipal-building-police/config"
	"github.com/labstack/echo/v4"
	"net/http"
	"strings"
)

func Cors(cfg config.CorsConfig) func(next echo.HandlerFunc) echo.HandlerFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			c.Response().Header().Set("Access-Control-Allow-Origin", cfg.AllowOrigin)
			c.Response().Header().Set("Access-Control-Allow-Methods", cfg.AllowMethods)
			c.Response().Header().Set("Access-Control-Allow-Headers", cfg.AllowHeaders)
			c.Response().Header().Set("Access-Control-Allow-Credentials", "true")

			origin := c.Request().Header.Get(echo.HeaderOrigin)
			if c.Request().Method == http.MethodOptions && strings.HasPrefix(origin, "http") {
				return c.NoContent(http.StatusNoContent)
			}

			return next(c)
		}
	}
}
