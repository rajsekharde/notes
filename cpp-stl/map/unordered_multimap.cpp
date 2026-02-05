#include <bits/stdc++.h>
using namespace std;

/*
Hash-table key-value, duplicates allowed, unordered

Time complexity- O(1) for insertion, deletion, lookup
*/

int main() {
	// Creating an empty map
	unordered_multimap<int, string> m1;

	// Creating and initializing a map using a list
	unordered_multimap<string, int> m2 = {{"Hello", 1}, {"World", 2}};

	// Inserting values
	m1.insert({1, "A"});
    m1.insert({1, "B"});

    // Iteration
    for(auto it = m1.equal_range(1).first; it != m1.equal_range(1).second; ++it)
        cout << it->second << endl;

    return 0;
}