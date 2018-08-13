/*
    calculate the number of coins (and which coins) are needed to make change for a given amount
*/

package change

import "math"

/*
    iterate through the coins available, using previously calculated minumum numbers to calculate the next number
*/
func CalculateChange(amount int, coins []int) (int, map[int]int) {
    var minCoins = make(map[int]int)
    minCoins[0] = 0
    var coinsUsed = make(map[int]map[int]int)
    coinsUsed[0] = make(map[int]int)

    for amt := 1; amt <= amount; amt++ {
        minCoins[amt] = math.MaxUint32
        for i := range coins {
            coin := coins[i]
            if amt >= coin {
                if 1 + minCoins[amt - coin] < minCoins[amt] {
                    minCoins[amt] = 1 + minCoins[amt - coin]
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
