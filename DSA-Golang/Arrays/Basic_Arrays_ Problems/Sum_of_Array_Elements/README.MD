# Sum of All Elements in an Array

## Problem

Given an array `arr` of size `n`, find the sum of all the elements in the array.

Return the total sum.

### Example

**Input**
```
arr = [1, 2, 3, 4, 5]
```

**Output**
```
15
```

---

## Intuition

Every element in the array contributes exactly once to the final sum.

By traversing the array from the first element to the last and continuously adding each element to a running total, we can efficiently compute the sum. After processing all elements, the running total represents the sum of the entire array.

---

## Approach

Initialize a variable `sum` with `0`. Traverse the array using a loop and add each element to `sum`. Once all elements have been processed, return the value of `sum`.

This approach visits each element exactly once, making it efficient for arrays of any size.

---

## Algorithm

1. Initialize `sum = 0`.
2. Traverse the array.
3. Add each element to `sum`.
4. Return `sum`.

---

## Complexity

- **Time Complexity:** `O(N)`
- **Space Complexity:** `O(1)`

---

## Solution

See `main.go`.