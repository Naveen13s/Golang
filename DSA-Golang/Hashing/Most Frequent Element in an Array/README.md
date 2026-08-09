# Most Frequent Element in an Array

## Problem

Given an integer array `nums`, return the element with the highest frequency.

If multiple elements have the same maximum frequency, return the smallest one.

### Example

**Input**
```
nums = [1,2,2,3,3]
```

**Output**
```
2
```

---

## Intuition

Instead of counting occurrences repeatedly, use a hash map to store the frequency of each element.

After building the frequency map, iterate through it to find:

- The highest frequency.
- The smallest element when frequencies are equal.

---

## Approach

1. Create a frequency map.
2. Count occurrences of every element.
3. Traverse the map:
   - Update the answer when a higher frequency is found.
   - If frequencies are equal, keep the smaller element.
4. Return the answer.

---

## Algorithm

1. Initialize an empty hash map.
2. Count the frequency of each element.
3. Track:
   - `maxFreq`
   - `answer`
4. Return the most frequent element.

---

## Complexity

- **Time Complexity:** `O(N)`
- **Space Complexity:** `O(N)`

---

## Solution

See `main.go`.