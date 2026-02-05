#include <bits/stdc++.h>
using namespace std;

/*
Binary heap (max by default), duplicates allowes.
Maintains Max/Min at top
*/

int main() {
    // Creation of Max-Heap
    priority_queue<int> pq;

    // Insertion
    pq.push(5);
    pq.push(20);

    // Accessing top element
    cout << pq.top();

    // Removing top element
    pq.pop();


    // Min-heap
    priority_queue<int, vector<int>, greater<int>> minHeap;

    return 0;
}