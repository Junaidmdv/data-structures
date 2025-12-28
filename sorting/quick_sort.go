package main

func Qs(ar []int, st, end int) int {
	i := st
	j := end
	pivote := ar[st]

	for i < j {
		for i < end && ar[i] <= pivote {
			i++
		}

		for j > st && ar[j] > pivote {
			j--
		}

		if i < j {
			ar[i], ar[j] = ar[j], ar[i]
		}

	}

	ar[st], ar[j] = ar[j], ar[st]

	return j
}

func QuickSort(ar []int, st, end int) {

	if st < end {
		pivote := Qs(ar, st, end)
		QuickSort(ar, st, pivote-1)
		QuickSort(ar, pivote+1, end)
	}

}
