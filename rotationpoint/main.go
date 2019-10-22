package main

import "fmt"

/*
I want to learn some big words so people think I'm smart.

I opened up a dictionary to a page in the middle and started flipping through, looking for words I didn't know. I put each word I didn't know at increasing indices in a huge list I created in memory. When I reached the end of the dictionary, I started from the beginning and did the same thing until I reached the page I started at.

Now I have a list of words that are mostly alphabetical, except they start somewhere in the middle of the alphabet, reach the end, and then start from the beginning of the alphabet. In other words, this is an alphabetically ordered list that has been "rotated." For example:

  words = [
    'ptolemaic',
    'retrograde',
    'supplant',
    'undulate',
    'xenoepist',
    'asymptote',  # <-- rotates here!
    'babka',
    'banoffee',
    'engender',
    'karpatka',
    'othellolagkage',
]

Write a function for finding the index of the "rotation point," which is where I started working from the beginning of the dictionary. This list is huge (there are lots of words I don't know) so we want to be efficient here.
*/

func isGreaterThan(s1 string, s2 string) bool {
	return s1 > s2
}

func isLessThan(s1 string, s2 string) bool {
	return s1 < s2
}

func main() {
	input1 := []string{"cape", "cake"}

	fmt.Printf("input:%v,output:%v\n", input1, rotationIndex(input1))

	input2 := []string{"ptolemaic", "retrograde", "supplant",
                                      "undulate", "xenoepist", "asymptote",
                                      "babka", "banoffee", "engender",
                                      "karpatka", "othellolagkage"}
	fmt.Printf("input:%v, output:%v\n", input2, rotationIndex(input2))
	input3 := []string{"grape", "orange", "plum", "radish", "apple"}
	fmt.Printf("input:%v, output:%v\n", input3, rotationIndex(input3))
}

func rotationIndex(words []string) int{
        var leftPart, rightPart bool
	length := len(words)
	leftPart = false
	rightPart  = false


	half := length / 2

	for half > 0 && half <= length - 1 {
		if half - 1 < length -1 {
			leftPart =  isGreaterThan(words[half - 1], words[half])
		}

		if  half + 1 >  length -1 && leftPart {
			return half
		}

                rightPart  = isLessThan(words[half], words[half+1])

		if leftPart && rightPart && isLessThan(words[half], words[length - 1]) {
			return half
		}

		if leftPart == false && rightPart && isLessThan(words[half], words[length - 1]) {
			length = half
			half = length / 2
		} else {
			half = (length + half) / 2
		}

	}

	return -1
}
