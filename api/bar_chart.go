package api

import (
	"encoding/base64"
	"encoding/json"
	"log"
	"net/http"
	"strings"

	"github.com/ricardomarquesramos/charttaker/charts"

	"github.com/labstack/echo"
	//exposes "chart"
)

// BarChartParams parameters to render a bar chart
type getBarChartInput struct {
	Series    []charts.Series `json:"series"`
	Statistic string          `json:"statistic"`
}

// GetBarChart gets a bar chart with the given params
func GetBarChart(c echo.Context) error {
	params := c.Param("params")

	if strings.TrimSpace(params) == "" {
		return c.String(http.StatusOK, "")
	}

	data, err := base64.StdEncoding.DecodeString(params)

	if err != nil {
		log.Println("Could not decode input params:", err)
		return c.String(http.StatusUnprocessableEntity, "Cound not decode params")
	}

	var input getBarChartInput
	err = json.Unmarshal(data, &input)

	if err != nil {
		log.Println("Malformed input params:", err)
		return c.String(http.StatusUnprocessableEntity, "Malformed input params")
	}

	barChartBytes, err := charts.RenderBarChartAsBytes(input.Series, "percentage")

	if err != nil {
		log.Println("Could not generate the chart:", err)
		return c.String(http.StatusUnprocessableEntity, "Could not generate the chart")
	}

	return c.Blob(http.StatusOK, "image/png", barChartBytes)
}
