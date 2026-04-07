# 3Sum

Given: Integer array nums, return all triplets [nums[i], nums[j], nums[k]] such that i != j != k, nums[i] + nums[j] + nums[k] = 0. Order of triplets and output does not matter

## Brute force approach

n = nums.size()

Three nested loops: i=0 to n-1, j = 0 to n-1, k = 0 to n-1
for each i,j,k: check if nums[i] + nums[j] + nums[k] == 0. yes -> store nums

time complexity: O(n^3). Difficult to remove duplicate triplets

## Optimised approach

nums = [-1,0,1,2,-1,-4]

sorted: nums = [-4,-1,-1,0,1,2]

take one number from nums at a time starting from left and check all pairs to the right of the number. if sum of three no.s = 0, store no.s, l++, r--. if sum > 0 : right-- else left++

```bash
for i = 0 to n-1
{
    // check for duplicates
    if nums[i] == nums[i-1] : continue

    l = i+1, r = n-1
    while(l<r)
    {
        sum = nums[i] + nums[l] + nums[r]
        if s = 0:
            store {n-i, n-l, n-r}
            l++, r--
            keep skipping duplicate values of l & r
        else if s > 0 : r--
        else if s < 0 : l++
    }
}
```

time complexity: O(n^2). Iterating throught the array once. for each iteration, checking all elements to right.

nums = [-4,-1,-1,0,1,2]

i   l   r   s
-4  -1  2   -3  l++
-4  -1  2   -3  l++
-4  0   2   -2  l++
-4  1   2   -1  l++
-4  2   2   l=r, break
-1  -1  2   0   store {-1,-1,2}, l++, r--
-1  0   1   0   store {-1,0,1}, l++, r--
-1  1   0   0   l>r, break
-1  i- same as prev, skip
0   1   2   3   r--
0   1   1   l=r, break
1   2   2   l=r, break
2   (6)   (6)   invalid l,r, break

final ans: {-1,-1,2}, {-1,0,1}