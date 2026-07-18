# Armstrong Number

## Problem

Given an integer `n`, determine whether it is an Armstrong number.

An Armstrong number is equal to the sum of its digits, where each digit is raised to the power of the total number of digits.

### Example

**Input**
```
153
```

**Output**
```
true
```

---

## Approach

Count the number of digits. Then traverse each digit, raise it to the power of the digit count, and accumulate the result. Finally, compare the sum with the original number.

---

## Algorithm

1. Store the original number.
2. Count the total number of digits.
3. Extract each digit.
4. Add `digit^count` to the sum.
5. Compare the sum with the original number.

---

## Complexity

- **Time Complexity:** `O(log₁₀N)`
- **Space Complexity:** `O(1)`

---

## Solution

See `main.go`.