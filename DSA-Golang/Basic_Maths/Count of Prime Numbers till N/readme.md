# Count of Prime Numbers till N

## Problem

Given an integer `n`, find the number of prime numbers in the range **[1, n]** (inclusive).

Return the total number of prime numbers.

### Example

**Input**
```
10
```

**Output**
```
4
```

**Explanation**

Prime numbers are:

```
2, 3, 5, 7
```

---

## Approach

A brute-force solution checks whether every number from `2` to `n` is prime, resulting in a time complexity of **O(N√N)**.

A more efficient approach is the **Sieve of Eratosthenes**.

We initially assume that every number is prime. Starting from `2`, we mark all multiples of each prime number as non-prime. After processing all numbers up to `√n`, the remaining unmarked numbers are prime. Finally, we count all prime numbers in the range.

This algorithm avoids repeated divisibility checks and is the standard approach for finding all prime numbers up to `n`.

---

## Algorithm

1. Create a boolean array of size `n + 1`.
2. Mark all numbers as prime initially.
3. Mark `0` and `1` as non-prime.
4. Iterate from `2` to `√n`.
5. If the current number is prime, mark all its multiples as non-prime.
6. Count all remaining prime numbers.
7. Return the count.

---

## Complexity

- **Time Complexity:** `O(N log log N)`
- **Space Complexity:** `O(N)`

---

## Solution

See `main.go`.