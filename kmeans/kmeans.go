// Package kmeans The kmeans algorithm for 1 Dimension
package kmeans

import (
	"fmt"
	"github.com/montanaflynn/stats"
	"math"
)

func createClusters(dataset []float64, centers []float64) ([][]float64, error) {
	clusters := make([][]float64, len(centers))
	for _, data := range dataset {
		centerIndex := 0
		minDis := math.MaxFloat64
		for i, center := range centers{
			dis := (data - center) * (data - center)
			fmt.Printf("the difference between centers: %v\n", dis)
			if minDis > dis {
				minDis = dis
				centerIndex = i
			}
		}
		fmt.Printf("the centerIndex: %v\n", centerIndex)
		clusters[centerIndex] = append(clusters[centerIndex], data)
	}
	return clusters, nil
}

func updateCenters(centers []float64, clusters [][]float64) []float64{
	var newCenters []float64
	for _, clu := range clusters {
		mean, _ := stats.Mean(clu)
		newCenters = append(newCenters, mean)
	}
	return newCenters
}

// CalculateCenter calculate the center of each cluster in k-means++ algorithm
//	The main improvement of kmeans++ algorithm: enhance the stability of the clusters' number by determining the centers stably
func initCenter(dataset []float64, k int) ([]float64, error){
	centers := make([]float64, k)
	n := len(dataset)
	distance := make([]float64, n)
	centerIndex := 0
	centers[0] = dataset[centerIndex]
	for i, _ := range centers {
		for j, data := range dataset {
			distance[j] = nearest(data, centers[0:i])
		}
		max := distance[0]
		maxIndex := 0
		for i2, dis := range distance {
			if max < dis {
				max = dis
				maxIndex = i2
			}
		}
		centers[i] = dataset[maxIndex]
	}
	return centers,nil
}

func nearest(data float64, centers []float64) float64{
	minDis := math.MaxFloat64
	for _, center := range centers {
		dis := (data - center) * (data - center)
		if minDis > dis {
			minDis = dis
		}
	}
	return minDis
}

func Kmeans(dataset []float64, k int) ([]float64, [][]float64, error) {
	var clusters [][]float64
	var centers []float64
	var newCenters []float64
	maxIterations := 100
	varepsilon := 0.001
	flag := true
	//step1: calculate the centers of each cluster
	centers, err := initCenter(dataset, k)
	fmt.Printf("the initial centers: %v\n", centers)
	if err != nil {
		fmt.Printf("fail to calculate the centers: %v\n", err)
	}
	clusters, _ = createClusters(dataset, centers)
	fmt.Printf("the initial clusters: %v\n", clusters)
	for j := 0; j < maxIterations && flag; j++ {
		clusters, _ = createClusters(dataset, centers)
		newCenters = updateCenters(centers, clusters)
		var i int
		for i = 0; i < k; i++ {
			if (newCenters[i] - centers[i]) > varepsilon {
				break
			}
		}
		if i == k {
			flag = false
		}
		centers = newCenters
	}
	return centers, clusters, nil
}

// BestKmeans calculate the best k during [k1, k2]
func BestKmeans(dataset []float64, k1 int, k2 int) (int, [][]float64){
	num := len(dataset)
	fmt.Printf("database: %v\n", dataset)
	var sse []float64
	for k := k1; k <= k2; k++ {
		//step3: evaluate the cluster results
		fmt.Printf("the k = or cluster number: %v\n", k)
		centers, clusters, err := Kmeans(dataset, k)
		fmt.Printf("centers: %v\n", centers)
		fmt.Printf("clusters: %v\n", clusters)
		if err != nil {
			fmt.Printf("fail to run Kmeans!")
		}
		w := float64(0)
		for i, clu := range clusters {
			aver := centers[i]
			n := 0
			d := float64(0)
			for j, c := range clu {
				d += (c - aver) * (c - aver)
				n = j + 1
				fmt.Printf("d of cluster %v is : %v\n", n, d)
			}
			w = w + d / float64(2 * n)
			fmt.Printf("Centered at x: %.2f\n", centers[i])
			fmt.Printf("Matching data points: %+v\n\n", clu)
		}
		sse = append(sse, w)
	}
	fmt.Printf("sse : %v\n", sse)
	gap := make([]float64, num)
	gapSquare := make([]float64, num)
	maxElbow := float64(0)
	perfectCluster := 0
	for i := 0; i < len(sse) - 1; i++ {
		gap[i] = math.Abs(sse[i + 1] - sse[i])
	}
	for i := 0; i < len(gap) - 1; i++ {
		gapSquare[i] = math.Abs(gap[i + 1] - gap[i])
		if maxElbow < gapSquare[i] {
			maxElbow = gapSquare[i]
			perfectCluster = i + 2
		}
	}
	fmt.Printf("perfect cluster number is : %v\n", perfectCluster)

	_, clusters, _ := Kmeans(dataset,perfectCluster)
	fmt.Printf("perfect cluster is : %v\n", clusters)
	return perfectCluster, clusters
}