package main

import (
	"fmt"
	"image/color"
	"log"
	"math"

	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/vg"

	ef "github.com/Leonid-Sarmatov/golang-signal-filters/internal/exponential_filter"
	rnd "github.com/Leonid-Sarmatov/golang-signal-filters/internal/random_utils"
)

func main() {
	//in := []float64{1, 2, 3, 4, 5, 6, 7, 23, 40, 30, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20}
	in := make([]float64, 1000)
	for i := 0; i < len(in); i += 1 {
		in[i] = math.Sin(float64(i)*0.01)
	}

	for i := 0; i < len(in)-100; i += 1 {
		r := rnd.RandomNumber(0, 30) 
		if r < 2 {
			r := float64(rnd.RandomNumber(-1000, 1000))/5000
			//in[i-3] += 0.1 * r
			//in[i-2] += 0.2 * r
			in[i-1] += 0.5 * r
			in[i+0] += 2.0 * r
			in[i+1] += 0.5 * r
			//in[i+2] += 0.2 * r
			//in[i+3] += 0.1 * r
			/*for j := i; j < i + 5; j += 1 {
				in[j] += (float64(j-i)/20)*r
			}
			for j := i + 5; j >= 0; j -= 1 {
				in[j] += (float64(j-i)/20)*r
			}*/
		}
	}

	out := ef.ExponentialFilter(in, 0.5)

	inputPoints := make(plotter.XYs, len(in))
	for i := range in {
		inputPoints[i].X = float64(i)
		inputPoints[i].Y = in[i]
	}

	outputPoints := make(plotter.XYs, len(in))
	for i := range in {
		outputPoints[i].X = float64(i)
		outputPoints[i].Y = out[i]
	}

	inputLine, err := plotter.NewLine(inputPoints)
	if err != nil {
		log.Fatal(err)
	}
	inputLine.Color = color.RGBA{R: 255, A: 255}

	outputLine, err := plotter.NewLine(outputPoints)
	if err != nil {
		log.Fatal(err)
	}
	outputLine.Color = color.RGBA{B: 255, A: 255}

	p := plot.New()
	p.Add(inputLine, outputLine)
	p.Title.Text = "Экспоненциальный фильтр\nКрасный сигнал - входной сигнал, синий - выходной"
	p.X.Label.Text = "Время"
	p.Y.Label.Text = "Значение сигнала"

	if err := p.Save(10*vg.Inch, 6*vg.Inch, "signals.png"); err != nil {
		log.Fatal(err)
	}

	fmt.Println("OK")
}