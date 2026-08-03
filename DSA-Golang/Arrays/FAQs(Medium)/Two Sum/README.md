# Two Sum

## Problem

Given an integer array `nums` and an integer `target`, return the indices of the two numbers whose sum equals `target`.

Each input has exactly one solution, and the same element cannot be used twice.

### Example

**Input**
```
nums = [2,7,11,15]
target = 9
```

**Output**
```
[0,1]
```

---

## Intuition

Instead of checking every possible pair, use a hash map to store previously visited numbers along with their indices.

For each number, compute the required complement:

```
complement = target - currentNumber
```

If the complement already exists in the map, the answer has been found.

Otherwise, store the current number and continue.

---

## Approach

1. Create an empty hash map.
2. Traverse the array.
3. Compute the complement.
4. If the complement exists in the map, return both indices.
5. Otherwise, store the current number and its index.
6. Continue until the solution is found.

---

## Algorithm

1. Initialize an empty hash map.
2. Traverse the array.
3. Compute:
   ```
   complement = target - nums[i]
   ```
4. If the complement exists:
   - Return its index and the current index.
5. Otherwise, store:
   ```
   nums[i] -> i
   ```
6. Return the result.

---

## Complexity

- **Time Complexity:** `O(N)`
- **Space Complexity:** `O(N)`

---

## Solution

See `main.go`.