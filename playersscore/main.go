package main

import "fmt"

/*
You created a game that is more popular than Angry Birds.

Each round, players receive a score between 0 and 100, which you use to rank them from highest to lowest. So far you're using an algorithm that sorts in O(n\lg{n})O(nlgn) time, but players are complaining that their rankings aren't updated fast enough. You need a faster sorting algorithm.

Write a function that takes:

a list of unsorted_scores
the highest_possible_score in the game
and returns a sorted list of scores in less than O(n\lg{n})O(nlgn) time.

For example:

  unsorted_scores = [37, 89, 41, 65, 91, 53]
HIGHEST_POSSIBLE_SCORE = 100

# Returns [91, 89, 65, 53, 41, 37]
sort_scores(unsorted_scores, HIGHEST_POSSIBLE_SCORE)

We’re defining nn as the number of unsorted_scores because we’re expecting the number of players to keep climbing.

And, we'll treat highest_possible_score as a constant instead of factoring it into our big O time and space costs because the highest possible score isn’t going to change. Even if we do redesign the game a little, the scores will stay around the same order of magnitude.
*/


func sortScores(unsortedScores []int, HighestPossibleScore int) []int{
	var scoresTable [101]int
	var index int

	sortedScores := make([]int, 0, 100)

	for _, v := range unsortedScores {
		scoresTable[v] += 1
	}

	for  i := 100;  i >= 0; i-- {
		if scoresTable[i] > 0 {
			index = 0
			for index < scoresTable[i] {
				sortedScores = append(sortedScores, i)
				index += 1
			}
		}
	}

	return sortedScores
}

func main() {
	input1 := []int{20, 100, 99, 100, 5, 0}
	fmt.Printf("%v\n", sortScores(input1, 100))

        input2 := []int{}
	fmt.Printf("%v\n", sortScores(input2, 100))

        input3 := []int{55}
	fmt.Printf("%v\n", sortScores(input3, 100))

        input4 := []int{30, 60}
	fmt.Printf("%v\n", sortScores(input4, 100))

        input5 := []int{37, 89, 41, 65, 91, 53}
	fmt.Printf("%v\n", sortScores(input5, 100))

        input6 := []int{20, 10, 30, 30, 10, 20}
	fmt.Printf("%v\n", sortScores(input6, 100))

}
