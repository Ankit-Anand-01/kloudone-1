package main

import "fmt"

func main() {
	var sample []int = []int{3, 4, 5, 2, 1, 7, 8, -1, -3}
	fmt.Println("Before Sorting")
	for _, val := range sample {
		fmt.Println(val)
	}
	selectionSort(sample)
}

func selectionSort(arr []int) {
	len := len(arr)
	for i := 0; i < len-1; i++ {
		minIndex := i
		for j := i + 1; j < len; j++ {
			if arr[j] < arr[minIndex] {
				arr[j], arr[minIndex] = arr[minIndex], arr[j]
			}
		}
	}
	fmt.Println("\nAfter SelectionSort")
	for _, val := range arr {
		fmt.Println(val)
	}
}
