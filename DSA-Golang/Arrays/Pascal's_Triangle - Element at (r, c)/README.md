# Pascal's Triangle - Element at (r, c)

## Problem

Given two integers `r` and `c`, return the value at the `r`th row and `c`th column (1-indexed) of Pascal's Triangle.

### Example

**Input**
```
r = 5
c = 3
```

**Output**
```
6
```

---

## Intuition

The value at the `r`th row and `c`th column is equal to the binomial coefficient:

```
C(r-1, c-1)
```

Instead of generating the entire triangle, compute this value directly using an iterative combination formula.

---

## Approach

1. Convert the indices:
   - `n = r - 1`
   - `k = c - 1`
2. Use `k = min(k, n-k)` to minimize computations.
3. Compute the combination iteratively.
4. Return the result.

---

## Algorithm

1. Compute `n = r - 1` and `k = c - 1`.
2. Reduce `k` using symmetry.
3. Initialize `result = 1`.
4. For each `i` from `0` to `k-1`:
   - Multiply by `(n - i)`.
   - Divide by `(i + 1)`.
5. Return `result`.

---

## Complexity

- **Time Complexity:** `O(min(c-1, r-c))`
- **Space Complexity:** `O(1)`

---

## Solution

See `main.go`.