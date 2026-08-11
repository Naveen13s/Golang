# Sum of Highest and Lowest Frequency

## Problem

Given an integer array `arr`, return the sum of the highest occurring frequency and the lowest occurring frequency.

### Example

**Input**
```
arr = [1,2,2,3,3,3]
```

**Output**
```
4
```

---

## Intuition

Use a hash map to count the frequency of each element.

Once all frequencies are known, find the maximum and minimum frequency and return their sum.

---

## Approach

1. Count the frequency of every element using a hash map.
2. Traverse the frequency map.
3. Track:
   - Highest frequency.
   - Lowest frequency.
4. Return their sum.

---

## Algorithm

1. Create an empty frequency map.
2. Count the occurrences of each element.
3. Find the maximum and minimum frequency.
4. Return `maxFreq + minFreq`.

---

## Complexity

- **Time Complexity:** `O(N)`
- **Space Complexity:** `O(K)`

---

## Solution

See `main.go`.