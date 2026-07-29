# Intersection of Two Sorted Arrays

## Problem

Given two sorted arrays `nums1` and `nums2`, return an array containing the intersection of these arrays.

Each element should appear as many times as it appears in both arrays. If an element appears `x` times in `nums1` and `y` times in `nums2`, it should appear `min(x, y)` times in the result.

### Example

**Input**
```
nums1 = [1, 2, 2, 3, 4]
nums2 = [2, 2, 4, 6]
```

**Output**
```
[2, 2, 4]
```

---

## Intuition

Since both arrays are sorted, we can efficiently compare elements using two pointers.

If the current elements are equal, they belong to the intersection. Otherwise, move the pointer pointing to the smaller element.

This naturally handles duplicate values by adding an element only when both pointers point to the same value.

---

## Approach

1. Initialize two pointers, one for each array.
2. Compare the current elements.
3. If both elements are equal:
   - Add the element to the result.
   - Move both pointers.
4. If the first element is smaller, move the first pointer.
5. Otherwise, move the second pointer.
6. Continue until one array is exhausted.

---

## Algorithm

1. Set `i = 0` and `j = 0`.
2. While both pointers are within bounds:
   - If `nums1[i] == nums2[j]`, add the element to the result and increment both pointers.
   - If `nums1[i] < nums2[j]`, increment `i`.
   - Otherwise, increment `j`.
3. Return the result.

---

## Complexity

- **Time Complexity:** `O(N + M)`
- **Space Complexity:** `O(min(N, M))`

---

## Solution

See `main.go`.