package main

func merge(nums1 []int, m int, nums2 []int, n int) {

	for j := 0; j < n; j++ {
		nums1[m+j] = nums2[j]
	}

	for i := 0; i < m+n; i++ {
		for j := i + 1; j < m+n; j++ {
			if nums1[i] > nums1[j] {
				nums1[i], nums1[j] = nums1[j], nums1[i]
			}
		}
	}
	//fmt.Printf("%v", nums1)

}
