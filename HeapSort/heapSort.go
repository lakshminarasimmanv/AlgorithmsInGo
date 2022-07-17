/* Heap Sort. */

package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	slice := generateSlice(20)
	fmt.Println("\n--- Unsorted --- \n\n", slice)
	heapsort(slice)
	fmt.Println("\n--- Sorted ---\n\n", slice)
}

// Generates a slice of size, size filled with random numbers
func generateSlice(size int) []int {

	slice := make([]int, size, size)
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < size; i++ {
		slice[i] = rand.Intn(999) - rand.Intn(999)
	}
	return slice
}

func heapsort(a []int) {
	buildMaxHeap(a)
	for i := len(a) - 1; i >= 0; i-- {
		a[0], a[i] = a[i], a[0]
		maxHeapify(a, 0, i)
	}
}

func buildMaxHeap(a []int) {
	for i := len(a)/2 - 1; i >= 0; i-- {
		maxHeapify(a, i, len(a))
	}
}

func maxHeapify(a []int, i, length int) {
	left := 2*i + 1
	right := 2*i + 2
	largest := i
	if left < length && a[left] > a[largest] {
		largest = left
	}
	if right < length && a[right] > a[largest] {
		largest = right
	}
	if largest != i {
		a[i], a[largest] = a[largest], a[i]
		maxHeapify(a, largest, length)
	}
}
