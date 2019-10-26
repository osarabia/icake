package main

import "fmt"
/*
Find a duplicate, Space Edition™.

We have a list of integers, where:

The integers are in the range 1..n1..n
The list has a length of n+1n+1
It follows that our list has at least one integer which appears at least twice. But it may have several duplicates, and each duplicate may appear more than twice.

Write a function which finds an integer that appears more than once in our list. (If there are multiple duplicates, you only need to find one of them.)

We're going to run this function on our new, super-hip MacBook Pro With Retina Display™. Thing is, the damn thing came with the RAM soldered right to the motherboard, so we can't upgrade our RAM. So we need to optimize for space!
*/

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
