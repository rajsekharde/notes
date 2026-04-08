# Set Matrix Zeros

Given: m*n integer matrix. If an element is 0, set its entire row and column to 0's.

Example:
```bash
{
    {0, 1, 2, 0},
    {1, 2, 5, 6},
    {4, 7, 3, 8},
    {9, 3, 1, 5}
}

result:
{
    {0, 0, 0, 0},
    {0, 2, 5, 0},
    {0, 7, 3, 0},
    {0, 3, 1, 0}
}
```

## Brute force solution

maintain two hash sets for storing values of x, y that have been changed to 0

traverse the grid. for each (i,j), if number is 0, add i to x, j to y

traverse the grid again. this time if i belongs to x, or j belongs to y, change number to 0


Pattern: hash table, matrix


Time complexity:
- first traversal + storing hash: O(m*n)
- final traversal + retrieving hash: O(m*n)
- TC: O(m*n)
- Space complexity: O(m + n)

example:
```bash
{
    {1, 1, 1},
    {1, 0, 1},
    {1, 1, 1}
}

x: {1}
y: {1}

final:
{
    {1, 0, 1},
    {0, 0, 0},
    {1, 0, 1}
}
```