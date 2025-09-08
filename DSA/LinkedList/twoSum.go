package twoSum

import (
	"fmt"
)

func TwoSum(nums []int, target int) []int {
	seen := make(map[int]int)
	for i, num := range nums {
		remaining := target - num
		if idx, ok := seen[remaining]; ok {
			return []int{idx, i}
		}
		seen[num] = i
	}
	return []int{}
}

func main() {
	result := TwoSum([]int{2, 7, 11, 15}, 9)
	fmt.Println(result)
}
