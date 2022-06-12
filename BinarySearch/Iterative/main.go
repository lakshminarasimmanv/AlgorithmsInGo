package main

import "fmt"

func main() {
	arr := []int{1, 5, 8, 13, 23, 55, 67, 85, 98, 104}
	val := 85

	result := binary_search(val, arr)

	if result != -1 {
		fmt.Printf("Element %v is present at index %v.\n", val, result)
	} else {
		fmt.Printf("The Element %v is not present in the given list!\n", val)
	}
}

func binary_search(x int, data []int) int {
	low := 0
	high := len(data) - 1

	for mid := 0; low <= high; mid = (high + low) / 2 {

		if data[mid] < x {
			low = mid + 1
		} else if data[mid] > x {
			high = mid - 1
		} else {
			return mid
		}
	}
	return -1
}
