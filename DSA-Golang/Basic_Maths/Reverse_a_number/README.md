# Reverse a Number

## Problem

Given an integer `n`, return the integer formed by reversing its digits.

### Example

**Input**
```
1234
```

**Output**
```
4321
```

## Approach

Extract the last digit of the number using the modulo operator (`% 10`), append it to the reversed number, and remove the last digit using integer division (`/ 10`). Repeat until the number becomes `0`.

## Algorithm

1. Initialize `reverse = 0`.
2. While `n > 0`:
   - Extract the last digit.
   - Append it to `reverse`.
   - Remove the last digit from `n`.
3. Return `reverse`.

## Complexity

- **Time Complexity:** `O(log₁₀ n)`
- **Space Complexity:** `O(1)`

## Solution

See `main.go`.
