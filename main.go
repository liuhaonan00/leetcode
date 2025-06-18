package main

import (
	"fmt"
	"leetcode/array_question"
)

func main() {
	nums1 := []int{0}
	m := 0

	nums2 := []int{2}
	n := 1

	array_question.Merge(nums1, m, nums2, n)

	fmt.Println(nums1)
}
