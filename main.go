package main

import (
	"awesomeProject/kmeans"
	"fmt"
)

func main() {
	dataset := []float64{0.01, 0.02, 0.23, 0.34, 0.98, 0.99}
	k, cluster := kmeans.BestKmeans(dataset, 1, 5)
	fmt.Printf("the best cluster number: %v\n", k)
	fmt.Printf("the best cluster : %v\n", cluster)
}
