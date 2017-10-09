package main

import (
	"github.com/labstack/echo"
	"github.com/ricardomarquesramos/charttaker/api"
)

func main() {
	e := echo.New()
	api.Router(e)

	e.Logger.Fatal(e.Start(":1323"))
}
