package main

import "fmt"

// keep swapping till the largest element is at the end (update end each time)
func bubbleSort(a []int) {
	for i := len(a) - 1; i > 0; i-- {
		swap := false
		for j := 0; j < i; j++ {
			if a[j+1] < a[j] {
				a[j], a[j+1] = a[j+1], a[j]
				swap = true
			}
		}
		if !swap {
			break
		}
	}
}

func main() {
	a := []int{23, 12, 44, 223, 123, 32}
	fmt.Println("Before sorting: ", a)
	bubbleSort(a)
	fmt.Println("After sorting: ", a)
}
