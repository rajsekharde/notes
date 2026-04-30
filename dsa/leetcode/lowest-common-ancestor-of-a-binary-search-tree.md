# Lowest common ancestor of a binary search tree

Given: The root node of a binary search tree. Find the lowest common ancestor node of two
given nodes in the BST.

# Approach 1

A node in a BST will be the LCA of two nodes if its value is between the values of p, q.
(p < val < q) or (q < val < p). If both are greater than val, the LCA will be to the right
of current node, else if both are smaller than val, the LCA will be to its left. The other
subtree where p,q cannot exist will be skipped completely.

Time Complexity:
Search space gets reduced by half in each iteration, since either left/right subtree is considered
for next iterations. So time is equal to finding height of tree O(H)

Overall: O(H): O(logn) best case, O(n) worst case for skewed tree

```bash
    TreeNode* lowestCommonAncestor(TreeNode* root, TreeNode* p, TreeNode* q) {
        while(root != nullptr) {
            if(p->val > root->val && q->val > root->val) {
                root = root->right;
            } else if(p->val < root->val && q->val < root->val) {
                root = root->left;
            } else {
                break;
            }
        }
        return root;
    }
```