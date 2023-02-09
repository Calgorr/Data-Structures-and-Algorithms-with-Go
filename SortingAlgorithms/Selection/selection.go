package main

import (
	"fmt"
)

func main() {
	nums := []int{10, 4, 2, 6, 7, 4, 2}
	SelectionSort(nums)
	fmt.Println(nums)
}

func SelectionSort(nums []int) {
	for i := 0; i < len(nums); i++ {
		index := minElem(nums, i)
		nums[index], nums[i] = nums[i], nums[index]
	}
}

func minElem(nums []int, index int) int {
	min := nums[index]
	for i := index + 1; i < len(nums); i++ {
		if nums[i] < min {
			min = nums[i]
			index = i
		}
	}
	return index
}
