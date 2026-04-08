# Longest substring without repeating characters

Given: String s. Find longest substring without duplicate characters

Example: "abcabcbb". Ans: "abc" or "bca" or "cab"

## Solution

Maintain a window of characters using two pointers: l, r, such that all characters inside the window are unique. r expands the window to the right in each iteraation and l shrinks it to right when a duplicate letter is found. 

Maintain a hash map seen(letter(char), position(int)) for the current window

shift r one position at a time. if seen[new char] = present, and within window, set l = seen[new char] + 1, seen[new char] = r, update maximum length

```bash
int lengthOfLongestSubstring(string s) {
        int n = s.length(), ls=0;
        unordered_map<char, int> seen;
        for(int l=0, r=0; r<n; r++) {
            if(seen.count(s[r]) && seen[s[r]]>=l)
                l = seen[s[r]]+1;
            ls = max(ls, (r-l+1));
            seen[s[r]] = r;
        }
        return ls;
    }
```

Time complexity = O(n)