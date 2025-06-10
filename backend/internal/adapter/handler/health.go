package handler

import (
	"github.com/amir-moshfegh/municipal-building-police/internal/adapter/helper"
	"github.com/labstack/echo/v4"
	"net/http"
)

type HealthHandler struct{}

func NewHealthHandler() *HealthHandler {
	return &HealthHandler{}
}

// HealthCheck godoc
// @Summary Health Check
// @Description Health Check
// @Tags health
// @Accept  json
// @Produce  json
// @Success 200 {object} helper.BaseHttpResponse "Success"
// @Failure 400 {object} helper.BaseHttpResponse "Failed"
// @Router /v1/health/ [get]
func (h *HealthHandler) Health(c echo.Context) error {
	return c.JSON(http.StatusOK, helper.GenerateBaseResponse(true, "Working!", map[string]string{"status": "ok"}))
}
