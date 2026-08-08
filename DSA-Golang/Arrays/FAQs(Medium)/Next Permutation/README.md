# Next Permutation

## Problem

Given an integer array `nums`, rearrange it into the next lexicographically greater permutation.

If no greater permutation exists, rearrange it into the smallest possible order (ascending order).

The rearrangement must be performed in-place using constant extra space.

### Example

**Input**
```
nums = [1,2,3]
```

**Output**
```
[1,3,2]
```

---

## Intuition

To obtain the next permutation, we need to make the smallest possible increase to the current arrangement.

We first find the pivot where the increasing order breaks from the right.

Then, we swap it with the next greater element and reverse the suffix to obtain the smallest lexicographically larger permutation.

---

## Approach

1. Find the first decreasing element from the right (pivot).
2. Find the smallest element greater than the pivot.
3. Swap the two elements.
4. Reverse the suffix after the pivot.
5. If no pivot exists, reverse the entire array.

---

## Algorithm

1. Traverse from right to left to find the pivot.
2. If a pivot exists:
   - Find the successor from the right.
   - Swap them.
3. Reverse all elements after the pivot.
4. Return the modified array.

---

## Complexity

- **Time Complexity:** `O(N)`
- **Space Complexity:** `O(1)`

---

## Solution

See `main.go`.