#include <bits/stdc++.h>
using namespace std;

/*
Binary Search Tree

Left Child < Parent < Right Child
*/

struct Node {
    int value;
    struct Node* left;
    struct Node* right;
};

Node* addNode(Node* root, int value) {
    Node* newNode = new Node{value, nullptr, nullptr};

    if(!root)
        return newNode;
    
    Node* temp = root;
    while(true) {
        if(value < temp->value) {
            if(!temp->left) {
                temp->left = newNode;
                break;
            }
            temp = temp->left;
        } else {
            if(!temp->right) {
                temp->right = newNode;
                break;
            }
            temp = temp->right;
        }
    }
    
    return root;
}

int main() {
    Node* root = nullptr;
    root = addNode(root, 10);
    for(int i=20; i<=100; i+=10)
        root = addNode(root, i);
    
    cout << root->left->value;
    
    return 0;
}