#include <bits/stdc++.h>
using namespace std;

class Solution
{
public:
    bool isPalindrome(string s)
    {
        string sc = "";
        for (char ch : s)
        {
            ch = tolower(ch);
            if ((ch >= 'a' && ch <= 'z') || (ch >= '0' && ch <= '9'))
                sc += ch;
        }
        if (sc == "")
            return true;
        int l = 0, r = sc.length()-1;
        while (l < r)
        {
            if (sc[l] != sc[r])
                return false;
            l++;
            r--;
        }
        return true;
    }
};

int main()
{
    Solution s;
    cout << s.isPalindrome("A man, a plan, a canal: Panama");
    cout << "\n";
    return 0;
}