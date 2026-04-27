# Invert binary tree

Given: root of a binary tree. Invert the tree and return its root

```bash
Example:

input:
      4
  2       7
1   3   6   9

[4, 2, 7, 1, 3, 6, 9]

output:
      4
  7       2
9   6   3   1

[4, 7, 2, 9, 6, 3, 1]
```

# Approach 1

Using recursion. Visit each node and swap its left and right nodes. Base case: null pointer.

Time Complexity: O(n)

```bash
TreeNode* invertTree(TreeNode* root) {
        if(root == nullptr) return root;
        TreeNode *temp = root->right;
        root->right = root->left;
        root->left = temp;
        invertTree(root->right);
        invertTree(root->left);
        return root;
    }
```

# Approach 2

Iteration using stack/queue. Swap the left and right nodes of each node and push the childred to
a stack or a queue.

Time Complexity: O(n)

```bash
TreeNode* invertTree(TreeNode* root) {
        if(root == nullptr) return root;
        queue<TreeNode*> q;
        q.push(root);
        while(!q.empty()) {
            TreeNode *node = q.front();
            q.pop();
            TreeNode *temp = node->right;
            node->right = node->left;
            node->left = temp;
            if(node->right != nullptr) q.push(node->right);
            if(node->left != nullptr) q.push(node->left);
        }
        return root;
    }
```

For using stack, just replace the queue<TreeNode*> with a stack<TreeNode*>.