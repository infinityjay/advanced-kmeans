package zscore

import (
	"github.com/montanaflynn/stats"
	"math"
)
func ZScore(dataset []float64) ([]float64, error){
	var z []float64
	mu, _ := stats.Mean(dataset)
	sum := float64(0)
	for _, data := range dataset {
		sum = sum + (data - mu)*(data - mu)
	}
	v := sum / float64(len(dataset) -1)
	sigma := math.Sqrt(v)
	for _, data := range dataset{
		z = append(z, math.Abs(data - mu) / sigma)
	}
	return z, nil
}