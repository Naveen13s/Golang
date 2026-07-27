# Remove Duplicates from Sorted Array

## Problem

Given an integer array `nums` sorted in non-decreasing order, remove all duplicates in-place so that each unique element appears only once.

Return the number of unique elements in the array.

If the number of unique elements is `k`, the first `k` elements of `nums` should contain the unique values in their original order.

The elements after the first `k` positions do not matter.

### Example

**Input**
```
nums = [1, 1, 2, 2, 2, 3, 4, 4]
```

**Output**
```
k = 4
```

The first `k` elements become:

```
[1, 2, 3, 4]
```

---

## Intuition

Since the array is already sorted, duplicate values will always appear next to each other.

This allows us to use two pointers.

One pointer keeps track of the position of the last unique element, while another pointer traverses the array looking for new unique values.

Whenever the current element is different from the last unique element, we have found a new unique value. We move the unique-element pointer forward and place the new value at that position.

After processing the entire array, all unique values are stored at the beginning of the array.

---

## Approach

Use two pointers:

- `i` points to the last unique element.
- `j` traverses the array from left to right.

For each `j`:

1. Compare `nums[j]` with `nums[i]`.
2. If they are different, a new unique element has been found.
3. Increment `i`.
4. Store `nums[j]` at `nums[i]`.
5. Continue until the end of the array.
6. Return `i + 1`, which represents the number of unique elements.

---

## Algorithm

1. If the array is empty, return `0`.
2. Initialize `i = 0`.
3. Traverse the array using `j` from index `1`.
4. If `nums[j] != nums[i]`:
   - Increment `i`.
   - Set `nums[i] = nums[j]`.
5. Return `i + 1`.

---

## Complexity

- **Time Complexity:** `O(N)`
- **Space Complexity:** `O(1)`

The solution modifies the original array in-place.

---

## Solution

See `main.go`.