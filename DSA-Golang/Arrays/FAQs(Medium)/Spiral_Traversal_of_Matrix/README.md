# Spiral Traversal of Matrix

## Problem

Given an `M × N` matrix, return all its elements in clockwise spiral order.

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
[1,2,3,6,9,8,7,4,5]
```

---

## Intuition

Instead of marking cells as visited, maintain four boundaries representing the current unvisited rectangle of the matrix.

At each iteration:

- Traverse the top row.
- Traverse the right column.
- Traverse the bottom row.
- Traverse the left column.

After each traversal, shrink the corresponding boundary.

Repeat until all elements are visited.

---

## Approach

1. Initialize:
   - `top`
   - `bottom`
   - `left`
   - `right`
2. Traverse:
   - Left → Right
   - Top → Bottom
   - Right → Left
   - Bottom → Top
3. Update boundaries after each traversal.
4. Continue while valid boundaries remain.

---

## Algorithm

1. Initialize four boundary pointers.
2. Repeat while `top <= bottom` and `left <= right`.
3. Traverse the four sides of the current layer.
4. Shrink the boundaries.
5. Return the result.

---

## Complexity

- **Time Complexity:** `O(M × N)`
- **Space Complexity:** `O(1)` (excluding output array)

---

## Solution

See `main.go`.