package main

import "fmt"

func main() {
	ar := []int{52, 32, 51, 25, 55, 2, 12, 8}
	// BubleSort(ar)
	// fmt.Println(ar)
	// SelectionSort(ar)
	// fmt.Println(ar)
	// InsertionSort(ar)
	// fmt.Println(ar)
	// MergeSort(ar,0,len(ar)-1)
	// fmt.Println(ar)

	QuickSort(ar, 0, len(ar)-1)
	fmt.Println(ar)
}
