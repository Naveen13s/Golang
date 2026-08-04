# 3Sum

## Problem

Given an integer array `nums`, return all unique triplets whose sum equals `0`.

Each triplet must satisfy:

- `i != j`
- `i != k`
- `j != k`

Duplicate triplets are not allowed.

### Example

**Input**
```
nums = [-1,0,1,2,-1,-4]
```

**Output**
```
[
 [-1,-1,2],
 [-1,0,1]
]
```

---

## Intuition

A brute-force solution checks every possible triplet, resulting in `O(N³)` time.

A more efficient approach is to first sort the array.

After sorting, fix one element and use the Two Pointer technique to find the remaining two elements whose sum equals the negative of the fixed element.

Sorting also makes it easy to skip duplicate values and avoid repeated triplets.

---

## Approach

1. Sort the array.
2. Iterate through each element as the first element of the triplet.
3. Skip duplicate values.
4. Use two pointers:
   - `left = i + 1`
   - `right = n - 1`
5. Calculate the sum.
6. Move pointers based on the sum.
7. Store valid triplets while skipping duplicates.

---

## Algorithm

1. Sort the array.
2. Traverse the array.
3. Skip duplicate first elements.
4. Use two pointers to search for pairs.
5. Add valid triplets.
6. Skip duplicate second and third elements.
7. Return all unique triplets.

---

## Complexity

- **Time Complexity:** `O(N²)`
- **Space Complexity:** `O(1)` (excluding output)

---

## Solution

See `main.go`.