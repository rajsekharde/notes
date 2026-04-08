# Letter Combinations of a phone number

Given, a string containing digits from 2-9. Return all possible letter combinations that can be formed using telephone buttons

digit-letter mapping:
```bash
2: a,b,c
3: d,e,f
4: g,h,i
5: j,k,l
6: m,n,o
7: p,q,r,s
8: t,u,v
9: w,x,y,z
```

Example: "23". Ans: {ad, ae, af, bd, be, bf, cd, ce, cf}

## Solution

Number of combinations = no.of digits * letters mapped to each digit.
Worst case: n 7's or 9's : Combinations = 4^9.

Time complexity = O(4^n)

Final strings need to be formed using recursion. Starting from first digit, and adding mapped letters in each recursive call.

```bash
"23"

f(2, "", {})

f(3, "a", {}), f(3, "b", {}), f(3, "c", {})

f(_, "ad", {ad}), f(_, "ae", {ad, ae}), f(_, "af", {ad, ae, af})...
```

Recursive function:
```bash
lt = vector<vector<char>> (digit-letter map)

void build(int c, string &s, string &digits, vector<string> &sol)
    {
        if(c == digits.length())
        {
            sol.push_back(s);
            return;
        }
        for(char ch: lt[digits[c]-'0'])
        {
            s.push_back(ch);
            build(c+1, s, digits, sol);
            s.pop_back();
        }
    }
```