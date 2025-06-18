package array_question

func removeElement(nums []int, val int) int {
	start, end := 0, len(nums)-1
	result := 0
	for start <= end {
		//fmt.Printf("start: %d, nums[start]: %d end: %d, nums[end]: %d\n", start, nums[start], end, nums[end])
		if nums[end] == val {
			end--
			continue
		}

		if nums[start] == val {
			nums[start] = nums[end]
			start++
			end--
			result++
			continue
		}

		result++
		start++
	}
	return result
}

func RemoveElement(nums []int, val int) int {
	return removeElement(nums, val)
}
