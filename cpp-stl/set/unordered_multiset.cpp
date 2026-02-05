#include <bits/stdc++.h>
using namespace std;

/*
Hash-based, Duplicates allowed, unordered

Time complexity:

Insertion- O(1)

Deletion- O(1)

Lookup- O(1)
*/

int main() {
    // Declaration
    unordered_multiset<int> s1;
    unordered_multiset<int> s2 = {1, 2, 5, 3, 4};

    // Insertion
    s1.insert(1);
    s1.insert(1);

    // Deletion
    s1.erase(1);

    // Lookup
    auto r = s2.find(2); // returns 1 or 0
    int c = s2.count(5); // returns count or 0

    // Traversal
    for (auto it = s2.begin(); it != s2.end(); ++it)
        cout << *it << " ";
}
