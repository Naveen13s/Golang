# Count All Digits

## Problem

Given an integer N, count the total number of digits.

### Example

Input

```
6678
```

Output

```
4
```

---

## Approach

Divide the number by 10 repeatedly until it becomes 0.
Each division removes one digit, and we increment a counter.

---

## Algorithm

1. Initialize `count = 0`.
2. While `n > 0`:
   - Increment `count`
   - Divide `n` by `10`
3. Return `count`.

---

## Complexity

- Time Complexity: **O(log₁₀N)**
- Space Complexity: **O(1)**

---

## Go Solution

See `main.go`.
