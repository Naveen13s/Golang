# Count Subarrays With Sum K

## Problem

Given an integer array `nums` and an integer `k`, return the total number of subarrays whose sum is exactly `k`.

### Example

**Input**
```text
nums = [1,2,3]
k = 3
```

**Output**
```text
2
```

The valid subarrays are `[1,2]` and `[3]`.

---

## Intuition

Use the **Prefix Sum + Hash Map** technique.

For the current prefix sum `prefixSum`, we need a previous prefix sum of:

```text
prefixSum - k
```

because:

```text
currentPrefixSum - previousPrefixSum = k
```

The frequency of `prefixSum - k` tells us how many valid subarrays end at the current index.

Unlike the longest subarray problem, here we store the **frequency** of each prefix sum because multiple occurrences can create multiple valid subarrays.

---

## Approach

1. Create a hash map to store:
   ```text
   prefixSum -> frequency
   ```
2. Initialize `prefixSum = 0` and `count = 0`.
3. Store:
   ```text
   0 -> 1
   ```
4. Traverse the array.
5. Update the prefix sum.
6. Add the frequency of `prefixSum - k` to the answer.
7. Increase the frequency of the current prefix sum.
8. Return the count.

---

## Algorithm

1. Initialize `prefixSum = 0`.
2. Initialize `count = 0`.
3. Set `freq[0] = 1`.
4. For every element:
   - Add it to `prefixSum`.
   - Add `freq[prefixSum-k]` to `count`.
   - Increment `freq[prefixSum]`.
5. Return `count`.

---

## Complexity

- **Time Complexity:** `O(N)`
- **Space Complexity:** `O(N)`

---

## Solution

See `main.go`.