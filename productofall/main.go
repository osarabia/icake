package main

import "fmt"

/*
You have a list of integers, and for each index you want to find the product of every integer except the integer at that index.

Write a function get_products_of_all_ints_except_at_index() that takes a list of integers and returns a list of the products.

For example, given:

  [1, 7, 3, 4]

your function would return:

  [84, 12, 28, 21]

by calculating:

  [7 * 3 * 4,  1 * 3 * 4,  1 * 7 * 4,  1 * 7 * 3]

Here's the catch: You can't use division in your solution!
*/


func ProductOfAll(numbers []int) []int{
	productsOfAll := make([]int, len(numbers))

	productsOfAll[0] = 1
	for r := 1; r < len(numbers); r++ {
		productsOfAll[r] = productsOfAll[r - 1] * numbers[r - 1]
	}

	accum := 1
	for l := len(numbers) - 2; l >= 0; l--{
		accum = accum * numbers[l+1]
		productsOfAll[l] *=  accum
	}

	return productsOfAll
}

func main() {
	numbers := []int{1, 7, 3, 4}
	fmt.Printf("%v\n",ProductOfAll(numbers))

        numbers1 := []int{1, 2, 3}
	fmt.Printf("%v\n",ProductOfAll(numbers1))

        numbers2 := []int{8, 2, 4, 3, 1, 5}
	fmt.Printf("%v\n",ProductOfAll(numbers2))

        numbers4 := []int{6, 2, 0, 3}
	fmt.Printf("%v\n",ProductOfAll(numbers4))

	numbers5 := []int{4, 0, 9, 1, 0}
        fmt.Printf("%v\n",ProductOfAll(numbers5))

        numbers6 := []int{-3, 8, 4}
        fmt.Printf("%v\n",ProductOfAll(numbers6))

        numbers7 := []int{-7, -1, -4, -2}
        fmt.Printf("%v\n",ProductOfAll(numbers7))

}
