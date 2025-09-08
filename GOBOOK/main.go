package main

import (
	"fmt"
	"gobook\gobook"
)

func main() {
	arr := []int{1, 5, 1, 2, 7, 8}
	k := 3
	result := gobook.maxSumArray(arr, k)
	fmt.Println("Maximum subarray sum:", result)
}
