#include <bits/stdc++.h>
using namespace std;

/*
LIFO container, duplicates allowed, order maintained
*/

int main() {
    // Creation
    stack<int> st;

    // Insertion at end
    st.push(1);
    st.push(2);

    // Accesss last element
    cout << st.top();

    // Delete last element
    st.pop();
    return 0;
}