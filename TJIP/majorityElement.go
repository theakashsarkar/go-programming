package main

import "fmt"

func majorityElement(nums []int) int {
	var winner, voteCount int
	for _, num := range nums {
		if voteCount == 0 {
			winner = num
		}
		if num == winner {
			voteCount++
		} else {
			voteCount--
		}
		// voteCount += map[bool]int{true: 1, false: -1}[num == winner]
	}
	return winner
}

func main() {
	result := majorityElement([]int{3, 2, 3})
	fmt.Println(result)
}
