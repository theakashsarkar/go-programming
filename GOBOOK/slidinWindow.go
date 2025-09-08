package main

import "fmt"

// Corrected function
func MaxSumArray(num []int, k int) int {
	if len(num) < k || k <= 0 {
		return 0
	}

	windowSlid := 0
	for i := 0; i < k; i++ {
		windowSlid += num[i]
	}

	maxNum := windowSlid

	for i := k; i < len(num); i++ {
		windowSlid += num[i] - num[i-k]
		if windowSlid > maxNum {
			maxNum = windowSlid
		}
	}

	return maxNum
}

func main() {
	arr := []int{1, 5, 1, 2, 7, 8}
	k := 3
	result := MaxSumArray(arr, k)
	fmt.Println("Maximum subarray sum:", result)
}
