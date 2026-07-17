# Palindrome Number

## Problem

Given an integer `n`, determine whether it is a palindrome.

A palindrome number remains the same when its digits are reversed.

### Example

**Input**
```
121
```

**Output**
```
true
```

---

## Approach

Store the original number and reverse it digit by digit. If the reversed number is equal to the original number, then the number is a palindrome.

---

## Algorithm

1. Store the original number.
2. Reverse the digits of the number.
3. Compare the reversed number with the original.
4. Return `true` if they are equal; otherwise, return `false`.

---

## Complexity

- **Time Complexity:** `O(log₁₀N)`
- **Space Complexity:** `O(1)`

---

## Solution

See `main.go`.