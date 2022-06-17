/* Go package for Clustering Algorithm. */
package main

import (
	"fmt"
	"math"
	"math/rand"
	"time"
)

/*
 * K-Means Clustering Algorithm.
 *
 * @param dataSet: The data set to be clustered.
 * @param k: The number of clusters.
 * @param maxIterations: The maximum number of iterations.
 * @param distance: The distance function.
 * @param centroids: The initial centroids.
 * @return: The clusters.
 */
func KMeans(dataSet [][]float64, k int, maxIterations int, distance func([][]float64, [][]float64) float64, centroids [][]float64) [][][]float64 {
	// Initialize the centroids.
	if centroids == nil {
		centroids = make([][]float64, k)
		for i := 0; i < k; i++ {
			centroids[i] = make([]float64, len(dataSet[0]))
		}
		for i := 0; i < k; i++ {
			for j := 0; j < len(dataSet[0]); j++ {
				centroids[i][j] = dataSet[rand.Intn(len(dataSet))][j]
			}
		}
	}

	// Initialize the clusters.
	clusters := make([][][]float64, k)
	for i := 0; i < k; i++ {
		clusters[i] = make([][]float64, 0)
	}

	// Initialize the distance matrix.
	distanceMatrix := make([][]float64, len(dataSet))
	for i := 0; i < len(dataSet); i++ {
		distanceMatrix[i] = make([]float64, k)
	}

	// Initialize the iteration number.
	iteration := 0

	// Initialize the flag.
	flag := true

	// Initialize the random number generator.
	rand.Seed(time.Now().UnixNano())

	// Iterate until the maximum number of iterations is reached or the clusters do not change.
	for iteration < maxIterations && flag {
		// Update the distance matrix.
		for i := 0; i < len(dataSet); i++ {
			for j := 0; j < k; j++ {
				distanceMatrix[i][j] = distance(dataSet[i:i+1], centroids[j:j+1])
			}
		}

		// Update the clusters.
		clusters = make([][][]float64, k)
		for i := 0; i < k; i++ {
			clusters[i] = make([][]float64, 0)
		}
		for i := 0; i < len(dataSet); i++ {
			min := math.MaxFloat64
			index := 0
			for j := 0; j < k; j++ {
				if distanceMatrix[i][j] < min {
					min = distanceMatrix[i][j]
					index = j
				}
			}
			clusters[index] = append(clusters[index], dataSet[i])
		}

		// Update the centroids.
		for i := 0; i < k; i++ {
			for j := 0; j < len(dataSet[0]); j++ {
				sum := 0.0
				for l := 0; l < len(clusters[i]); l++ {
					sum += clusters[i][l][j]
				}
				centroids[i][j] = sum / float64(len(clusters[i]))
			}
		}

		// Update the iteration number.
		iteration++

		// Update the flag.
		flag = false
		for i := 0; i < k; i++ {
			if len(clusters[i]) == 0 {
				flag = true
				break
			}
		}
	}

	// Return the clusters.
	return clusters
}

/*
 * K-Means++ Clustering Algorithm.
 *
 * @param dataSet: The data set to be clustered.
 * @param k: The number of clusters.
 * @param maxIterations: The maximum number of iterations.
 * @param distance: The distance function.
 * @return: The clusters.
 */
func KMeansPlusPlus(dataSet [][]float64, k int, maxIterations int, distance func([][]float64, [][]float64) float64) [][][]float64 {
	// Initialize the centroids.
	centroids := make([][]float64, k)
	for i := 0; i < k; i++ {
		centroids[i] = make([]float64, len(dataSet[0]))
	}
	centroids[0] = dataSet[rand.Intn(len(dataSet))]
	for i := 1; i < k; i++ {
		// Initialize the distance matrix.
		distanceMatrix := make([][]float64, len(dataSet))
		for j := 0; j < len(dataSet); j++ {
			distanceMatrix[j] = make([]float64, i)
		}

		// Update the distance matrix.
		for j := 0; j < len(dataSet); j++ {
			for l := 0; l < i; l++ {
				distanceMatrix[j][l] = distance(dataSet[j:j+1], centroids[l:l+1])
			}
		}

		// Initialize the minimum distance.
		min := math.MaxFloat64

		// Initialize the index.
		index := 0

		// Find the minimum distance.
		for j := 0; j < len(dataSet); j++ {
			sum := 0.0
			for l := 0; l < i; l++ {
				sum += distanceMatrix[j][l]
			}
			if sum < min {
				min = sum
				index = j
			}
		}

		// Update the centroids.
		centroids[i] = dataSet[index]
	}

	// Return the clusters.
	return KMeans(dataSet, k, maxIterations, distance, centroids)
}

/*
 * K-Medoids Clustering Algorithm.
 *
 * @param dataSet: The data set to be clustered.
 * @param k: The number of clusters.
 * @param maxIterations: The maximum number of iterations.
 * @param distance: The distance function.
 * @param centroids: The initial centroids.
 * @return: The clusters.
 */
