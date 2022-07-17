package main

import "fmt"

func main() {
	arr := []int{23, 43, 22, 12, 54, 67, 16, 67, 98, 76}
	fmt.Println(bubble_sort(arr))
}

func bubble_sort(arr []int) []int {
	swapped := true

	for swapped {
		swapped = false

		for i := 0; i < len(arr)-1; i++ {
			if arr[i] > arr[i+1] {
				arr[i], arr[i+1] = arr[i+1], arr[i]
				swapped = true
			}
		}
	}
	return arr
}
