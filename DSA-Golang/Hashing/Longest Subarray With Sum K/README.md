# Longest Subarray With Sum K

## Problem

Given an integer array `nums` and an integer `k`, return the length of the longest subarray whose sum is exactly `k`.

If no such subarray exists, return `0`.

### Example

**Input**
```
nums = [10,5,2,7,1,9]
k = 15
```

**Output**
```
4
```

---

## Intuition

Use the **Prefix Sum + Hash Map** technique.

While traversing the array, maintain the current prefix sum.

If the current prefix sum is `prefixSum`, then we need an earlier prefix sum of:

```
prefixSum - k
```

because:

```
currentPrefixSum - previousPrefixSum = k
```

If that prefix sum exists, the elements between the two indices have sum `k`.

To get the longest subarray, store only the **first occurrence** of each prefix sum.

---

## Approach

1. Create a hash map to store:
   ```
   prefixSum -> first index
   ```
2. Initialize:
   ```
   prefixSum = 0
   maxLength = 0
   ```
3. Store:
   ```
   0 -> -1
   ```
   to handle subarrays starting from index `0`.
4. Traverse the array.
5. Update the prefix sum.
6. Check whether `prefixSum - k` exists.
7. If it exists, calculate the subarray length and update the maximum.
8. Store the current prefix sum only if it is not already present.
9. Return `maxLength`.

---

## Complexity

- **Time Complexity:** `O(N)` average
- **Space Complexity:** `O(N)`

---

## Solution

See `main.go`.