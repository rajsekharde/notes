#include <iostream>
#include <queue>

using namespace std;

/*
Finding top k largest elements in an array

Approach 1:
    Create a min-heap
    Push first k elements of array into the min-heap
    Iterate through remaining elemets
        If a number is greater than heap.top(),
        Pop the heap and add the number
    After iteration, the heap will have the top k largest numbers

    Time complexity:
        Pushing first k elements: O(klog(k))
        Processing remaining elements: O(nlog(k))
        Total: O(nlog(k))
*/