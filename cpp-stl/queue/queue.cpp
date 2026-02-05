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

    // Remove last element
    q.pop();

    return 0;
}