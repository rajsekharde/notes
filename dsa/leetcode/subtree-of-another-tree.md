# Subtree of another tree

Given: two tree root nodes root and subRoot. Return true if root contains a subtree with same
structure and nodes as subRoot, otherwise false.

# Approach 1 : BFS + Recursive similarity check

Visit each node in order using a queue, and if value is same as subRoot, call a recursive function
to check if the subtrees node and subRoot are equal.

Time complexity:
n = nodes in root, m = nodes in subRoot
BFS in root: O(n)
Similarity check: O(m)
Since it needs to be checked for each node, in worst case,
Overall: O(m * n)

```bash
    bool isSame(TreeNode *p, TreeNode *q) {
        if(!p || !q) return p==q;
        if(p->val != q->val) return false;
        bool l = isSame(p->left, q->left);
        bool r = isSame(p->right, q->right);
        return (l && r);
    }
    bool isSubtree(TreeNode* root, TreeNode* subRoot) {
        queue<TreeNode*> q;
        q.push(root);
        while(!q.empty()) {
            TreeNode* node = q.front();
            q.pop();
            if(node->val == subRoot->val && isSame(node, subRoot))
                return true;
            if(node->left) q.push(node->left);
            if(node->right) q.push(node->right);
        }
        return false;
    }
```