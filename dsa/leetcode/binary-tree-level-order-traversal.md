# Binary tree level order traversal

Given: The root of a binary tree, return the level order traversal of its nodes' values,
from left to right, level by level. 

# Approach 1

Perform BFS and store {node, level} in the queue. Add 1 to the level of current node while
storing its children in the queue.

Time complexity: O(n)

```bash
    vector<vector<int>> levelOrder(TreeNode* root) {
        queue<pair<TreeNode*, int>> q;
        vector<vector<int>> ans;
        if(!root) return {};
        q.push({root, 0});
        while(!q.empty()) {
            auto [node, pos] = q.front();
            q.pop();
            if(ans.size() <= pos) ans.resize(pos+1);
            ans[pos].push_back(node->val);
            if(node->left) q.push({node->left, pos+1});
            if(node->right) q.push({node->right, pos+1});
        }
        return ans;
    }
```

# Approach 2

Store nodes in a queue, but process only the nodes of one level at a time.

Time complexity: O(n)

```bash
    vector<vector<int>> levelOrder(TreeNode* root) {
        queue<TreeNode*> q;
        vector<vector<int>> ans;
        if(!root) return {};
        q.push(root);
        while(!q.empty()) {
            int size = q.size();
            vector<int> level;
            for(int i=0; i<size; i++) {
                auto node = q.front();
                q.pop();
                level.push_back(node->val);
                if(node->left) q.push(node->left);
                if(node->right) q.push(node->right);
            }
            ans.push_back(level);
        }
        return ans;
    }
```