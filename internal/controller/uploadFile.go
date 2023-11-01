package controller

import (
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/labstack/echo/v4"
)

func (b *BackendHandler) RouteUploadFile(c echo.Context) error {
	if c.Request().Method == "GET" {
		return c.Render(200, "uploadFiles.html", nil)
	} else {
		return c.HTML(http.StatusBadRequest, "<h1>Request is not accepted</h1>")
	}
}

func (b *BackendHandler) UploadFile(c echo.Context) error {
	name := c.FormValue("name")
	email := c.FormValue("email")

	file, err := c.FormFile("image")
	if err != nil {
		return err
	}

	src, err := file.Open()
	if err != nil {
		return err
	}
	defer src.Close()

	// Destination
	dst, err := os.Create(fmt.Sprintf("assets/uploads/%s", file.Filename))
	if err != nil {
		return err
	}
	defer dst.Close()

	if _, err = io.Copy(dst, src); err != nil {
		return err
	}

	return c.HTML(http.StatusOK, fmt.Sprintf("<p>File %s uploaded successfully with fields name=%s and email=%s.</p>", file.Filename, name, email))
}
