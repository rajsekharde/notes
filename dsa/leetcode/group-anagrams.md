# Group Anagrams

Given: An array of strings. Return an array of groups of strings that are anagrams of each other, in any order.

Example:
```bash
Input: strs = ["act","pots","tops","cat","stop","hat"]

Output: [["hat"],["act", "cat"],["stop", "pots", "tops"]]
```

## Approach 1

Maintain a hash map of sorted strings and its anagrams. iterate through the array, add the current string to anagrams list where key = sorted string


Time complexity: Iterating: O(n), Sorting: O(klogk), hashing/searching: O(1). Overall: O(n*klogk), n = number of strings, k = maximum length of a string.


```bash
vector<vector<string>> groupAnagrams(vector<string>& strs) {
        unordered_map<string, vector<string>> mp;
        vector<vector<string>> ans;

        for(string s: strs)
        {
            string a = s;
            sort(s.begin(), s.end());
            mp[s].push_back(a);
        }

        for(auto it: mp)
        {
            ans.push_back(it.second);
        }

        return ans;
    }
```