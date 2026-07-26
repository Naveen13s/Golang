# Left Rotate Array by One

## Problem

Given an integer array `nums`, rotate the array to the left by one position.

There is no need to return anything. Modify the given array in-place.

### Example

**Input**
```
nums = [1, 2, 3, 4, 5]
```

**Output**
```
[2, 3, 4, 5, 1]
```

---

## Intuition

In a left rotation by one position, every element moves one position to the left, while the first element moves to the last position.

Before shifting the elements, we need to save the first element because its value will be overwritten.

After storing the first element, shift every remaining element one position to the left. Finally, place the saved first element at the end of the array.

This allows us to perform the rotation in-place without creating another array.

---

## Approach

1. Store the first element in a temporary variable.
2. Shift every element from index `1` to `n-1` one position to the left.
3. Place the stored first element at index `n-1`.
4. The original array is now left-rotated by one position.

---

## Algorithm

1. Find the length of the array.
2. If the array contains zero or one element, no rotation is required.
3. Store `nums[0]` in `temp`.
4. Iterate from index `1` to `n-1`.
5. Set `nums[i-1] = nums[i]`.
6. Set `nums[n-1] = temp`.

---

## Complexity

- **Time Complexity:** `O(N)`
- **Space Complexity:** `O(1)`

---

## Solution

See `main.go`.