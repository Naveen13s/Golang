# Largest Element in an Array

## Problem

Given an array of integers `nums`, return the value of the largest element in the array.

### Example

**Input**
```
nums = [3, 8, 2, 10, 5]
```

**Output**
```
10
```

---

## Intuition

To find the largest element, we need to examine every element in the array.

We start by considering the first element as the largest. Then, we traverse the remaining elements one by one.

Whenever we encounter an element greater than the current largest value, we update the largest value.

After traversing the entire array, the stored value represents the largest element.

---

## Approach

1. Initialize `largest` with the first element of the array.
2. Traverse the array starting from the second element.
3. Compare each element with `largest`.
4. If the current element is greater, update `largest`.
5. Return `largest`.

---

## Algorithm

1. Set `largest = nums[0]`.
2. Iterate from index `1` to the end of the array.
3. If `nums[i] > largest`, set `largest = nums[i]`.
4. Return `largest`.

---

## Complexity

- **Time Complexity:** `O(N)`
- **Space Complexity:** `O(1)`

---

## Solution

See `main.go`.