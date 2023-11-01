package controller

import (
	"fmt"

	"github.com/labstack/echo/v4"
)

func (b *BackendHandler) GetFile(c echo.Context) error {
	fileName := c.Param("file_name")
	return c.File(fmt.Sprintf("assets/uploads/%s", fileName))
}
