# Pascal's Triangle - Print the r-th Row

## Problem

Given an integer `r`, return all the values in the `r`th row (1-indexed) of Pascal's Triangle.

### Example

**Input**
```
r = 5
```

**Output**
```
[1,4,6,4,1]
```

---

## Intuition

The elements of the `r`th row are binomial coefficients:

```
C(r-1,0), C(r-1,1), ..., C(r-1,r-1)
```

Instead of generating the entire Pascal's Triangle, we can compute each coefficient iteratively using the previous one.

This gives an efficient linear-time solution.

---

## Approach

1. Let `n = r - 1`.
2. Start with `current = 1`.
3. Add `current` to the answer.
4. Generate the next coefficient using:

```
current = current × (n - i) / (i + 1)
```

5. Repeat until all elements are generated.

---

## Algorithm

1. Initialize `current = 1`.
2. Add it to the result.
3. Iterate `i` from `0` to `n-1`.
4. Compute the next coefficient.
5. Append it to the answer.
6. Return the row.

---

## Complexity

- **Time Complexity:** `O(r)`
- **Space Complexity:** `O(r)`

---

## Solution

See `main.go`.