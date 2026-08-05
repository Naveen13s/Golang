# 4Sum

## Problem

Given an integer array `nums` and an integer `target`, return all unique quadruplets whose sum equals `target`.

All four indices must be distinct, and duplicate quadruplets are not allowed.

### Example

**Input**
```
nums = [1,0,-1,0,-2,2]
target = 0
```

**Output**
```
[
 [-2,-1,1,2],
 [-2,0,0,2],
 [-1,0,0,1]
]
```

---

## Intuition

4Sum is an extension of the 3Sum problem.

Instead of checking every possible combination of four elements, first sort the array.

Then fix the first two elements using two loops and use the Two Pointer technique to find the remaining two elements.

Sorting also allows duplicate values to be skipped efficiently.

---

## Approach

1. Sort the array.
2. Fix the first element using `i`.
3. Fix the second element using `j`.
4. Set two pointers:
   - `left = j + 1`
   - `right = n - 1`
5. Calculate the sum of the four elements.
6. If the sum equals `target`, store the quadruplet.
7. If the sum is smaller than `target`, move `left`.
8. If the sum is greater than `target`, move `right`.
9. Skip duplicate values to avoid duplicate quadruplets.

---

## Algorithm

1. Sort `nums`.
2. Iterate `i` from `0` to `n-4`.
3. Skip duplicate values of `nums[i]`.
4. Iterate `j` from `i+1` to `n-3`.
5. Skip duplicate values of `nums[j]`.
6. Initialize `left = j+1` and `right = n-1`.
7. Use the two pointers to find the required pair.
8. Skip duplicates after finding a valid quadruplet.
9. Return all unique quadruplets.

---

## Complexity

- **Time Complexity:** `O(N³)`
- **Auxiliary Space Complexity:** `O(1)` excluding the output array and sorting implementation.

---

## Solution

See `main.go`.