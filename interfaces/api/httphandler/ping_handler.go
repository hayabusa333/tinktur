package httphandler

import (
	"net/http"

	"github.com/labstack/echo"
)

func Ping(c echo.Context) error {
	return c.String(http.StatusOK, "pong")
}
