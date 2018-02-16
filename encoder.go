package stringsizer

import (
	"math"
)

var baseChars = "abcdefghjijklmnopqrstuvwxyzABCDEFGHJIJKLMNOPQRSTUVWXYZ0123456789,.;[]-=_+()^&!"
var base = float64(len(baseChars))

// Transform will take a number and convert it to a string that
// resembles its base formula.
// Consider the base character library "abc". The number 0 will just be "a".
// The number 1 will be "b", 2 will be "c", and 3 will be "ba".
// (The preceding "a" is not shown).
func transform(originalNumber int) (encoded string) {
	number := float64(originalNumber)
	if number < 1 {
		return string(baseChars[0])
	}

	nums := []int{}
	for {
		power := math.Floor(math.Log(number) / math.Log(base))
		if len(nums) == 0 {
			nums = make([]int, int(power)+1)
			for i := range nums {
				nums[i] = 0
			}
		}
		numAtPower := math.Floor(number / math.Pow(base, power))
		nums[int(power)] = int(numAtPower)
		number = number - numAtPower*math.Pow(base, power)
		if number < base {
			if number > 0 {
				nums[0] = int(number)
			}
			break
		}
	}

	// s := make([]string, len(nums))
	// for power, numAtPower := range nums {
	// 	s[power] = fmt.Sprintf("%d*%d^%d", numAtPower, int(base), power)
	// }
	// fmt.Println(originalNumber, strings.Join(s, " + "))

	for _, numAtPower := range nums {
		encoded = string(baseChars[numAtPower]) + encoded
	}
	return
}
