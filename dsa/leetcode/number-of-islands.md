# Number of islands

Given an m * n 2D binary grid 'grid' which represents a map of 1's(land) and 0's(water), return the number
of islands. island is formed by connecting adjacent land cells (left, right, up, down)

# Approach 1

Traverse the matrix and if land is encountered, mark all connected land cells as visited using bfs/dfs.
Count the number of land cells encountered that are not visited.

Time complexity: O(m * n)

Space complexity: O(m * n)

```bash
    int numIslands(vector<vector<char>>& grid) {
        int m = grid.size(), n = grid[0].size(), ans = 0;
        vector<vector<bool>> visited(m, vector<bool>(n, false));
        for(int i=0; i<m; i++) {
            for(int j=0; j<n; j++) {
                char curr = grid[i][j];
                if(curr == '1' && !visited[i][j]) {
                    queue<pair<int, int>> q;
                    q.push({i, j});
                    visited[i][j] = true;
                    ans += 1;
                    while(!q.empty()) {
                        auto [x, y] = q.front();
                        q.pop();
                        if(x-1 >= 0 && grid[x-1][y] == '1' && !visited[x-1][y]) {q.push({x-1, y}); visited[x-1][y] = true;}
                        if(x+1 < m && grid[x+1][y] == '1' && !visited[x+1][y]) {q.push({x+1, y}); visited[x+1][y] = true;}
                        if(y-1 >= 0 && grid[x][y-1] == '1' && !visited[x][y-1]) {q.push({x, y-1}); visited[x][y-1] = true;}
                        if(y+1 < n && grid[x][y+1] == '1' && !visited[x][y+1]) {q.push({x, y+1}); visited[x][y+1] = true;}
                    }
                }
            }
        }
        return ans;
    }
```

# Approach 2

Visit each cell and instead of storing updating a visited matrix, change the cell values in-place if
connected land is encountered.

Time complexity: O(m * n)

Space complexity: O(1)

```bash
    int numIslands(vector<vector<char>>& grid) {
        int m = grid.size(), n = grid[0].size(), ans = 0;
        for(int i=0; i<m; i++) {
            for(int j=0; j<n; j++) {
                char curr = grid[i][j];
                if(curr == '1') {
                    queue<pair<int, int>> q;
                    q.push({i, j});
                    grid[i][j] = '0';
                    ans += 1;
                    while(!q.empty()) {
                        auto [x, y] = q.front();
                        q.pop();
                        if(x-1 >= 0 && grid[x-1][y] == '1') {q.push({x-1, y}); grid[x-1][y] = '0';}
                        if(x+1 < m && grid[x+1][y] == '1') {q.push({x+1, y}); grid[x+1][y] = '0';}
                        if(y-1 >= 0 && grid[x][y-1] == '1') {q.push({x, y-1}); grid[x][y-1] = '0';}
                        if(y+1 < n && grid[x][y+1] == '1') {q.push({x, y+1}); grid[x][y+1] = '0';}
                    }
                }
            }
        }
        return ans;
    }
```

# Approach 3

Recursive DFS solution. Cleaner code, same time/space complexity as #2

```bash
    int numIslands(vector<vector<char>>& grid) {
        if(grid.empty() || grid[0].empty()) return 0;
        int ans = 0;
        for(int i=0; i<grid.size(); i++) {
            for(int j=0; j<grid[0].size(); j++) {
                if(grid[i][j] == '1') {
                    ans +=1 ;
                    dfs(grid, i, j);
                }
            }
        }
        return ans;
    }
    void dfs(vector<vector<char>> &grid, int i, int j) {
        if(i < 0 || i >= grid.size() || j < 0 || j >= grid[0].size() || grid[i][j] != '1') return;
        grid[i][j] = '0';
        dfs(grid, i-1, j);
        dfs(grid, i+1, j);
        dfs(grid, i, j-1);
        dfs(grid, i, j+1);
    }
```