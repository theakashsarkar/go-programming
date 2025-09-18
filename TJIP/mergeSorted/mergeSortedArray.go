package mergeSorted

func merge(nums1 []int, m int, nums2 []int, n int) []int { //v1[1, 2, 3, 0, 0, 0]  v2[2, 5, 6]
	left, right := 0, 0
	newArr := []int{} // m + n - 1
	for left < m && right < n {
		if nums1[left] < nums2[right] {
			newArr = append(newArr, nums1[left])
			left++
		} else {
			newArr = append(newArr, nums2[right])
			right++
		}
	}
	newArr = append(newArr, nums1[left:m]...)
	newArr = append(newArr, nums2[right:n]...)
	return newArr
}

// func main() {
// 	fmt.Println(merge([]int{1, 2, 3, 0, 0, 0}, 3, []int{2, 5, 6}, 3))
// }
