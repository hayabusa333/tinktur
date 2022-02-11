package main

import (
	"github.com/labstack/echo"

	"tinktur/interfaces/api/httphandler"
)

func main() {
	e := echo.New()

	_ = e.GET("/ping", httphandler.Ping)

	e.Logger.Fatal(e.Start(":1323"))
}
