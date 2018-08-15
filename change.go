/*
   calculate the number of coins (and which coins) are needed to make change for a given amount
*/

package change

import "math"

/*
   iterate through the coins available, using previously calculated minumum numbers to calculate the next number.  uses a package variable.
*/
func CalculateChange(amount int, coins []int) (int, map[int]int) {
	var minCoins = make(map[int]int)
	var coinsUsed = make(map[int]map[int]int)
	if numCoins, ok := minCoins[amount]; ok {
		return numCoins, coinsUsed[amount]
	}
	minCoins[0] = 0
	coinsUsed[0] = make(map[int]int)
	for amt := 1; amt <= amount; amt++ {
		minCoins[amt] = math.MaxUint32
		for i := range coins {
			coin := coins[i]
			if amt >= coin {
				if 1+minCoins[amt-coin] < minCoins[amt] {
					minCoins[amt] = 1 + minCoins[amt-coin]
					coinsUsed[amt] = make(map[int]int)
					for k, v := range coinsUsed[amt-coin] {
						coinsUsed[amt][k] = v
					}
					coinsUsed[amt][coin]++
				}
			}
		}
	}
	return minCoins[amount], coinsUsed[amount]
}

/*
   accept requests on channels, send out responses based on package state.  run inside a goroutine. example:

   var amount chan int = make(chan int)

   var coins chan []int = make(chan []int)

   var numCoins chan int = make(chan int)

   var coinsUsed chan map[int]int = make(chan map[int]int)

   go change.PersistentChangeCalculator(amount, coins, numCoins, coinsUsed)

   amount <- 50

   coins <- []int{25, 10, 5, 1}

   fmt.Println(<-numCoins, <-coinsUsed)

   XXX: need to re-factor minCoins and coinsUsed to store state (using a package variable confuses CalculateChange if different coins are used across runs.  allow as input, or store coins used as part of state storage (eg: build a struct.  this would allow for a DDoS on memory footprint if someone sent a lot of coin types, so making the coin set finite by having named coin sets would seem most logical.
*/
func PersistentChangeCalculator(amount <-chan int, coins <-chan []int, numCoins chan<- int, coinsUsed chan<- map[int]int) {
	n, c := CalculateChange(<-amount, <-coins)
	numCoins <- n
	coinsUsed <- c
}
