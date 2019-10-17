package main

import (
	"fmt"
	"math"
)

func MaxProductOf3(numbers []float64) float64 {
	var maxProductOf3, highestProductOf2, highest, lowestProductOf2, lowest, current float64

	maxProductOf3 = -100000000

	highest = numbers[0]

	if highest > numbers[1] {
		lowest = numbers[1]
	} else {
		lowest = highest
		highest = numbers[1]
		highestProductOf2 = lowest * highest
                lowestProductOf2 = lowest * highest
	}

	for i := 2; i < len(numbers); i++ {
		current = numbers[i]
		if i == 2 {
			maxProductOf3 = lowest * highest * current
			highestProductOf2 = math.Max(highestProductOf2, lowest * current)
                        highestProductOf2 = math.Max(highestProductOf2, highest * current)

			lowestProductOf2 = math.Min(lowestProductOf2, lowest * current)
                        lowestProductOf2 = math.Min(lowestProductOf2, highest * current)

			if current <  lowest {
				lowest = current
			}

			if current > highest {
				highest = current
			}
		} else {

			maxProductOf3 = math.Max(maxProductOf3, highestProductOf2 * current)
                        maxProductOf3 = math.Max(maxProductOf3, lowestProductOf2 * current)

			highestProductOf2 = math.Max(highestProductOf2, lowest * current)
                        highestProductOf2 = math.Max(highestProductOf2, highest * current)

			lowestProductOf2 = math.Min(lowestProductOf2, lowest * current)
			lowestProductOf2 = math.Min(lowestProductOf2, highest * current)

                        if current <  lowest {
				lowest = current
			}

			if current > highest {
				highest = current
			}

		}
	}

	return maxProductOf3
}

func main() {
	numbers1 := []float64{1.0, 2.0, 3.0, 4.0}
	fmt.Printf("%v\n", MaxProductOf3(numbers1))

        numbers2 := []float64{6.0, 1.0, 3.0, 5.0, 7.0, 8.0, 2.0}
	fmt.Printf("%v\n", MaxProductOf3(numbers2))

	numbers3 := []float64{-5.0, 4.0, 8.0, 2.0, 3.0}
        fmt.Printf("%v\n", MaxProductOf3(numbers3))

	numbers4 := []float64{-10.0, 1.0, 3.0, 2.0, -10.0}
        fmt.Printf("%v\n", MaxProductOf3(numbers4))

	numbers5 := []float64{-5.0, -1.0, -3.0, -2.0}
	fmt.Printf("%v\n", MaxProductOf3(numbers5))
}
