# Greatest Common Divisor (GCD)

## Problem

Given two integers `n1` and `n2`, find their Greatest Common Divisor (GCD).

The GCD is the largest positive integer that divides both numbers without leaving a remainder.

### Example

**Input**
```
12 18
```

**Output**
```
6
```

---

## Approach

Instead of checking every common divisor, we use **Euclid's Algorithm**, which is much more efficient.

The algorithm is based on the property:

- `GCD(a, b) = GCD(b, a % b)`

We repeatedly replace the larger number with the remainder of dividing the larger number by the smaller one. This process continues until the remainder becomes `0`.

The last non-zero value is the Greatest Common Divisor.

This approach is significantly faster than checking all possible divisors.

---

## Algorithm

1. While `n2` is not `0`:
   - Store `n2` in a temporary variable.
   - Update `n2 = n1 % n2`.
   - Update `n1 = temp`.
2. Return `n1`.

---

## Complexity

- **Time Complexity:** `O(log(min(n1, n2)))`
- **Space Complexity:** `O(1)`

---

## Solution

See `main.go`.