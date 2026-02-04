#include <bits/stdc++.h>
using namespace std;

int main() {

    // Frequency of characters in a string
    string s = "asdegfgsdtgjkfgh";
    map<char, int> m1;
    for (char ch: s)
        m1[ch]++;
    for (auto it:m1)
        cout << it.first << " " << it.second << "\n";
    
    return 0;
}