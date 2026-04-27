# Maximum depth of binary tree

Given: Root of a binary tree. Return its maximum depth. A binary tree's maximum depth is the number of
nodes along the longest path from the root node down to the farthest leaf node.

# Approach 1

Using recursion. return 1 + max (depth of right subtree, depth of left subtree). if root = null, return 0.

Time complexity: O(n)

```bash
    int maxDepth(TreeNode* root) {
        if(root == nullptr) return 0;
        int maxD = max(maxDepth(root->right), maxDepth(root->left));
        return maxD + 1;
    }
```

# Approach 2

Iteration using stack/queue.

Time complexity: O(n)

```bash
    int maxDepth(TreeNode* root) {
        if(root == nullptr) return 0;
        int maxD = 1;
        stack<pair<TreeNode*, int>> s;
        s.push({root, 1});
        while(!s.empty()) {
            auto[node, d] = s.top();
            s.pop();
            maxD = max(maxD, d);
            if(node->right != nullptr) {
                s.push({node->right, d+1});
            }
            if(node->left != nullptr) {
                s.push({node->left, d+1});
            }
        }
        return maxD;
    }
```