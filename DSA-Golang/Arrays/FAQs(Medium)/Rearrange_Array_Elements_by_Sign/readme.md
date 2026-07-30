# Rearrange Array Elements by Sign

## Problem

Given an integer array `nums` of even length containing an equal number of positive and negative integers, rearrange the array such that:

- Every consecutive pair has opposite signs.
- The relative order of positive numbers is preserved.
- The relative order of negative numbers is preserved.
- The array starts with a positive number.

### Example

**Input**
```
nums = [3,1,-2,-5,2,-4]
```

**Output**
```
[3,-2,1,-5,2,-4]
```

---

## Intuition

Since the number of positive and negative integers is equal, we know exactly where each type of element should be placed.

Positive numbers occupy even indices (`0, 2, 4, ...`) and negative numbers occupy odd indices (`1, 3, 5, ...`).

By placing each number at its respective index while traversing the array once, we preserve the original order of both positive and negative elements.

---

## Approach

1. Create a result array of size `n`.
2. Initialize two pointers:
   - `pos = 0` for even indices.
   - `neg = 1` for odd indices.
3. Traverse the input array.
4. Place positive numbers at `pos` and increment `pos` by `2`.
5. Place negative numbers at `neg` and increment `neg` by `2`.
6. Return the result array.

---

## Algorithm

1. Create an empty result array.
2. Set `pos = 0` and `neg = 1`.
3. Traverse every element:
   - If positive, place it at `result[pos]`.
   - Otherwise, place it at `result[neg]`.
4. Return the rearranged array.

---

## Complexity

- **Time Complexity:** `O(N)`
- **Space Complexity:** `O(N)`

---

## Solution

See `main.go`.