package main

import "fmt"

func main() {
        slice1 := []int{2, 4, 6}
        slice2 := []int{1, 3, 7}

	output := mergeSlices(slice1, slice2)
	fmt.Printf("%v\n", output)

	slice1 = []int{3, 4, 6, 10, 11, 15}
        slice2 = []int{1, 5, 8, 12, 14, 19}

	output = mergeSlices(slice1, slice2)
	fmt.Printf("%v\n", output)

        slice1 = []int{2, 4, 6, 8}
        slice2 = []int{1, 7}

	output = mergeSlices(slice1, slice2)
	fmt.Printf("%v\n", output)

}

func mergeSlices(slice1 []int, slice2 []int) []int {
        index1 := 0
	index2 := 0
        output := make([]int, 0, len(slice1) -1 + len(slice2) - 1)

	for len(slice1) > 0 && len(slice2) > 0 && index1 < len(slice1) && index2 < len(slice2) {
		if slice1[index1] < slice2[index2] {
			output = append(output, slice1[index1])
			index1 += 1
		} else {
                        output = append(output, slice2[index2])
			index2 += 1
		}
	}

	for index1 < len(slice1) {
		output = append(output, slice1[index1])
		index1 += 1
	}

	for index2 < len(slice2) {
		output = append(output, slice2[index2])
		index2 += 1
	}

	return output
}
