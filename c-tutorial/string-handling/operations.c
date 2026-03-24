#include <stdio.h>
#include <string.h>

int main()
{
    char str[] = "Hello";

    // length of string
    int len = strlen(str);
    printf("Length of \"%s\" = %d\n", str, len);

    // copy one string to another
    char str2[10];
    strcpy(str2, str); // (DEST, SOURCE)
    printf("Copied %s\n", str2);

    // copy at most n chars from source to dest
    char str3[10];
    strncpy(str3, str, 2);
    printf("Copied %s\n", str3);

    // find the first occurrence of a char in a string
    char *pf = strchr(str, 'l');
    printf("Pos = %p\n", pf);

    // find the last occurrence of a char in a string
    char *pl = strrchr(str, 'l');
    printf("Pos = %p\n", pl);

    // find the first occurrence of a string in a string
    char *spf = strstr(str, "He");
    printf("Pos = %p\n", spf);

    return 0;
}