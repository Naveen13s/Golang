# Second Largest Element in an Array

## Problem

Given an array of integers `nums`, return the second-largest distinct element in the array.

If the second-largest element does not exist, return `-1`.

### Example

**Input**
```
nums = [8, 8, 7, 6, 5]
```

**Output**
```
7
```

---

## Intuition

Instead of sorting the array, we can find the largest and second-largest elements in a single traversal.

We maintain two variables:

- `largest` stores the largest element found so far.
- `secondLargest` stores the second-largest distinct element found so far.

If the current element becomes the new largest, the previous largest becomes the second largest.

Otherwise, if the current element is smaller than the largest but greater than the current second largest, we update the second largest.

Duplicate occurrences of the largest element are ignored because we need the second-largest distinct element.

---

## Approach

1. Initialize `largest` with the first element.
2. Initialize `secondLargest` with `-1`.
3. Traverse the remaining elements.
4. If the current element is greater than `largest`:
   - Move `largest` to `secondLargest`.
   - Update `largest`.
5. Otherwise, if the current element is smaller than `largest` but greater than `secondLargest`, update `secondLargest`.
6. Return `secondLargest`.

---

## Algorithm

1. If the array has fewer than two elements, return `-1`.
2. Set `largest = nums[0]`.
3. Set `secondLargest = -1`.
4. Traverse the array.
5. Update `largest` and `secondLargest` when necessary.
6. Return `secondLargest`.

---

## Complexity

- **Time Complexity:** `O(N)`
- **Space Complexity:** `O(1)`

---

## Solution

See `main.go`.