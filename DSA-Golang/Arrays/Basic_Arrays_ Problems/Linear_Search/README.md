# Linear Search - Find the Smallest Index

## Problem

Given an array of integers `nums` and an integer `target`, find the smallest index (0-based indexing) where the target appears in the array.

If the target is not found, return `-1`.

### Example

**Input**
```
nums = [2, 3, 4, 3, 5]
target = 3
```

**Output**
```
1
```

**Explanation**

The target `3` appears at indices `1` and `3`. Since we need the smallest index, we return `1`.

---

## Intuition

Since the array is not guaranteed to be sorted, we can search for the target by traversing the array from left to right.

Because we visit indices in increasing order, the first occurrence of the target will always be its smallest index.

Therefore, as soon as we find the target, we can immediately return the current index.

If the target is never found, we return `-1`.

---

## Approach

Traverse the array starting from index `0`.

For every element:

- Compare it with `target`.
- If `nums[i] == target`, return `i` immediately.
- If the traversal finishes without finding the target, return `-1`.

This is known as **Linear Search**.

---

## Algorithm

1. Start traversing the array from index `0`.
2. Compare each element with `target`.
3. If the current element equals `target`, return its index.
4. If no element matches the target, return `-1`.

---

## Complexity

- **Time Complexity:** `O(N)`
- **Space Complexity:** `O(1)`

---

## Solution

See `main.go`.