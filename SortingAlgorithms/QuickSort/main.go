package main

import "fmt"

func main() {
	arr := []int{4, 65, 2, 34, 77, 3, 8, 44}

	index, result := partition(arr, 1, (len(arr) - 1))
	fmt.Println(result)
	fmt.Println(index)
}

func partition(arr []int, l, r int) (int, []int) {
	pivot := arr[r]
	i := l - 1

	for j := l; j < r; j++ {
		if arr[j] < pivot {
			arr[i], arr[j] = arr[j], arr[i]
			i++
		}
	}
	arr[i+1], arr[r] = arr[r], arr[i+1]
	return i, arr
}