func KMedoids(dataSet [][]float64, k int, maxIterations int, distance func([][]float64, [][]float64) float64, centroids [][]float64) [][][]float64 {
	// Initialize the centroids.
	if centroids == nil {
		centroids = make([][]float64, k)
		for i := 0; i < k; i++ {
			centroids[i] = make([]float64, len(dataSet[0]))
		}
		for i := 0; i < k; i++ {
			for j := 0; j < len(dataSet[0]); j++ {
				centroids[i][j] = dataSet[rand.Intn(len(dataSet))][j]
			}
		}
	}

	// Initialize the clusters.
	clusters := make([][][]float64, k)
	for i := 0; i < k; i++ {
		clusters[i] = make([][]float64, 0)
	}

	// Initialize the distance matrix.
	distanceMatrix := make([][]float64, len(dataSet))
	for i := 0; i < len(dataSet); i++ {
		distanceMatrix[i] = make([]float64, k)
	}

	// Initialize the iteration number.
	iteration := 0

	// Initialize the flag.
	flag := true

	// Initialize the random number generator.
	rand.Seed(time.Now().UnixNano())

	// Iterate until the maximum number of iterations is reached or the clusters do not change.
	for iteration < maxIterations && flag {
		// Update the distance matrix.
		for i := 0; i < len(dataSet); i++ {
			for j := 0; j < k; j++ {
				distanceMatrix[i][j] = distance(dataSet[i:i+1], centroids[j:j+1])
			}
		}

		// Update the clusters.
		clusters = make([][][]float64, k)
		for i := 0; i < k; i++ {
			clusters[i] = make([][]float64, 0)
		}
		for i := 0; i < len(dataSet); i++ {
			min := math.MaxFloat64
			index := 0
			for j := 0; j < k; j++ {
				if distanceMatrix[i][j] < min {
					min = distanceMatrix[i][j]
					index = j
				}
			}
			clusters[index] = append(clusters[index], dataSet[i])
		}

		// Update the centroids.
		for i := 0; i < k; i++ {
			min := math.MaxFloat64
			index := 0
			for j := 0; j < len(clusters[i]); j++ {
				sum := 0.0
				for l := 0; l < len(clusters[i]); l++ {
					sum += distance(clusters[i][j:j+1], clusters[i][l:l+1])
				}
				if sum < min {
					min = sum
					index = j
				}
			}
			centroids[i] = clusters[i][index]
		}

		// Update the iteration number.
		iteration++

		// Update the flag.
		flag = false
		for i := 0; i < k; i++ {
			if len(clusters[i]) == 0 {
				flag = true
				break
			}
		}
	}

	// Return the clusters.
	return clusters
}

/*
 * K-Medoids++ Clustering Algorithm.
 *
 * @param dataSet: The data set to be clustered.
 * @param k: The number of clusters.
 * @param maxIterations: The maximum number of iterations.
 * @param distance: The distance function.
 * @return: The clusters.
 */
func KMedoidsPlusPlus(dataSet [][]float64, k int, maxIterations int, distance func([][]float64, [][]float64) float64) [][][]float64 {
	// Initialize the centroids.
	centroids := make([][]float64, k)
	for i := 0; i < k; i++ {
		centroids[i] = make([]float64, len(dataSet[0]))
	}
	centroids[0] = dataSet[rand.Intn(len(dataSet))]
	for i := 1; i < k; i++ {
		// Initialize the distance matrix.
		distanceMatrix := make([][]float64, len(dataSet))
		for j := 0; j < len(dataSet); j++ {
			distanceMatrix[j] = make([]float64, i)
		}

		// Update the distance matrix.
		for j := 0; j < len(dataSet); j++ {
			for l := 0; l < i; l++ {
				distanceMatrix[j][l] = distance(dataSet[j:j+1], centroids[l:l+1])
			}
		}

		// Initialize the minimum distance.
		min := math.MaxFloat64

		// Initialize the index.
		index := 0

		// Find the minimum distance.
		for j := 0; j < len(dataSet); j++ {
			sum := 0.0
			for l := 0; l < i; l++ {
				sum += distanceMatrix[j][l]
			}
			if sum < min {
				min = sum
				index = j
			}
		}

		// Update the centroids.
		centroids[i] = dataSet[index]
	}

	// Return the clusters.
	return KMedoids(dataSet, k, maxIterations, distance, centroids)
}

/*
 * Hierarchical Clustering Algorithm.
 *
 * @param dataSet: The data set to be clustered.
 * @param k: The number of clusters.
 * @param distance: The distance function.
 * @return: The clusters.
 */
