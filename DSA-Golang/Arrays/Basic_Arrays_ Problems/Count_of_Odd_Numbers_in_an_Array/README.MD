# Count of Odd Numbers in an Array

## Problem

Given an array of `n` elements, return the count of odd numbers present in the array.

### Example

**Input**
```
arr = [1, 2, 3, 4, 5]
```

**Output**
```
3
```

**Explanation**

Odd numbers in the array are `1`, `3`, and `5`.

---

## Intuition

To determine the number of odd elements, we examine each array element one by one.

A number is considered odd if it leaves a remainder of `1` when divided by `2`. By checking every element and maintaining a counter, we can efficiently determine the total number of odd numbers in the array.

Since each element is processed exactly once, this approach is both simple and efficient.

---

## Approach

Initialize a counter with `0`. Traverse the array and check each element using the modulo operator (`%`).

If an element is odd (`element % 2 != 0`), increment the counter.

After traversing the entire array, return the value of the counter.

---

## Algorithm

1. Initialize `count = 0`.
2. Traverse the array.
3. If the current element is odd:
   - Increment `count`.
4. Return `count`.

---

## Complexity

- **Time Complexity:** `O(N)`
- **Space Complexity:** `O(1)`

---

## Solution

See `main.go`.