# Majority Element

## Problem

Given an integer array `nums` of size `n`, return the majority element.

The majority element is the element that appears more than `n/2` times.

It is guaranteed that a majority element always exists.

### Example

**Input**
```
nums = [2,2,1,1,1,2,2]
```

**Output**
```
2
```

---

## Intuition

Since one element appears more than half the time, it cannot be completely cancelled out by the remaining elements.

Moore's Voting Algorithm repeatedly cancels occurrences of different elements. The element left at the end is the majority element.

---

## Approach

1. Initialize a candidate and a count.
2. Traverse the array.
3. If the count becomes `0`, choose the current element as the new candidate.
4. If the current element equals the candidate, increment the count.
5. Otherwise, decrement the count.
6. Return the candidate.

Since the problem guarantees a majority element exists, no verification step is required.

---

## Algorithm

1. Set `candidate = 0` and `count = 0`.
2. Traverse the array.
3. If `count == 0`, update the candidate.
4. Increase or decrease the count depending on whether the current element matches the candidate.
5. Return the candidate.

---

## Complexity

- **Time Complexity:** `O(N)`
- **Space Complexity:** `O(1)`

---

## Solution

See `main.go`.