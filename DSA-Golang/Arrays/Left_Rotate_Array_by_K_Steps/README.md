# Left Rotate Array by K Steps

## Problem

Given an integer array `nums` and a non-negative integer `k`, rotate the array to the left by `k` steps.

### Example

**Input**
```
nums = [1, 2, 3, 4, 5, 6, 7]
k = 3
```

**Output**
```
[4, 5, 6, 7, 1, 2, 3]
```

---

## Intuition

A simple approach would be to left-rotate the array by one position `k` times. However, this would require `O(N × K)` time.

We can perform the rotation more efficiently using the reversal algorithm.

For a left rotation by `k` positions:

1. Reverse the first `k` elements.
2. Reverse the remaining `n-k` elements.
3. Reverse the entire array.

These three reversals rearrange the elements into the required rotated order without using an additional array.

We also use `k % n` because rotating an array by its length brings it back to its original position.

---

## Approach

First, reduce `k` using:

```
k = k % n
```

Then:

1. Reverse the elements from index `0` to `k-1`.
2. Reverse the elements from index `k` to `n-1`.
3. Reverse the entire array from index `0` to `n-1`.

The array is modified in-place.

---

## Algorithm

1. Find the length `n` of the array.
2. If the array is empty, return.
3. Set `k = k % n`.
4. Reverse the first `k` elements.
5. Reverse the remaining `n-k` elements.
6. Reverse the entire array.

---

## Complexity

- **Time Complexity:** `O(N)`
- **Space Complexity:** `O(1)`

---

## Solution

See `main.go`.