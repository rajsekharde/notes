# Valid Anagram

Given two strings, return true if they are anagrams of each other, else false. Anagram: if two strings contain the same set of characters, but in any order.

Example: "anagram", "gmranaa"

## Approach 1

Sort the two strings, and compare letters one by one, starting from left.

Time complexity: sorting: O(nlogn), Iteration: O(n). Overall: O(nlogn)

```bash
bool isAnagram(string s, string t) {
        if(s.length() != t.length()) return false;
        sort(s.begin(), s.end());
        sort(t.begin(), t.end());
        for(int i=0; i<s.length(); i++)
        {
            if(s[i] != t[i]) return false;
        }
        return true;
    }
```

## Approach 2

store the frequency of letters of the strings in two hash maps, compare final hash maps.

Time complexity: Hashing: O(n), Comparing maps: O(n). Overall: O(n)

```bash
bool isAnagram(string s, string t) {
        if(s.length() != t.length()) return false;
        unordered_map<char, int> m1;
        unordered_map<char, int> m2;
        for(char ch: s) m1[ch]++;
        for(char ch: t) m2[ch]++;
        if(m1 == m2) return true;
        else return false;
    }
```