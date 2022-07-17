/* Linerar Search. */

package main

import "fmt"

func main() {
	var arr = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	var key int
	fmt.Println("Enter the key to be searched: ")
	fmt.Scanf("%d", &key)
	var flag int
	for i := 0; i < len(arr); i++ {
		if arr[i] == key {
			fmt.Println("Key found at index: ", i)
			flag = 1
			break
		}
	}
	if flag == 0 {
		fmt.Println("Key not found")
	}
}
