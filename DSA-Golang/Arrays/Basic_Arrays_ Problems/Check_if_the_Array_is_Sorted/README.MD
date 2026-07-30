# Check if the Array is Sorted

## Problem

Given an array `arr` of size `n`, determine whether it is sorted in ascending (non-decreasing) order.

Return `true` if the array is sorted; otherwise, return `false`.

### Example

**Input**
```
arr = [1, 2, 2, 4, 5]
```

**Output**
```
true
```

---

## Intuition

A sorted array follows the property that every element is less than or equal to the element immediately after it.

By comparing each element with its next element, we can quickly detect whether the ordering is violated. If any element is greater than the next one, the array is not sorted. Otherwise, if all adjacent pairs satisfy the condition, the array is sorted.

This allows us to determine the result in a single traversal.

---

## Approach

Traverse the array from the first element to the second-last element.

For each index, compare the current element with the next element.

- If `arr[i] > arr[i + 1]`, return `false`.
- If no violation is found after checking all pairs, return `true`.

---

## Algorithm

1. Traverse the array from index `0` to `n - 2`.
2. Compare `arr[i]` with `arr[i + 1]`.
3. If `arr[i] > arr[i + 1]`, return `false`.
4. Otherwise, continue.
5. Return `true` after the loop.

---

## Complexity

- **Time Complexity:** `O(N)`
- **Space Complexity:** `O(1)`

---

## Solution

See `main.go`.