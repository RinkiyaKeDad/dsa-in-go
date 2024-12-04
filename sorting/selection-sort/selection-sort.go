package main

import "fmt"

// store min_index and then swap at the end (update start each time)
func sortArray(nums []int) {
	for i := 0; i < len(nums)-1; i++ {
		min_index := i
		for j := i + 1; j < len(nums); j++ {
			if nums[j] < nums[min_index] {
				min_index = j
			}
		}
		nums[i], nums[min_index] = nums[min_index], nums[i]
	}

}

func main() {
	nums := []int{64, 23, 42, 19, 22}
	fmt.Println("Before Sorting: ", nums)
	sortArray(nums)
	fmt.Println("After Sorting: ", nums)
}
