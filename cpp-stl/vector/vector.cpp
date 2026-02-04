#include <iostream>
#include <vector>
#include <algorithm>
using namespace std;

/*
Time complexity:

Insertion in the end by push_back()- O(1)
Insertion in the middle by insert()- O(n)

Deletion from end by pop_back()- O(1)
Deletion from middle by erase()- O(n)

Search by index- O(1)
Search by value- O(n)
*/

void display(vector<int> &v) {
    for (int i: v)
        printf("%d\n", i);
}

int main() {
    // Declare an empty vector
    vector<int> v1;

    // Declare a vector with given size and fill with value
    vector<int> v2(5, 10);

    // Declare a vector and initialize it using a list
    vector<int> v3 = {1, 5, 3, 2, 8, 10, 7, 6, 4, 9};
    vector<int> v4 = v3;

    // Add values to a vector
    for (int i=1; i<=10; i++)
        v1.push_back(i);
    
    // Delete last value
    v1.pop_back();

    // Access a value using it's index
    printf("%d\n", v1.at(5));
    
    // Iterate through a vector
    for (int i: v2)
        printf("%d\n", i);
    
    // Sort a vector in ascending order
    sort(v3.begin(), v3.end());
    display(v3);

    // Sort a vector in descending order
    sort(v4.begin(), v4.end());
    display(v4);

    // Reverse a vector
    reverse(v4.begin(), v4.end());
    display(v4);

    // Check if a vector is empty
    cout << v1.empty() << "\n"; // prints true(1) if empty

    return 0;
}