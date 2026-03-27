#include <bits/stdc++.h>
using namespace std;

/*
Double-ended queue, duplicates allowed.
Allows push, pop from both ends
*/

int main() {
    // Creation
    deque<int> dq;

    // Insertion at end
    dq.push_back(1);

    // Insertion at front
    dq.push_front(0);

    // Access first element
    cout << dq.front();

    // Access last element
    cout << dq.back();

    // Remove first element
    dq.pop_front();

    // Remove last element
    dq.pop_back();

    // Get size
    int s = dq.size();

    // Check if empty
    bool e = q.empty();
    
    return 0;
}
