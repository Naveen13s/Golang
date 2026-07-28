# Union of Two Sorted Arrays

## Problem

Given two sorted arrays `nums1` and `nums2`, return an array containing the union of the two arrays.

The union should:

- Contain only distinct elements.
- Be sorted in ascending order.

### Example

**Input**
```
nums1 = [1, 2, 2, 3, 5]
nums2 = [2, 3, 4, 5, 6]
```

**Output**
```
[1, 2, 3, 4, 5, 6]
```

---

## Intuition

Since both arrays are already sorted, we can efficiently merge them using two pointers.

At every step:

- Compare the current elements.
- Add the smaller element if it hasn't already been added.
- If both elements are equal, add it once and move both pointers.

This preserves the sorted order while removing duplicates.

---

## Approach

1. Initialize two pointers for both arrays.
2. Compare the current elements.
3. Add the smaller element if it is not already present in the result.
4. If both elements are equal, add one copy and move both pointers.
5. Process any remaining elements from either array.

---

## Algorithm

1. Set `i = 0`, `j = 0`.
2. While both pointers are within bounds:
   - Compare `nums1[i]` and `nums2[j]`.
   - Add the appropriate unique element.
3. Add the remaining unique elements from the unfinished array.
4. Return the result.

---

## Complexity

- **Time Complexity:** `O(N + M)`
- **Space Complexity:** `O(N + M)`

---

## Solution

See `main.go`.