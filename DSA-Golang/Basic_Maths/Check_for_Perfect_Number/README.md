# Perfect Number

## Problem

Given an integer `n`, determine whether it is a perfect number.

A perfect number is a number whose proper divisors (excluding the number itself) add up to the number itself.

### Example

**Input**
```
28
```

**Output**
```
true
```

---

## Approach

A proper divisor is a positive divisor of a number excluding the number itself.

Instead of checking every number from `1` to `n-1`, we only iterate up to the square root of `n`.

Whenever we find a divisor `i` such that `n % i == 0`, we also get another divisor `n / i`. We add both divisors to the sum. For perfect squares, we ensure the square root is counted only once.

Finally, if the sum of all proper divisors equals the original number, then the number is a perfect number.

---

## Algorithm

1. If `n <= 1`, return `false`.
2. Initialize `sum = 1` because `1` is a proper divisor of every number greater than `1`.
3. Iterate from `2` to `√n`.
4. If `i` divides `n`:
   - Add `i` to the sum.
   - If `i != n / i`, add `n / i` as well.
5. Compare the sum with `n`.
6. Return the result.

---

## Complexity

- **Time Complexity:** `O(√N)`
- **Space Complexity:** `O(1)`

---

## Solution

See `main.go`.