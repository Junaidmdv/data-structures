package main

func BubleSort(ar []int) []int {

	for i := 0; i < len(ar); i++ {
		for j := 0; j < len(ar)-i-1; j++ {
			if ar[j] > ar[j+1] {
				ar[j], ar[j+1] = ar[j+1], ar[j]
			}
		}
	}
	return ar
}
