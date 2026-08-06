# Sort an Array of 0s, 1s and 2s

## Problem

Given an array `nums` containing only `0`, `1`, and `2`, sort the array in non-decreasing order.

The sorting must be performed in-place without using another array.

### Example

**Input**
```
nums = [2,0,2,1,1,0]
```

**Output**
```
[0,0,1,1,2,2]
```

---

## Intuition

Since the array contains only three distinct values, a comparison-based sorting algorithm is unnecessary.

The Dutch National Flag Algorithm partitions the array into three regions:

- `0`s at the beginning
- `1`s in the middle
- `2`s at the end

Using three pointers, the array can be sorted in a single traversal.

---

## Approach

1. Initialize three pointers:
   - `low`
   - `mid`
   - `high`
2. Traverse while `mid <= high`.
3. If the current element is:
   - `0`: swap with `low`, increment both `low` and `mid`.
   - `1`: increment `mid`.
   - `2`: swap with `high`, decrement `high`.
4. Continue until all elements are processed.

---

## Algorithm

1. Set `low = 0`, `mid = 0`, `high = n - 1`.
2. Traverse the array while `mid <= high`.
3. Perform swaps based on the value at `nums[mid]`.
4. Return the sorted array.

---

## Complexity

- **Time Complexity:** `O(N)`
- **Space Complexity:** `O(1)`

---

## Solution

See `main.go`.