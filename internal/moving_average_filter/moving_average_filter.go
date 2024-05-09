package moving_average_filter

import (

)

func MovingAverageFilter(inputSignal []float64, windowSize int) []float64 {
	outputSignal := make([]float64, len(inputSignal))

	for i := 0; i < len(inputSignal); i += 1 {
		sum := 0.0
		counter := 0

		for j := i - windowSize + 1; j <= i; j += 1 {
			if j >= 0 && j < len(inputSignal) {
				sum += inputSignal[j]
				counter += 1
			}
		}

		if counter > 0 {
			outputSignal[i] = sum/float64(counter)
		} else {
			outputSignal[i] = 0
		}
	}
	return outputSignal
}