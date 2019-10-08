package main

import (
	"fmt"
	"strings"
)

func main() {
	input1 := "omar"
	fmt.Printf("%v \n", reverseStringInPlace(input1))

	input2 := "mar"
        fmt.Printf("%v \n", reverseStringInPlace(input2))

        input3 := ""
        fmt.Printf("%v \n", reverseStringInPlace(input3))

}

func reverseStringInPlace(input string) string {

	inputData := strings.Split(input, "")
	startIndex, endIndex := 0, len(inputData) -1

	for startIndex < endIndex {
		inputData[startIndex], inputData[endIndex] = inputData[endIndex], inputData[startIndex]
		startIndex += 1
		endIndex -= 1
	}

	return strings.Join(inputData, "")
}
