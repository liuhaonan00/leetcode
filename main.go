package main

import (
	"fmt"
	"leetcode/array_question"
)

func main() {
	nums := []int{3, 2, 2, 3}
	val := 3

	res := array_question.RemoveElement(nums, val)
	fmt.Println(res)

	fmt.Println(nums)
}
