#include <bits/stdc++.h>
using namespace std;

// Time complexity- O(1) for insertion, deletion, lookup

int main() {
	// Creating an empty map
	unordered_map<int, string> m1;

	// Creating and initializing a map using a list
	unordered_map<string, int> m2 = {{"Hello", 1}, {"World", 2}};

	// Inserting values
	m1.insert({1, "A"});
	m1[2] = "B";

	// Accessing values by keys
	cout << m1[1] << "\n";
	cout << m1.at(2) << "\n";

	// Updating values
	m1[1] = "AA";
	m1.at(2) = "BB";

	// Deleting values by keys
	m1.erase(1);

	// Finding elements
	auto f = m1.find(1); // Returns key if found, otherwise m1.end()
    int c = m1.count(1); // Returns 0 if not found

	// Traversal
	for (auto &p: m2)
		cout << p.first << " " << p.second << "\n";
	for (auto it = m2.begin(); it != m2.end(); ++it)
		cout << it->first << " " << it->second << "\n";
	return 0;
}