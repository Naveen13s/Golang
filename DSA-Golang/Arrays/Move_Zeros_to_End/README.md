# Move Zeros to End

## Problem

Given an integer array `nums`, move all the `0`s to the end of the array.

The relative order of the non-zero elements must remain the same.

The operation must be performed in-place without making a copy of the array.

### Example

**Input**
```
nums = [0, 1, 0, 3, 12]
```

**Output**
```
[1, 3, 12, 0, 0]
```

---

## Intuition

We need to move all non-zero elements toward the beginning of the array while maintaining their original relative order.

To achieve this efficiently, we use two pointers.

One pointer traverses the entire array, while another pointer keeps track of the position where the next non-zero element should be placed.

Whenever we encounter a non-zero element, we swap it with the element at the second pointer and move that pointer forward.

As a result, all non-zero elements are moved to the front in their original order, while zeros are automatically pushed toward the end.

---

## Approach

Use two pointers:

- `i` traverses the array.
- `j` represents the position where the next non-zero element should be placed.

For every element:

1. If `nums[i]` is non-zero, swap `nums[i]` with `nums[j]`.
2. Increment `j`.
3. If `nums[i]` is zero, continue searching.

After the traversal, all non-zero elements will be at the beginning and all zeros will be at the end.

---

## Algorithm

1. Initialize `j = 0`.
2. Traverse the array using `i`.
3. If `nums[i] != 0`:
   - Swap `nums[i]` and `nums[j]`.
   - Increment `j`.
4. Continue until the end of the array.
5. The array is modified in-place.

---

## Complexity

- **Time Complexity:** `O(N)`
- **Space Complexity:** `O(1)`

---

## Solution

See `main.go`.