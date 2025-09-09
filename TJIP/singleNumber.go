package main

import "fmt"

// func singelNumber(nums []int) int {
// 	var winner, voteCount int

// 	for _, num := range nums {
// 		if voteCount == 0 {
// 			winner = num
// 		}
// 		if num == winner {
// 			voteCount--
// 		} else {
// 			voteCount++
// 		}
// 	}
// 	return winner
// }

// this problem set-based
func singleNumber(nums []int) int {
	set := make(map[int]struct{})
	for _, num := range nums {
		if _, exists := set[num]; !exists {
			set[num] = struct{}{}
		} else {
			delete(set, num)
		}
	}

	for value := range set {
		return value
	}
	return -1
}

func main() {
	result := singleNumber([]int{2, 2, 1})
	fmt.Println(result)
}
