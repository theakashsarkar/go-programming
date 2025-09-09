package main

import "fmt"

func firstMissingPositive(nums []int) {
	set := make(map[int]struct{})
	for _, num := range nums {
		set[num] = struct{}{}
	}
	for i := 0; i <= len(nums)+1; i++ {
		if _, exists := set[i]; !exists {
			return i
		}
	}
	return 0
}

func main() {
	fmt.Println(firstMissingPositive([]int{1, 2, 0}))
}
