# Largest Digit in a Number

## Problem

Given an integer `n`, return the largest digit present in the number.

### Example

**Input**
```
5832
```

**Output**
```
8
```

---

## Approach

Extract each digit using the modulo operator (`%`) and compare it with the current maximum digit. Update the maximum whenever a larger digit is found.

---

## Algorithm

1. Initialize `maxDigit` to `0`.
2. Traverse each digit of the number.
3. Compare the current digit with `maxDigit`.
4. Update `maxDigit` if necessary.
5. Return `maxDigit`.

---

## Complexity

- **Time Complexity:** `O(log₁₀N)`
- **Space Complexity:** `O(1)`

---

## Solution

See `main.go`.