func Hierarchical(dataSet [][]float64, k int, distance func([][]float64, [][]float64) float64) [][][]float64 {
	// Initialize the clusters.
	clusters := make([][][]float64, len(dataSet))
	for i := 0; i < len(dataSet); i++ {
		clusters[i] = make([][]float64, 1)
		clusters[i][0] = dataSet[i]
	}

	// Initialize the distance matrix.
	distanceMatrix := make([][]float64, len(dataSet))
	for i := 0; i < len(dataSet); i++ {
		distanceMatrix[i] = make([]float64, len(dataSet))
	}

	// Update the distance matrix.
	for i := 0; i < len(dataSet); i++ {
		for j := 0; j < len(dataSet); j++ {
			distanceMatrix[i][j] = distance(dataSet[i:i+1], dataSet[j:j+1])
		}
	}

	// Iterate until the number of clusters is reached.
	for len(clusters) > k {
		// Initialize the minimum distance.
		min := math.MaxFloat64

		// Initialize the indices.
		index1 := 0
		index2 := 0

		// Find the minimum distance.
		for i := 0; i < len(dataSet); i++ {
			for j := 0; j < len(dataSet); j++ {
				if distanceMatrix[i][j] < min {
					min = distanceMatrix[i][j]
					index1 = i
					index2 = j
				}
			}
		}

		// Merge the clusters.
		clusters[index1] = append(clusters[index1], clusters[index2]...)
		clusters = append(clusters[:index2], clusters[index2+1:]...)

		// Update the distance matrix.
		for i := 0; i < len(dataSet); i++ {
			distanceMatrix[index1][i] = distance(clusters[index1], clusters[i])
			distanceMatrix[i][index1] = distanceMatrix[index1][i]
		}
	}

	// Return the clusters.
	return clusters
}

/*
 * DBSCAN Clustering Algorithm.
 *
 * @param dataSet: The data set to be clustered.
 * @param epsilon: The maximum distance between two points.
 * @param minPts: The minimum number of points.
 * @param distance: The distance function.
 * @return: The clusters.
 */
func DBSCAN(dataSet [][]float64, epsilon float64, minPts int, distance func([][]float64, [][]float64) float64) [][][]float64 {
	// Initialize the clusters.
	clusters := make([][][]float64, 0)

	// Initialize the distance matrix.
	distanceMatrix := make([][]float64, len(dataSet))
	for i := 0; i < len(dataSet); i++ {
		distanceMatrix[i] = make([]float64, len(dataSet))
	}

	// Update the distance matrix.
	for i := 0; i < len(dataSet); i++ {
		for j := 0; j < len(dataSet); j++ {
			distanceMatrix[i][j] = distance(dataSet[i:i+1], dataSet[j:j+1])
		}
	}

	// Initialize the visited points.
	visited := make([]bool, len(dataSet))

	// Initialize the noise points.
	noise := make([]bool, len(dataSet))

	// Iterate over the data set.
	for i := 0; i < len(dataSet); i++ {
		// Check if the point has been visited.
		if visited[i] {
			continue
		}

		// Mark the point as visited.
		visited[i] = true

		// Initialize the neighbors.
		neighbors := make([][]float64, 0)

		// Find the neighbors.
		for j := 0; j < len(dataSet); j++ {
			if distanceMatrix[i][j] <= epsilon {
				neighbors = append(neighbors, dataSet[j])
			}
		}

		// Check if the point is a noise point.
		if len(neighbors) < minPts {
			noise[i] = true
			continue
		}

		// Initialize the cluster.
		cluster := make([][]float64, 0)

		// Add the point to the cluster.
		cluster = append(cluster, dataSet[i])

		// Iterate over the neighbors.
		for j := 0; j < len(neighbors); j++ {
			// Check if the neighbor has been visited.
			if visited[j] {
				continue
			}

			// Mark the neighbor as visited.
			visited[j] = true

			// Initialize the neighbor's neighbors.
			neighbors2 := make([][]float64, 0)

			// Find the neighbor's neighbors.
			for l := 0; l < len(dataSet); l++ {
				if distanceMatrix[j][l] <= epsilon {
					neighbors2 = append(neighbors2, dataSet[l])
				}
			}

			// Check if the neighbor is a noise point.
			if len(neighbors2) < minPts {
				noise[j] = true
				continue
			}

			// Add the neighbor to the cluster.
			cluster = append(cluster, dataSet[j])

			// Add the neighbor's neighbors to the neighbors.
			neighbors = append(neighbors, neighbors2...)
		}

		// Add the cluster to the clusters.
		clusters = append(clusters, cluster)
	}

	// Initialize the noise cluster.
	noiseCluster := make([][]float64, 0)

	// Add the noise points to the noise cluster.
	for i := 0; i < len(dataSet); i++ {
		if noise[i] {
			noiseCluster = append(noiseCluster, dataSet[i])
		}
	}

	// Add the noise cluster to the clusters.
	clusters = append(clusters, noiseCluster)

	// Return the clusters.
	return clusters
}

/*
 * Main function.
 */
func main() {
	// Initialize the data set.
	dataSet := [][]float64{
		{1.0, 1.0},
		{1.5, 2.0},
		{3.0, 4.0},
		{5.0, 7.0},
		{3.5, 5.0},
		{4.5, 5.0},
		{3.5, 4.5},
	}

	// Initialize the distance function.
	distance := func(x [][]float64, y [][]float64) float64 {
		sum := 0.0
		for i := 0; i < len(x[0]); i++ {
			sum += math.Pow(x[0][i]-y[0][i], 2)
		}
		return math.Sqrt(sum)
	}

	// Initialize the clusters.
	clusters := KMeans(dataSet, 2, 100, distance, nil)

	// Print the clusters.
	for i := 0; i < len(clusters); i++ {
		fmt.Println(clusters[i])
	}
}
