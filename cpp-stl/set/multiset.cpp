#include <bits/stdc++.h>
using namespace std;

/*
Balanced BST, duplicates allowed, sorted in ascending order

Time complexity:

Insertion- O(log(n))

Deletion- O(log(n))

Lookup- O(log(n))
*/

int main() {
    // Declaration
    multiset<int> s1;
    multiset<int> s2 = {1, 2, 5, 3, 4};

    // Insertion
    s1.insert(1);
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
    
    auto range = s1.equal_range(1); // Returns iterators to all 1s
}