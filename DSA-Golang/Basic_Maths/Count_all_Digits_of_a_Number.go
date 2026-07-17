//Count all Digits of a Number
/* You are given an integer n. You need to return the number of digits in the number.
   The number will have no leading zeroes, except when the number is 0 itself.
*/

package main

import "fmt"

func countDigit(n int) int {
	if n == 0 {
		return 1
	}

	count := 0

	for n > 0 {
		count++
		n /= 10
	}

	return count
}

func main() {
	n := 6678

	ans := countDigit(n)

	fmt.Println("Number:", n)
	fmt.Println("Digit Count:", ans)
}

/* // 2nd approach the optimal O(1) solution using math.Log10

import "math"

func countDigit(n int) int {
    if n == 0 {
        return 1
    }

    return int(math.Log10(float64(n))) + 1
} */

/*Time Complexity

The loop runs once for each digit.

If the number has d digits:

Time Complexity: O(d) (or O(log₁₀ n))
Space Complexity: O(1)
*/
