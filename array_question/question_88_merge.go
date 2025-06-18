package array_question

//[1, 3, 5, 7, _, _, _]  [2, 5, 6]
//>                   >
//
//[1, 3, 5, _, _, _, 7]  [2, 5, 6]
//>                      >
//
//[1, 3, 5, _, _, 6, 7]  [2, 5, 6]
//>                   >
//
//[1, 3, 5, _, 5, 6, 7]  [2, 5, 6]
//>                >
//
//
//[1, 3, _, 5, 5, 6, 7]  [2, 5, 6]
//>                   >
//
//[1, 3, 5, 7, _, _, _]  [2, 5, 6]
//>                   >

// 倒序插入， 复杂度 o(N), 注意边界问题
func merge(nums1 []int, m int, nums2 []int, n int) {
	i, j := m-1, n-1

	for k := len(nums1) - 1; k >= 0; k-- {
		if j < 0 {
			break
		}

		if i < 0 || nums1[i] <= nums2[j] {
			nums1[k] = nums2[j]
			j--
			continue
		}

		nums1[k] = nums1[i]
		i--
	}
}

func Merge(nums1 []int, m int, nums2 []int, n int) {
	merge(nums1, m, nums2, n)
}
