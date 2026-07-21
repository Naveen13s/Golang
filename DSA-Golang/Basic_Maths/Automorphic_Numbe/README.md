# Check if a Number is an Automorphic Number

## Problem

Given an integer `n`, determine whether it is an Automorphic Number.

An Automorphic Number is a number whose square ends with the number itself.

### Example

**Input**
```
25
```

**Output**
```
true
```

**Explanation**

```
25² = 625

625 ends with 25.
```

---

## Approach

An automorphic number appears at the end of its square.

First, compute the square of the given number. Then count the number of digits in the original number. Using this digit count, calculate `10^digits` to extract the last digits of the square using the modulo operator.

If the extracted value matches the original number, then the number is automorphic.

This arithmetic approach avoids converting numbers into strings.

---

## Algorithm

1. Compute the square of the number.
2. Count the number of digits.
3. Compute `10^digits`.
4. Extract the last digits of the square using modulo.
5. Compare the extracted value with the original number.
6. Return the result.

---

## Complexity

- **Time Complexity:** `O(log N)`
- **Space Complexity:** `O(1)`

---

## Solution

See `main.go`.