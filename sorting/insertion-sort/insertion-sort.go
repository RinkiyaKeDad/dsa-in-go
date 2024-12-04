package main

import "fmt"

func insertionSort(nums []int) {
	for i := 1; i < len(nums); i++ {
		key := nums[i]
		j := i - 1
		for ; j >= 0 && nums[j] > key; j-- {
			nums[j+1] = nums[j]
		}
		nums[j+1] = key
	}
}

func main() {
	nums := []int{3, 9, 4, 5, 8, 7, 1}
	fmt.Println("Before sorting: ", nums)
	insertionSort(nums)
	fmt.Println("After sorting: ", nums)
}
