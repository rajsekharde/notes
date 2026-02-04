#include <bits/stdc++.h>
using namespace std;

/*
Time complexity:

Insertion- O(log(n))

Deletion- O(log(n))

Lookup- O(log(n))
*/

int main() {
    // Declaration
    set<int> s1;
    set<int> s2 = {1, 2, 5, 3, 4};

    // Insertion
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