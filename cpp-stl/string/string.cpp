#include <bits/stdc++.h>

string s = "hello";

// length of a string
int l = s.length();

// converting to lower case
string s2 = "";
for(char ch: s)
{
	s2 += tolower(ch);
}

// tolower() -> returns lower case of a character
// toupper() -> returns upper case of a character
