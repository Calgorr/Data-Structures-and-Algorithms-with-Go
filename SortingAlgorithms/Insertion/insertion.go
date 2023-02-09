package main

import "fmt"

func main() {
	nums := []int{10, 4, 2, 6, 7, 4, 2}
	InsertionSort(nums)
	fmt.Println(nums)
}

func InsertionSort(nums []int) {
	for i := 0; i < len(nums)-1; i++ {
		key := nums[i+1]
		j := i
		for j >= 0 && nums[j] > key {
			nums[j+1] = nums[j]
			j -= 1
		}
		nums[j+1] = key
	}
}
