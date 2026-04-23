#include <bits/stdc++.h>
using namespace std;

/*
Binary heap (max by default), duplicates allowed.
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

    // Get number of elements
    int s = pq.size();

    // Check if empty
    int ch = pq.empty();

    // Max-heap of pair<int, int>
    priority_queue<pair<int, int>, vector<pair<int, int>, greater<pair<int, int>>> maxHeap;

    // Min-heap of integers
    priority_queue<int, vector<int>, greater<int>> minHeap;

    // Min-heap of pair<int, int>
    priority_queue<pair<int, int>, vector<pair<int, int>>, greater<pair<int, int>>> minHeapP;

    return 0;
}

