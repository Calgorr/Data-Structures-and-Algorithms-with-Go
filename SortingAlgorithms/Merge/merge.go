package main

import "fmt"

func main() {
	a := []int{8, 5, 4, 523, 23, 236, 235}
	fmt.Println(MergeSort(a))
}

func MergeSort(nums []int) []int {
	if len(nums) == 1 {
		return nums
	}
	arr1 := nums[:len(nums)/2]
	arr2 := nums[len(nums)/2:]

	arr1 = MergeSort(arr1)
	arr2 = MergeSort(arr2)

	return merge(arr1, arr2)
}

func merge(a, b []int) []int {
	c := []int{}
	for len(a) > 0 && len(b) > 0 {
		if a[0] > b[0] {
			c = append(c, b[0])
			b = b[1:]
		} else {
			c = append(c, a[0])
			a = a[1:]
		}
	}
	for len(a) > 0 {
		c = append(c, a[0])
		a = a[1:]
	}
	for len(b) > 0 {
		c = append(c, b[0])
		b = b[1:]
	}
	return c
}
