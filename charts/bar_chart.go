package charts

import (
	"bytes"

	chart "github.com/wcharczuk/go-chart"
)

func RenderBarChartAsBytes(series []Series, statistic string) ([]byte, error) {
	chartSeries := buildChartSeries(series)

	barChart := buildBarChart(chartSeries)

	buffer := bytes.NewBuffer([]byte{})
	err := barChart.Render(chart.PNG, buffer)

	if err != nil {
		return nil, err
	}

	return buffer.Bytes(), nil
}

func buildChartSeries(series []Series) *[]chart.Value {
	chartSeries := make([]chart.Value, len(series))

	for i, bar := range series {
		chartSeries[i] = chart.Value{
			Label: bar.Label,
			Value: bar.Value,
		}
	}

	return &chartSeries
}

func buildBarChart(chartSeries *[]chart.Value) *chart.BarChart {
	barChart := chart.BarChart{
		Height:   512,
		BarWidth: 60,
		XAxis: chart.Style{
			Show: true,
		},
		YAxis: chart.YAxis{
			Style: chart.Style{
				Show: true,
			},
		},
		Bars: *chartSeries,
	}

	return &barChart
}
