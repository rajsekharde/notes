# Valid Palindrome

Give, a string containing ASCII characters, return true if it's a palindrome, otherwise false.

Topics: two pointers, string

Example:
```bash
Input: s = "Was it a car or a cat I saw?"

Output: true
```

## Approach 1

Iterate once and remove all non-alphanumeric characters, and convert letters to lower case
```bash
s = "Was it a car or a cat I saw"
t = "wasitacaroracatisaw"
```

Use two pointers: l, r and keep checking characters from opposite ends. if chars not equal, return false

Time Complexity: O(n)

```bash
bool isPalindrome(string s) {
        string t = "";
        for(char ch: s)
        {
            ch = tolower(ch);
            if((ch >= 'a' && ch <= 'z') || (ch >= '0' && ch <= '9'))
                t += ch;
        }
        int l=0,r=t.length()-1;
        while(l<r)
        {
            if(t[l] != t[r]) return false;
            l++;
            r--;
        }
        return true;
    }
```