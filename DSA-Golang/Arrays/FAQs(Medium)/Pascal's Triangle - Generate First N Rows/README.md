# Pascal's Triangle - Generate First N Rows

## Problem

Given an integer `n`, return the first `n` rows of Pascal's Triangle.

### Example

**Input**
```
n = 5
```

**Output**
```
[
 [1],
 [1,1],
 [1,2,1],
 [1,3,3,1],
 [1,4,6,4,1]
]
```

---

## Intuition

Each row of Pascal's Triangle starts and ends with `1`.

Every middle element is the sum of the two elements directly above it from the previous row.

Instead of calculating combinations separately, we build the triangle row by row using the values from the previous row.

---

## Approach

1. Create a 2D slice to store the triangle.
2. For each row:
   - Create a slice of size `i + 1`.
   - Set the first and last elements to `1`.
   - Compute the middle elements using the previous row.
3. Return the completed triangle.

---

## Algorithm

1. Initialize an empty 2D slice.
2. Iterate from row `0` to `n-1`.
3. Set the first and last elements of each row to `1`.
4. Fill the middle elements:
   - `triangle[i][j] = triangle[i-1][j-1] + triangle[i-1][j]`
5. Return the triangle.

---

## Complexity

- **Time Complexity:** `O(n²)`
- **Space Complexity:** `O(n²)`

---

## Solution

See `main.go`.