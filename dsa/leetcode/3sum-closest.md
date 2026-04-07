# 3Sum Closest

Given: integer array nums of n integers, int target. Find three integers at distinct indices such that their sum is closest to target. Return closest sum.

## Brute force approach

Check every 3-number combination in the array and find closest sum. Time complexity = O(n^3).

## Optimised approach

Sort the array. Take one number at a time, starting from left. For each number, search for a pair of numbers such that sum of three numbers is closest to target. Using two-pointers for the pair.

```bash
// pseudo-code

sort(nums) in ascending order

for i=0 to n-1
{
    l=i+1, r=n-1
    while l<r
    {
        sum = nums[i] + nums[l] + nums[r]
        if sum closer than cl-sum, cl-sum = sum

        if sum > target : r--
        if sum < target : l++
        if sum = target : return target // best case scenario
    }
}
```

Time complexity: O(n^2). The pair-checking requires iterating throught the array, for each i

Pattern used: Sorting + Two-Pointers. Take one element fixed and perform 2-Sum for the pair.

In each iteration inside while loop, keep decreasing the gap b/w sum and target by doing l++ if sum < target (increasing sum) or by doing r-- if sum > target (decreasing sum)