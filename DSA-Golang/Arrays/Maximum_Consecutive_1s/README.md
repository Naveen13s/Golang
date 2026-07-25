# Maximum Consecutive 1s

## Problem

Given a binary array `nums`, return the maximum number of consecutive `1`s in the array.

A binary array contains only `0`s and `1`s.

### Example

**Input**
```
nums = [1, 1, 0, 1, 1, 1]
```

**Output**
```
3
```

**Explanation**

The longest sequence of consecutive `1`s is:

```
[1, 1, 1]
```

Therefore, the answer is `3`.

---

## Intuition

We need to find the longest continuous sequence of `1`s.

While traversing the array, we maintain a counter that represents the length of the current sequence of consecutive `1`s.

Whenever we encounter a `1`, we increment the counter. If we encounter a `0`, the consecutive sequence is broken, so we reset the counter to `0`.

At every step, we compare the current count with the maximum count found so far and update the maximum when necessary.

---

## Approach

Maintain two variables:

- `count` to track the current number of consecutive `1`s.
- `maxCount` to track the maximum consecutive `1`s found so far.

Traverse the array:

1. If the current element is `1`, increment `count`.
2. Update `maxCount` if `count` is greater.
3. If the current element is `0`, reset `count` to `0`.
4. After traversing the entire array, return `maxCount`.

---

## Algorithm

1. Initialize `count = 0`.
2. Initialize `maxCount = 0`.
3. Traverse every element in `nums`.
4. If the element is `1`:
   - Increment `count`.
   - Update `maxCount` if necessary.
5. Otherwise, reset `count = 0`.
6. Return `maxCount`.

---

## Complexity

- **Time Complexity:** `O(N)`
- **Space Complexity:** `O(1)`

---

## Solution

See `main.go`.