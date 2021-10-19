package main

import (
	"awesomeProject/kmeans"
	"awesomeProject/zscore"
	"fmt"
	"github.com/montanaflynn/stats"
)

func main() {
	dataset := []float64 {
		1.52, 1.35, 1.52, 30.1, 30.1, 30, 22.9, 23.1, 22.9,160,160,160,30.4,0.024,27.3,143,23.4,
		27.6,0.21,138,163,142,138,0.02,30.5,0.047,27.4,142,23.5,27.3,0.43,138,163,142,138,0.55,
		30.4,0.066,27.3,142,23.4,27.3,0.42,138,163,142,138,0.67,
	}
	//dataset2 := []float64 {
	//	30.1, 30.1, 30, 22.9, 23.1, 22.9,160,160,160,27.3,143,
	//	138,27.4,142,138,27.3,142,138,
	//}
	data, _ := zscore.ZScore(dataset)
	_, cluster := kmeans.BestKmeans(data, 1, 5)
	for i, clu := range cluster {
		center, _ := stats.Mean(clu)
		fmt.Printf("the cluster %v :\n %v\n", i + 1, clu)
		fmt.Printf("the cluster components number: %v\n", len(clu))
		fmt.Printf("the cluster center : %v\n", center)
	}
}