package controller

import (
	"github.com/labstack/echo/v4"
)

func (b *BackendHandler) HealthcheckHandler(c echo.Context) error {
	return c.Render(200, "hello.html", nil)
}
