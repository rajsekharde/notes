#include <bits/stdc++.h>
using namespace std;

/*
Anagram: Two strings are considered anagrams if they consist of the same
characters but with or without same order

Eg: "cat" <-> "tac"; "anagram" <-> "nagaram";
"cat" <-x-> "car" (t, r not in both)
"abc" <-x-> "aabc" (length not same, extra a in 2nd string)

Algorithm for checking if two strings are anagrams:
1. Store the frequency of characters of first string in a hash map
2. Traverse through the second string. If a charcter has freq = 0 or
   doesn't exist in first string, return false
3. Reduce the freq of the character by one
*/

bool isAnagram(string& s, string& t) {
    if(s.length() != t.length())
        return false;
    
    unordered_map<char, int> m;
    for(char ch: s)
        m[ch]++;
    for(char ch: t) {
        if(m.count(ch)==0 || m[ch]==0)
            return false;
        m[ch]--;
    }
    return true;
}

int main() {
    string a = "anagramanagram", b = "cat", c = "nagaramanagram";
    cout << isAnagram(a, c) << " " << isAnagram(a, b) << "\n";
    return 0;
}