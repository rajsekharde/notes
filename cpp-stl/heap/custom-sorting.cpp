#include <bits/stdc++.h>
using namespace std;

int main() {
    // Sort pairs by their second value
    unordered_map<int, int> mp;
    for(int i=1; i<=10; i++) {
        mp.insert({i, i%4}); // Creating a vector of pairs
    }
    priority_queue<pair<int, int>> pq;
    for(auto& it:mp)
        pq.push({it.second, it.first});
    /*
    Priority queue automatically sorts pairs by their first values.
    So the top pair is always the one with largest first value (actually
    the one with largest second value since the order was changed)
    */
}