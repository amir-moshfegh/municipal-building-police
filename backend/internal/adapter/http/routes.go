package http

import (
	"fmt"
	"github.com/amir-moshfegh/municipal-building-police/config"
	"github.com/amir-moshfegh/municipal-building-police/docs"
	"github.com/amir-moshfegh/municipal-building-police/internal/adapter/middlewares"
	"github.com/amir-moshfegh/municipal-building-police/internal/adapter/router"
	"github.com/labstack/echo/v4"
	echoSwagger "github.com/swaggo/echo-swagger"
)

func InitRoutes(e *echo.Echo, cfg *config.Config) {

	api := e.Group("/api", middlewares.RedirectTrailingSlash())
	v1 := api.Group("/v1")

	health := v1.Group("/health")

	v1.GET("/test", func(c echo.Context) error {
		return c.JSON(200, map[string]string{"message": "Test endpoint is working"})
	})

	router.Health(health)
}

func InitSwaggerRoutes(e *echo.Echo, cfg config.ServerConfig) {
	docs.SwaggerInfo.Title = "Municipal Building Police API"
	docs.SwaggerInfo.Description = "API documentation for the Municipal Building Police"
	docs.SwaggerInfo.Version = "1.0.0"
	docs.SwaggerInfo.Host = fmt.Sprintf("%s%s", cfg.Host, cfg.Port)
	docs.SwaggerInfo.BasePath = "/api"
	docs.SwaggerInfo.Schemes = []string{"http", "https"}

	e.GET("/swagger/*", echoSwagger.WrapHandler)
}
