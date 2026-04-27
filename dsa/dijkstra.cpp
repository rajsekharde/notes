#include <bits/stdc++.h>
using namespace std;

typedef pair<int, int> pii;

vector<int> dijkstra(int n, vector<vector<pii>> &adj, int source) {
    priority_queue<pii, vector<pii>, greater<pii>> pq;
    vector<bool> visited(n, false);
    vector<int> distances(n, INT_MAX);

    pq.push({0, source});
    distances[source] = 0;
    visited[source] = true;
    
    while(!pq.empty()) {
        auto [dist, node] = pq.top();
        pq.pop();

        for(auto neighbour: adj[node]) {
            int next_node = neighbour.first;
            int d = neighbour.second;
            if(distances[node] + d < distances[next_node]) {
                distances[next_node] = distances[node] + d;
            }
            if(visited[next_node] == false) {
                pq.push({distances[next_node], next_node});
                visited[next_node] = true;
            }
        }
    }
        return distances;
}

int main() {
    int n = 3;
    vector<vector<pii>> adj(n);

    adj[0].push_back({1, 10});
    adj[1].push_back({0, 10});

    adj[1].push_back({2, 20});
    adj[2].push_back({1, 20});

    vector<int> distances = dijkstra(n, adj, 0);

    cout << "Lenght of shortest path from 0:\n";
    for(int i=0; i<n; i++) {
        cout << i << " " << distances[i] << "\n";
    }

    return 0;
}