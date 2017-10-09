package api

import "github.com/labstack/echo"

// Router initialization
func Router(e *echo.Echo) {
	e.GET("/bar/:params", GetBarChart)
}
