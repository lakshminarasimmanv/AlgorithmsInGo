package main

import "fmt"

func main() {
	arr := []int{1, 43, 34, 54, 23, 2, 65}
	x := 66

	result := binary_search(0, (len(arr) - 1), x, arr)

	if result != -1 {
		fmt.Printf("Element %v is present at index %v\n", x, result)
	} else {
		fmt.Printf("The Element %v is not present in the given list!\n", x)
	}
}

func binary_search(low, high, x int, arr []int) int {

	if high >= low {

		mid := (high + low) / 2

		if arr[mid] == x {
			return mid
		} else if arr[mid] > x {
			return binary_search(low, mid-1, x, arr)
		} else {
			return binary_search(mid+1, high, x, arr)
		}

	}
	return -1
}
