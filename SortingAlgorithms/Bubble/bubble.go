package main

import "fmt"

func main() {
	nums := []int{10, 4, 2, 6, 7, 4, 2}
	BubbleSort(nums)
	fmt.Println(nums)
}

func BubbleSort(nums []int) {
	for i := 0; i < len(nums); i++ {
		for j := 0; j < len(nums)-i-1; j++ {
			if nums[j] > nums[j+1] {
				nums[j], nums[j+1] = nums[j+1], nums[j]
			}
		}
	}
}
