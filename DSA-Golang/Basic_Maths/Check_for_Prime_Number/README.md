# Check for Prime Number

## Problem

Given an integer `n`, determine whether it is a prime number.

A prime number has exactly two positive divisors: `1` and itself.

### Example

**Input**
```
17
```

**Output**
```
true
```

---

## Approach

A prime number is divisible only by `1` and itself.

Instead of checking every number from `2` to `n-1`, we iterate only up to `√n`. If a divisor exists greater than `√n`, its corresponding divisor will always be less than or equal to `√n`. Therefore, checking up to the square root is sufficient.

If any number in this range divides `n`, the number is not prime. Otherwise, it is prime.

---

## Algorithm

1. If `n <= 1`, return `false`.
2. Iterate from `2` to `√n`.
3. If `n` is divisible by any number, return `false`.
4. Otherwise, return `true`.

---

## Complexity

- **Time Complexity:** `O(√N)`
- **Space Complexity:** `O(1)`

---

## Solution

See `main.go`.