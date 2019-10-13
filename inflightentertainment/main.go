package main

import "fmt"

/*
You've built an inflight entertainment system with on-demand movie streaming.

Users on longer flights like to start a second movie right when their first one ends, but they complain that the plane usually lands before they can see the ending. So you're building a feature for choosing two movies whose total runtimes will equal the exact flight length.

Write a function that takes an integer flight_length (in minutes) and a list of integers movie_lengths (in minutes) and returns a boolean indicating whether there are two numbers in movie_lengths whose sum equals flight_length.

When building your function:

Assume your users will watch exactly two movies
Don't make your users watch the same movie twice
Optimize for runtime over memory
Do you have an answer?
*/


func main() {
	movies := []int{2, 4}
	flightLength := 1

	fmt.Printf("%v\n", canTwoMoviesFillFlight(movies, flightLength))

        movies1 := []int{2, 4}
	flightLength1 := 6

	fmt.Printf("%v\n", canTwoMoviesFillFlight(movies1, flightLength1))

        movies2 := []int{3, 8}
	flightLength2 := 6

	fmt.Printf("%v\n", canTwoMoviesFillFlight(movies2, flightLength2))

        movies3 := []int{3, 8, 3}
	flightLength3 := 6

	fmt.Printf("%v\n", canTwoMoviesFillFlight(movies3, flightLength3))

        movies4 := []int{1, 2, 3, 4, 5, 6} 
	flightLength4 := 7

        fmt.Printf("%v\n", canTwoMoviesFillFlight(movies4, flightLength4))

        movies5 := []int{4,3,2} 
	flightLength5 := 5

        fmt.Printf("%v\n", canTwoMoviesFillFlight(movies5, flightLength5))

        movies6 := []int{6} 
	flightLength6 := 6

        fmt.Printf("%v\n", canTwoMoviesFillFlight(movies6, flightLength6))

        movies7 := []int{} 
	flightLength7 := 2

        fmt.Printf("%v\n", canTwoMoviesFillFlight(movies7, flightLength7))

}

func canTwoMoviesFillFlight(movies []int, flightLength int) bool {
	processed := make(map[int]bool)

	if len(movies) > 0 {
                processed[movies[0]]= true
        }

	var complement int

	for i := 1; i < len(movies); i++ {
		complement = flightLength - movies[i]
		if processed[complement] {
			return true
		}
	}

	return false
}
