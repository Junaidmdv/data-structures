package main

import "fmt"

func Heapify(ar []int, n, idx int) {
	max := idx
	left := idx*2 + 1
	right := idx*2 + 2

	if left < n && ar[max] < ar[left] {
		max = left
	}
	if right < n && ar[max] < ar[right] {
		max = right
	}

	if max != idx {
		ar[idx], ar[max] = ar[max], ar[idx]
		Heapify(ar, n, max)
	}

}

func HeapSort(ar []int) {
	n := len(ar)
	for i := n/2 - 1; i >= 0; i-- {
		Heapify(ar, n, i)
	}

	for i := n - 1; i >= 0; i-- {
		ar[0], ar[i] = ar[i], ar[0]
		Heapify(ar, i, 0)

	}
}

func main() {

	ar := []int{2, 51, 2, 51, 234, 12, 3}
	HeapSort(ar)
	fmt.Println(ar)
}
