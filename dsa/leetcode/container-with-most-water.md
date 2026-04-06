# Container with most water

Given: integer array of size n, arr[i] = height of column. Find two columns that form a container that can hold the maximum amount of water

Similar to two-sum problem. Need to find two positions that give an optimal result.

## Solutions

### Brute force

Calculate volume of container for each pair of columns, and find the one with highest volume.
Run nested loops: i=0 to n, j=i+1 to n. Volume = min(h[i], h[j]) * abs(i-j).
Time complexity = O(n^2)

### Optimal solution

Take the widest container using two pointers- l=0, r=n. in each iteration, shrink the container by moving the smaller line. if height[l]<height[r] then l++ else r--. Iterate till l, r cross.

Since height of container is limited by the smaller line, the smallest of the two lines is moved in each iteration. The optimal solution is never missed.

Pattern: Two-pointers + Greedy elimination

Dry run:

height = [1,8,6,2,5,4,8,3,7]

l   r   v   maxV
0   8   8   8
1   8   49  49
1   7   18  49
1   6   40  49
1   5   16  49
1   4   15  49
1   3   4   49
1   2   6   49
1   1   -   -

maxV = 49