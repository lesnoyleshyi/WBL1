package main

import (
	"fmt"
	"log"
)

const searchedVal int = 32

func main() {
	sortedNums := make([]int, 100, 100)
	for i, _ := range sortedNums {
		sortedNums[i] = i * 2
	}

	pos, ok := bSearch(sortedNums, searchedVal)
	if !ok {
		log.Println("No such value is in array")
	} else {
		fmt.Printf("Number %d is on %d position\n", searchedVal, pos)
	}
}

func bSearch(arr []int, val int) (int, bool) {
	var min, mid int
	max := len(arr) - 1

	for min <= max {
		mid = (min + max) / 2
		if arr[mid] == val {
			return mid, true
		}
		if arr[mid] < val {
			min = mid + 1
		}
		if arr[mid] > val {
			max = mid - 1
		}
	}
	return 0, false
}
