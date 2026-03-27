#include <bits/stdc++.h>
using namespace std;

/*
FIFO container, duplicates allowed, order maintained
*/

int main() {
    // Creation
    queue<int> q;

    // Insertion at end
    q.push(1);

    // Accessing first element
    cout << q.front();

    // Accessing last element
    cout << q.back();

    // Remove first element
    q.pop();

    // Get size
    int s = q.size();

    // Check if empty
    bool e = q.empty();

    return 0;
}
