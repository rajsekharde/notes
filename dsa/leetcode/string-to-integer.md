# String to Integer (atoi)

Convert a string to a 32-bit signed integer

## Rules
- Strip off leading spaces
- Read sign (-): negative (+)/no-sign: positive (only before digits)
- If a non-digit character is encountered at any position, return
- Range: -2^31 to 2^31-1. Round off integer to this limit
- Return 0 if no digits read

## Algorithm

Initialise sign = 0 (+ve), ans = 0, i=0

Strip off leading spaces: i = 1st non-space char

Check if s[i] == '-' or '+' -> i++. if '-' -> sign = 1 (-ve)

Read next characters sequentially

If a digit found (0-9), convert to int, else break

Convert digit char to int: int d = ch - '0';

Check overflow:
```bash
if(ans > (INT_MAX - d)/10)
{
    return sign==1 ? INT_MIN : INT_MAX;
}
```

return ans

### Time complexity

T = O(n), since the full string is traversed once

## cpp-syntax encountered
```bash
# Convert char to int:
char ch = '5';
int d = ch - '0'; // d = 5

# 32-bit int limits:
INT_MAX = 2^31 - 1
INT_MIN = -2^31

# Check if a char is a digit
if(ch > '0' && ch < '9')
```
