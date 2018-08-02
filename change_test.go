package change

import "reflect"
import "testing"

func TestChange(t *testing.T) {
    testCase(t, 1, 67, []int{25, 10, 5, 1}, 6, map[int]int{25: 2, 10: 1, 5: 1, 1: 2})
    testCase(t, 2, 67, []int{30, 5, 1}, 5, map[int]int{30: 2, 5: 1, 1: 2})
    testCase(t, 3, 40, []int{25, 20, 10, 1}, 2, map[int]int{20: 2})
}

// a test case for the change code
// testCase (each case has a test case for debugging
// amount to test
// the coins to use
// the total number of coins returned by the function
// the coins which should be returned by the function
func testCase(t *testing.T, testCase int, amount int, coins []int, totalCoins int, validCoins map[int]int) {
    var numCoins int
    var coinsUsed map[int]int
    numCoins, coinsUsed = calculateChange(amount, coins)

    if numCoins != totalCoins {
        t.Error("number of coins wrong for test case", testCase)
    } else if coinsDontMatch(coinsUsed, validCoins) {
        t.Error("coins don't match for test case", testCase)
    }
}

func coinsDontMatch(calculated map[int]int, valid map[int]int) bool {
    return !reflect.DeepEqual(calculated, valid)
}
