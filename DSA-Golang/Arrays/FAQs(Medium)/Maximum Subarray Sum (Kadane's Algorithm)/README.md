# Maximum Subarray Sum

## Problem

Given an integer array `nums`, return the maximum possible sum of a contiguous non-empty subarray.

### Example

**Input**
```
nums = [-2,1,-3,4,-1,2,1,-5,4]
```

**Output**
```
6
```

---

## Intuition

A negative running sum reduces the sum of any future subarray.

Whenever the current sum becomes negative, it is better to start a new subarray from the next element.

Kadane's Algorithm efficiently keeps track of the maximum subarray sum in a single traversal.

---

## Approach

1. Initialize:
   - `currentSum = 0`
   - `maxSum = nums[0]`
2. Traverse the array.
3. Add the current element to `currentSum`.
4. Update `maxSum` if needed.
5. If `currentSum` becomes negative, reset it to `0`.
6. Return `maxSum`.

---

## Algorithm

1. Initialize `currentSum` and `maxSum`.
2. Traverse the array once.
3. Update the running sum.
4. Update the maximum sum.
5. Reset the running sum if it becomes negative.
6. Return the maximum sum.

---

## Complexity

- **Time Complexity:** `O(N)`
- **Space Complexity:** `O(1)`

---

## Solution

See `main.go`.