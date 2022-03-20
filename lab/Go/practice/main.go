package main

import (
	"fmt"
	"sort"
)

func main() {

	fmt.Println(findMedianSortedArrays([]int{1, 2, 2}, []int{1, 2, 3}))
}

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {

	intSlice := append(nums1, nums2...)

	sort.Ints(intSlice)
	keys := make(map[int]bool)
	s1 := []int{}
	for _, entry := range intSlice {
		if _, value := keys[entry]; !value {
			keys[entry] = true
			s1 = append(s1, entry)
		}
	}

	var a int
	var l = len(s1)
	var k int

	if l == 1 {
		a = s1[0]
		k = 1

	} else if l == 2 {
		a = s1[0] + s1[1]
		k = 2

	} else if l%2 != 0 {
		a = s1[(l-1)/2]
		k = 1

	} else {
		for i := 1; i < l-1; i++ {
			a += s1[i]
			k++
		}
	}

	return float64(int64(a)) / float64(k)

}
