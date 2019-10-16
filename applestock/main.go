package main

import (
	"fmt"
	"math"
)
/*
First, I wanna know how much money I could have made yesterday if I'd been trading Apple stocks all day.

So I grabbed Apple's stock prices from yesterday and put them in a list called stock_prices, where:

The indices are the time (in minutes) past trade opening time, which was 9:30am local time.
The values are the price (in US dollars) of one share of Apple stock at that time.
So if the stock cost $500 at 10:30am, that means stock_prices[60] = 500.

Write an efficient function that takes stock_prices and returns the best profit I could have made from one purchase and one sale of one share of Apple stock yesterday.

For example:

  stock_prices = [10, 7, 5, 8, 11, 9]

get_max_profit(stock_prices)
# Returns 6 (buying for $5 and selling for $11)

No "shorting"—you need to buy before you can sell. Also, you can't buy and sell in the same time step—at least 1 minute has to pass.
*/


func getMaxProfit(stock []float64) float64{
	var maxProfitSoFar, profit, min float64
	//trackMaxProfit
	maxProfitSoFar = -100000.0

	if len(stock) <= 1 {
		return maxProfitSoFar
	}

	min = stock[0]

	for i := 1; i < len(stock); i++ {
		profit = (min * -1.0) + stock[i]
		maxProfitSoFar = math.Max(maxProfitSoFar, profit)
		//trackMinSeen
		min = math.Min(min, stock[i])
	}

	return maxProfitSoFar
}

func main() {
	stock1 := []float64{1.0, 5.0, 3.0, 2.0}

	fmt.Printf("%v\n", getMaxProfit(stock1))

	stock2 := []float64{7.0, 2.0, 8.0, 9.0}
        fmt.Printf("%v\n", getMaxProfit(stock2))

	stock3 := []float64{1.0, 6.0, 7.0, 9.0}
        fmt.Printf("%v\n", getMaxProfit(stock3))

	stock4 := []float64{9.0, 7.0, 4.0, 1.0}
        fmt.Printf("%v\n", getMaxProfit(stock4))

	stock5 := []float64{1.0, 1.0, 1.0, 1.0}
	fmt.Printf("%v\n", getMaxProfit(stock5))
}
