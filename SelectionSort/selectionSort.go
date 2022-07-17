/* Selection Sort. */

package main

import "fmt"

func main() {
	var n, i, j, min, temp int
	fmt.Print("Enter the number of elements: ")
	fmt.Scan(&n)
	a := make([]int, n)
	fmt.Println("Enter the elements:")
	for i = 0; i < n; i++ {
		fmt.Scan(&a[i])
	}
	for i = 0; i < n-1; i++ {
		min = i
		for j = i + 1; j < n; j++ {
			if a[j] < a[min] {
				min = j
			}
		}
		temp = a[i]
		a[i] = a[min]
		a[min] = temp
	}
	fmt.Println("Sorted array is:")
	for i = 0; i < n; i++ {
		fmt.Printf("%d ", a[i])
	}
	fmt.Println()
}
