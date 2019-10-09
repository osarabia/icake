package main

import (
	"fmt"
	"strings"
)

func main() {
	input1 := strings.Split("hello world", "")
	input1 = reverseWords(input1)
        fmt.Printf("%v\n", input1)

        input2 := strings.Split("vault", "")
	input2 = reverseWords(input2)
        fmt.Printf("%v\n", input2)

	input3 := strings.Split("thief cake", "")
	input3 = reverseWords(input3)
        fmt.Printf("%v\n", input3)

        input4 := strings.Split("thief cake", "")
	input4 = reverseWords(input4)
        fmt.Printf("%v\n", input4)

        input5 := strings.Split("one another get", "")
	input5 = reverseWords(input5)
        fmt.Printf("%v\n", input5)

}

func reverseWords(input []string) []string {
	startIndex, endIndex := 0, len(input) - 1
	//loop to reverse whole input
	for startIndex < endIndex {
		input[startIndex], input[endIndex] = input[endIndex], input[startIndex]
		startIndex += 1
		endIndex -= 1
	}

	startIndex = 0
	for i,v := range input {
		if v == " " || i == len(input) - 1 {
			if  v == " " {
			        endIndex = i -1
		        }

			if i == len(input) - 1 {
				endIndex = len(input) - 1
			}

			//loop to reverse each word
			for  startIndex < endIndex {
                                input[startIndex], input[endIndex] = input[endIndex], input[startIndex]
		                startIndex += 1
		                endIndex -= 1
			}

			//set startIndex
			startIndex = i + 1
		}
	}

	return input
}
