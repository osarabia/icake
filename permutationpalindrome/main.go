package main


import "fmt"

/*
Write an efficient function that checks whether any permutation ↴ of an input string is a palindrome. ↴

You can assume the input string only contains lowercase letters.

Examples:

"civic" should return True
"ivicc" should return True
"civil" should return False
"livci" should return False
*/


func main() {
	input1 := "civic"

	fmt.Printf("%v\n", isPermutationPalindrome(input1))

	input2 := "civil"

	fmt.Printf("%v\n", isPermutationPalindrome(input2))

        input3 := "aabcbcd"

	fmt.Printf("%v\n", isPermutationPalindrome(input3))

        input4 := "aabccbdd"

	fmt.Printf("%v\n", isPermutationPalindrome(input4))

	input5 := "aabcd"

        fmt.Printf("%v\n", isPermutationPalindrome(input5))

	input6 := "aabbcd"

        fmt.Printf("%v\n", isPermutationPalindrome(input6))
}

func isPermutationPalindrome(input string) bool {
        if len(input) == 0 || len(input) == 1 {
		return true
	}

	counter := make(map[rune]int)

	for _, c := range input {
		counter[c] += 1
	}

	if len(input) % 2 == 0 {
		for _, v := range counter {
			if v % 2 != 0 {
				return false
			}
		}
	} else {
		oddCounter := 0
                for _, v := range counter {
			if v % 2 != 0 && oddCounter > 0 {
				return false
			} else if v % 2 != 0 {
				oddCounter += 1
			}
                }
	}

	return true
}
