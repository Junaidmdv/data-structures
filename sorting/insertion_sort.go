package main

// 🔹 Example (Sorting [5, 3, 4, 1, 2]):

// Start with 5 → already sorted.

// Take 3 → insert before 5 → [3, 5, 4, 1, 2]

// Take 4 → insert between 3 and 5 → [3, 4, 5, 1, 2]

// Take 1 → insert at start → [1, 3, 4, 5, 2]

// Take 2 → insert after 1 → [1, 2, 3, 4, 5] ✅

func InsertionSort(ar []int) {

	for i := 1; i < len(ar); i++ {
		temp := ar[i]
		prev := i - 1

		for prev >= 0 && temp < ar[prev] {
			ar[prev+1] = ar[prev]
			prev--
		}

		ar[prev+1] = temp
	}
}
