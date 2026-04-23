#include <bits/stdc++.h>
using namespace std;

int main()
{
    vector<int> v;
    for(int i=0; i<10e8; i++) v.push_back(i);
    int t=10;
    int l=0, r=v.size()-1;
    while(l<=r)
    {
        int m = l + (r-l)/2;
        cout << v[m] << "\n";
        if(t == v[m]) {cout << "Found\n"; break;}
        if(v[m] > t) r = m - 1;
        else l = m + 1;
    }
    cout << "\n";
    return 0;
}