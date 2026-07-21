# Check if a Number is a Harshad Number

## Problem

Given an integer `n`, determine whether it is a Harshad Number.

A Harshad Number (also known as a Niven Number) is a number that is divisible by the sum of its digits.

### Example

**Input**
```
18
```

**Output**
```
true
```

**Explanation**

```
Sum of digits = 1 + 8 = 9

18 % 9 = 0
```

---

## Approach

To determine whether a number is a Harshad Number, first calculate the sum of its digits by extracting each digit using the modulo operator.

Once the digit sum is computed, check whether the original number is divisible by this sum.

If the remainder is `0`, the number is a Harshad Number; otherwise, it is not.

This approach processes each digit only once and requires constant extra space.

---

## Algorithm

1. Store the original number.
2. Initialize `sum = 0`.
3. Extract each digit and add it to `sum`.
4. Check if `original % sum == 0`.
5. Return the result.

---

## Complexity

- **Time Complexity:** `O(log N)`
- **Space Complexity:** `O(1)`

---

## Solution

See `main.go`.