# Same tree

Given: Root nodes of two trees, return true if both trees are same- structurally equal and nodes
have same values, else return false

# Approach 1 : BFS/DFS

Keep storing pairs of nodes at same position in a stack/queue and keep popping them together and
check if they are equal. handle edge cases

Time complexity: O(n)

```bash
bool isSameTree(TreeNode* p, TreeNode* q) {
        queue<TreeNode*> qu;
        if(p == nullptr || q == nullptr) return p == q;
        qu.push(p);
        qu.push(q);
        while(!qu.empty()) {
            TreeNode *a = qu.front();
            qu.pop();
            TreeNode *b = qu.front();
            qu.pop();

            if(a->val != b->val) return false;
            
            if(a->left != nullptr && b->left != nullptr) {
                qu.push(a->left);
                qu.push(b->left);
            } else if(a->left != b->left) return false;

            if(a->right != nullptr && b->right != nullptr) {
                qu.push(a->right);
                qu.push(b->right);
            } else if(a->right != b->right) return false;
        }
        return true;
    }
```

# Approach 2 : Recursion

Recursively check if each sub-tree is same or not, and compare the results of right & left sub-trees.
Base case: either p or q is null. 

Time complexity: O(n)

```bash
    bool isSameTree(TreeNode* p, TreeNode* q) {
        if(p == nullptr || q == nullptr) return p == q;
        if(p->val != q->val) return false;
        bool l = isSameTree(p->left, q->left);
        bool r = isSameTree(p->right, q->right);
        return (l && r);
    }
```