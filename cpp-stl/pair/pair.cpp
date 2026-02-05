#include <bits/stdc++.h>
using namespace std;

/*
Pair: Stores two values of any type as a single unit
*/

int main() {
    pair<int, string> p1 = {1, "A"};
    cout << p1.first << " " << p1.second;

    pair<char, int> p2;
    p2.first = 'B';
    p2.second = 2;

    p2.second = 5;

    return 0;
}