# Divisors of a Number

## Problem

Given an integer `n`, return all the divisors of `n` in ascending order.

A divisor is a number that divides `n` exactly without leaving a remainder.

### Example

**Input**
```
12
```

**Output**
```
1 2 3 4 6 12
```

---

## Approach

A brute-force solution checks every number from `1` to `n`, resulting in **O(N)** time.

A more efficient approach is to iterate only up to `√n`.

Divisors always occur in pairs. If `i` divides `n`, then `n/i` is also a divisor. We add both values whenever we find a divisor. Since the larger divisors are discovered before sorting, we sort the final list to return the divisors in ascending order.

---

## Algorithm

1. Create an empty list.
2. Iterate from `1` to `√n`.
3. If `i` divides `n`:
   - Add `i`.
   - If `i != n/i`, add `n/i`.
4. Sort the list.
5. Return the sorted list.

---

## Complexity

- **Time Complexity:** `O(√N + k log k)`
- **Space Complexity:** `O(k)`

where `k` is the number of divisors.

---

## Solution

See `main.go`.