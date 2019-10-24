package main

import "fmt"


func main() {
	input1 := []int{1, 2, 3, 2}
	fmt.Printf("%v\n", findDuplicate(input1))

	input2 := []int{1, 1}
	fmt.Printf("%v\n", findDuplicate(input2))

	input3 := []int{1, 2, 5, 5, 5, 5}
	fmt.Printf("%v\n", findDuplicate(input3))

	input4 := []int{4, 1, 4, 8, 3, 2, 7, 6, 5}
	fmt.Printf("%v\n", findDuplicate(input4))

	input5 := []int{1,2,3,4,5,5}
	fmt.Printf("%v\n", findDuplicate(input5))
}


func findDuplicate(numbers []int) int {
	var floor, ceiling, mid  int
	floor = 1
	ceiling = len(numbers) - 1

	var lowerRangeFloor, lowerRangeCeiling, upperRangeFloor, upperRangeCeiling, distinctInLowerRange, lowerCount int

	for floor < ceiling {
		mid = floor + ((ceiling - floor) / 2)

		lowerRangeFloor, lowerRangeCeiling = floor, mid
		upperRangeFloor, upperRangeCeiling = mid + 1, ceiling
		lowerCount = 0

		for _, v := range numbers {
			if v >= lowerRangeFloor && v <= lowerRangeCeiling {
				lowerCount += 1
			}
		}

		distinctInLowerRange = lowerRangeCeiling - lowerRangeFloor + 1

		if lowerCount > distinctInLowerRange  {
			//go to left
			floor, ceiling = lowerRangeFloor, lowerRangeCeiling
		} else {
			//go to right
			floor, ceiling = upperRangeFloor, upperRangeCeiling
		}
	}

	return floor
}
