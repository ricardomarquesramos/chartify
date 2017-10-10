package main

import (
	"os"

	"github.com/labstack/echo"
	"github.com/ricardomarquesramos/chartify/api"
)

func main() {
	e := echo.New()
	api.Router(e)

	e.File("/", "public/demo.html")

	port := ":" + os.Getenv("PORT")
	e.Logger.Fatal(e.Start(port))
}
