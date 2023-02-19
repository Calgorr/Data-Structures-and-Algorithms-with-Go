package quick

func QuickSort(nums []int, low, high int) {
	if low < high {
		pivot := Partition(nums, low, high)
		QuickSort(nums, pivot+1, high)
		QuickSort(nums, low, pivot-1)
	}
}

func Partition(nums []int, low, high int) int {
	pivot := nums[high-1]
	i := low - 1
	for j := low; j <= high-1; j++ {
		if nums[j] < pivot {
			i++
			nums[i], nums[j] = nums[j], nums[i]
		}
	}
	nums[i+1], nums[high-1] = nums[high-1], nums[i+1]
	return i + 1
}
