package main

import (
	"fmt"
	"math/cmplx"
	"os"

	"github.com/wcharczuk/go-chart"
)

const ITERATION_NUMBER = 50

func mandelbrot(point complex128) bool {
	n := 1
	var z complex128 = 0
	for cmplx.Abs(z) < 50 && n <= ITERATION_NUMBER {
		z = z*z + point
		n = n + 1
	}

	if n == ITERATION_NUMBER+1 {
		return true
	}
	return false
}

func main() {
	var XValues, YValues []float64
	for j := -100.00; j < 101.00; j += 0.1 {
		fmt.Println(j)
		for k := -100.00; k < 101.00; k += 0.1 {
			isMandelbrot := mandelbrot(complex(j, k))
			if isMandelbrot == true {
				XValues = append(XValues, j)
				YValues = append(YValues, k)
			}
		}
	}

	ticks := []chart.Tick{
		{Value: -3.0, Label: "-3.0"},
		{Value: -2.0, Label: "-2.0"},
		{Value: -1.0, Label: "-1.0"},
		{Value: 0.0, Label: "0.0"},
		{Value: 1.0, Label: "1.0"},
		{Value: 2.0, Label: "2.0"},
	}

	graph := chart.Chart{
		XAxis: chart.XAxis{
			Ticks: ticks,
			Style: chart.StyleShow(),
		},
		YAxis: chart.YAxis{
			Ticks: ticks,
			Style: chart.StyleShow(),
		},
		Series: []chart.Series{
			chart.ContinuousSeries{
				Style: chart.Style{
					StrokeWidth: chart.Disabled,
					DotWidth:    3,
					Show:        true,
				},
				XValues: XValues,
				YValues: YValues,
			},
		},
	}
	pngFile, _ := os.Create("output.png")
	chartError := graph.Render(chart.PNG, pngFile)
	fmt.Println(chartError)
	fmt.Println(-100.00 + 0.1)
}
