# Factorial of a Given Number

## Problem

Given an integer `n`, return the factorial of the number.

The factorial of a non-negative integer `n` is the product of all positive integers less than or equal to `n`.

### Example

**Input**
```
5
```

**Output**
```
120
```

---

## Approach

Initialize the result as `1` and multiply it by every integer from `2` to `n`.

---

## Algorithm

1. Initialize `result = 1`.
2. Iterate from `2` to `n`.
3. Multiply `result` by the current number.
4. Return `result`.

---

## Complexity

- **Time Complexity:** `O(n)`
- **Space Complexity:** `O(1)`

---

## Solution

See `main.go`.