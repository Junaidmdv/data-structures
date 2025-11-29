package main

func Merge(ar []int, st, mid, end int) {
	i := st
	j := mid + 1
	var tempAr []int

	for i <= mid && j <= end {
		if ar[i] <= ar[j] {
			tempAr = append(tempAr, ar[i])
			i++
		} else {
			tempAr = append(tempAr, ar[j])
			j++
		}
	}

	//

	for i <= mid {
		tempAr = append(tempAr, ar[i])
		i++
	}

	for j <= end {
		tempAr = append(tempAr, ar[j])
		j++
	}

	for i := 0; i < len(tempAr); i++ {
		ar[i+st] = tempAr[i]
	}

}

func MergeSort(ar []int, st, end int) {
	if st < end {
		mid := st + (end-st)/2
		MergeSort(ar, st, mid)
		MergeSort(ar, mid+1, end)
		Merge(ar, st, mid, end)
	}
}
