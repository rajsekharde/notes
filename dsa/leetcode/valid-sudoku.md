# Valid Sudoku

Given, a 9x9 matrix representing a sudoku board, check if the filled cells are valid.

Conditions for validity:
- Each row must contain digits 1-9 without repetition
- Each column must ""
- Each of the 9 3x3 grids must ""

## Brute force solution

Check all rows: O(9*9) = O(81)

Check all columns: O(81)

Check all grids: O(81)

Time complexity: O(3*81) ~ O(1)


For checking rows, iterate through a row and check if a hash set contains the number. if yes: return false, else add number to hash set