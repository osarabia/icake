package main

import (
    "fmt"
    "math/rand"
    "time"
)

/*
Write a function for doing an in-place ↴ shuffle of a list.

The shuffle must be "uniform," meaning each item in the original list must have the same probability of ending up in each spot in the final list.

Assume that you have a function get_random(floor, ceiling) for getting a random integer that is >= floor and <= ceiling.
*/

func random(min int, max int) int {
    value :=  max - min
    if value == 0 {
	    return rand.Intn(max)
    }
    return rand.Intn(max-min) + min
}

func main() {
    rand.Seed(time.Now().UnixNano())

    input := []int{5,3,7,1,9,6}

    fmt.Printf("%v\n", input)

    fmt.Printf("suffle:%v", suffleInPlace(input))
}

func suffleInPlace(input []int) []int{

	var randNumber int
	lastItem := len(input) - 1

	for i,_ := range input{
		randNumber = random(i, lastItem)
		if i != randNumber {
			input[i], input[randNumber] = input[randNumber], input[i]
		}
	}

	return input
}
