# Reverse an Array

## Problem

Given an array `arr` of `n` elements, reverse the array in-place.

The reversal should be performed without using any additional array.

### Example

**Input**
```
arr = [1, 2, 3, 4, 5]
```

**Output**
```
[5, 4, 3, 2, 1]
```

---

## Intuition

To reverse an array, the first element should exchange its position with the last element, the second element with the second-last element, and so on.

This can be efficiently achieved using two pointers. One pointer starts from the beginning of the array, while the other starts from the end. By swapping the elements at these positions and moving the pointers toward the center, the entire array is reversed without using extra memory.

---

## Approach

Initialize two pointers:

- `left = 0`
- `right = n - 1`

While `left < right`:

- Swap `arr[left]` and `arr[right]`.
- Move `left` one position to the right.
- Move `right` one position to the left.

Continue until both pointers meet or cross.

---

## Algorithm

1. Initialize `left = 0` and `right = n - 1`.
2. While `left < right`:
   - Swap `arr[left]` and `arr[right]`.
   - Increment `left`.
   - Decrement `right`.
3. Return the reversed array.

---

## Complexity

- **Time Complexity:** `O(N)`
- **Space Complexity:** `O(1)`

---

## Solution

See `main.go`.