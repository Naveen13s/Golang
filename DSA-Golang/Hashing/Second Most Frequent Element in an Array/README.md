# Second Most Frequent Element in an Array

## Problem

Given an integer array `nums`, return the second most frequent element.

If multiple elements have the same second highest frequency, return the smallest among them.

If no second most frequent element exists, return `-1`.

### Example

**Input**
```
nums = [1,2,2,3,3,3]
```

**Output**
```
2
```

---

## Intuition

Use a hash map to count the frequency of each element.

Find the highest frequency, then find the largest frequency smaller than it. If multiple elements have this frequency, choose the smallest element.

---

## Approach

1. Count the frequency of every element.
2. Find the maximum frequency.
3. Find the second highest distinct frequency.
4. Return the smallest element having that frequency.
5. If no second highest frequency exists, return `-1`.

---

## Algorithm

1. Build a frequency map.
2. Find the maximum frequency.
3. Traverse the map again to find the second highest frequency.
4. Handle ties by choosing the smaller element.
5. Return the answer.

---

## Complexity

- **Time Complexity:** `O(N)`
- **Space Complexity:** `O(N)`

---

## Solution

See `main.go`.