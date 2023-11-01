package controller

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

func (b *BackendHandler) HomePage(c echo.Context) error {

	return c.Redirect(http.StatusTemporaryRedirect, fmt.Sprintf("%s/dashboard", b.cfg.BaseURL))
}
