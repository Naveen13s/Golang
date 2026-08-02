# Rotate Matrix by 90 Degrees Clockwise

## Problem

Given an `N × N` matrix, rotate it by **90 degrees clockwise**.

The rotation must be performed **in-place**, without using another matrix.

### Example

**Input**
```
[
 [1,2,3],
 [4,5,6],
 [7,8,9]
]
```

**Output**
```
[
 [7,4,1],
 [8,5,2],
 [9,6,3]
]
```

---

## Intuition

Instead of creating another matrix, rotate the matrix in-place using two operations:

1. Transpose the matrix.
2. Reverse every row.

The transpose converts rows into columns, and reversing each row completes the clockwise rotation.

---

## Approach

1. Transpose the matrix by swapping:
   - `matrix[i][j]` with `matrix[j][i]`
2. Reverse every row.
3. The matrix is now rotated by 90 degrees clockwise.

---

## Algorithm

1. Find the size of the matrix.
2. Transpose the matrix.
3. Reverse each row.
4. Return the modified matrix.

---

## Complexity

- **Time Complexity:** `O(N²)`
- **Space Complexity:** `O(1)`

---

## Solution

See `main.go`.