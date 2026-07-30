# Leaders in an Array

## Problem

Given an integer array `nums`, return all the leaders in the array.

A leader is an element that is strictly greater than all the elements to its right.

The rightmost element is always a leader.

### Example

**Input**
```
nums = [16, 17, 4, 3, 5, 2]
```

**Output**
```
[17, 5, 2]
```

---

## Intuition

Instead of checking every element against all elements to its right, traverse the array from right to left.

Keep track of the maximum element seen so far.

If the current element is greater than this maximum, it is a leader.

Since leaders are collected in reverse order, reverse the result before returning it.

---

## Approach

1. Start from the last element.
2. Initialize `maxRight` with the last element.
3. Add the last element to the result.
4. Traverse from right to left.
5. If the current element is greater than `maxRight`:
   - Add it to the result.
   - Update `maxRight`.
6. Reverse the result.
7. Return the leaders.

---

## Algorithm

1. Initialize `maxRight` as the last element.
2. Traverse the array from right to left.
3. Add elements greater than `maxRight` to the result.
4. Update `maxRight`.
5. Reverse the result.
6. Return the result.

---

## Complexity

- **Time Complexity:** `O(N)`
- **Space Complexity:** `O(N)`

---

## Solution

See `main.go`.