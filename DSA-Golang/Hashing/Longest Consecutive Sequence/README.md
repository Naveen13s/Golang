# Longest Consecutive Sequence

## Problem

Given an integer array `nums`, return the length of the longest sequence of consecutive integers.

The integers in the sequence may appear in any order.

### Example

**Input**
```
nums = [100,4,200,1,3,2]
```

**Output**
```
4
```

---

## Intuition

Instead of sorting or checking every possible sequence, store all numbers in a hash set.

A number is considered the start of a sequence only if its previous number (`num - 1`) is not present.

From each starting point, count the consecutive numbers to determine the sequence length.

---

## Approach

1. Insert all numbers into a hash set.
2. Traverse the array.
3. If `num - 1` is not in the set, start counting the sequence.
4. Continue while `num + 1` exists.
5. Update the maximum sequence length.
6. Return the maximum length.

---

## Algorithm

1. Create a hash set containing all array elements.
2. For each number:
   - If it has no predecessor, start a new sequence.
   - Count consecutive numbers.
3. Track the maximum sequence length.
4. Return the result.

---

## Complexity

- **Time Complexity:** `O(N)`
- **Space Complexity:** `O(N)`

---

## Solution

See `main.go`.