package main

import "fmt"

// find the maximum sum of subarray with size k

func Maxsum(ar []int, k int) {

	// this method is used fixed sliding window pattern.
	//  Where the size of sliding window will be same.
	maxsum := 0
	sum := 0

	for i := 0; i < k; i++ {
		sum += ar[i]
	}

	for i := k; i < len(ar); i++ {
		sum += ar[i] - ar[i-k]

		if sum > maxsum {
			maxsum = sum
		}
	}

	fmt.Println(maxsum)
}


// Question 2 Longest Repeating Character Replacement(leetcode 424)

//brute force 








func main() {
	ar:=[]int{1,1,3,5,-2,2,-2}
	Maxsum(ar,3)

}
