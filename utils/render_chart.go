package utils

import (
	"bytes"
	"fmt"

	"github.com/wcharczuk/go-chart"
	"github.com/wcharczuk/go-chart/drawing"
)

// Generates a bar chart and returns the PNG image bytes
func GenerateBarChart(values []int64) ([]byte, error) {
	// times corresponding to time_periods
	labels := []string{
		"3:00 - 05:59",
		"6:00 - 06:59",
		"7:00 - 08:59",
		"9:00 - 13:29",
		"13:30 - 15:59",
		"16:00 - 18:29",
		"18:30 - 21:59",
		"22:00 - 23:59",
		"0:00 - 02:59",
	}

	if len(labels) != len(values) {
		fmt.Println(len(values))
		return nil, fmt.Errorf("number of labels and values must be the same")
	}

	graph := chart.BarChart{
		Title:      "Ridership Fall 2017",
		TitleStyle: chart.StyleShow(),
		Background: chart.Style{
			Padding: chart.Box{
				Top: 50,
			},
		},
		Height:   400,
		BarWidth: 40,
		XAxis:    chart.StyleShow(),
		YAxis: chart.YAxis{
			Style: chart.StyleShow(),
			Range: &chart.ContinuousRange{
				Min: 0,
				Max: 5500000,
			},
			Ticks: []chart.Tick{
				// Customize y-axis ticks to display as integers
				{Value: 0, Label: "0"},
				{Value: 1000000, Label: "1,000,000"},
				{Value: 2000000, Label: "2,000,000"},
				{Value: 3000000, Label: "3,000,000"},
				{Value: 4000000, Label: "4,000,000"},
				{Value: 5000000, Label: "5,000,000"},
			},
		},
		Bars: []chart.Value{},
	}

	for i, label := range labels {
		graph.Bars = append(graph.Bars, chart.Value{
			Label: label,
			Value: float64(values[i]),
			Style: chart.Style{
				Show:        true,
				FillColor:   drawing.ColorBlue,
				StrokeColor: drawing.ColorBlue,
			},
		})
	}

	// Render the chart as PNG image bytes
	buffer := bytes.NewBuffer([]byte{})
	err := graph.Render(chart.PNG, buffer)
	if err != nil {
		return nil, err
	}

	return buffer.Bytes(), nil
}
