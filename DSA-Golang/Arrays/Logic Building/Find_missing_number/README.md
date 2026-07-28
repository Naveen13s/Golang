# Missing Number

## Problem

Given an integer array `nums` of size `n` containing distinct numbers in the range `0` to `n` (inclusive), return the only missing number from the range.

### Example

**Input**
```
nums = [3, 0, 1]
```

**Output**
```
2
```

---

## Intuition

The numbers should contain every value from `0` to `n`.

If we calculate the expected sum of all numbers in this range and subtract the actual sum of the array elements, the remaining value is the missing number.

This approach avoids sorting and uses constant extra space.

---

## Approach

1. Calculate the expected sum using:

```
n × (n + 1) / 2
```

2. Calculate the actual sum of the array.
3. Return:

```
Expected Sum − Actual Sum
```

---

## Algorithm

1. Find the length of the array.
2. Compute the expected sum.
3. Traverse the array and compute the actual sum.
4. Return the difference.

---

## Complexity

### Sum Formula

- **Time Complexity:** `O(N)`
- **Space Complexity:** `O(1)`

### XOR Approach

- **Time Complexity:** `O(N)`
- **Space Complexity:** `O(1)`

---

## Solution

See `main.go`.