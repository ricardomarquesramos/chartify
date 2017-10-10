package main

import (
	"github.com/labstack/echo"
	"github.com/ricardomarquesramos/chartify/api"
)

func main() {
	e := echo.New()
	api.Router(e)

	e.Logger.Fatal(e.Start(":80"))
}
