package exponential_filter

import (

)

func ExponentialFilter(inputSignal []float64, alpha float64) []float64 {
	outputSignal := make([]float64, len(inputSignal))
	outputSignal[0] = inputSignal[0]

	for i := 1; i < len(inputSignal); i += 1 {
		outputSignal[i] = alpha*inputSignal[i] + (1-alpha)*outputSignal[i-1]
	}

	return outputSignal
}

