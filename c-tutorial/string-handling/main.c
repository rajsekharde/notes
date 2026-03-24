#include <stdio.h>
#include <string.h>

int main()
{
    // strings in C: null-terminated character arrays
    char str1[] = "Hello";
    // Stored in memory as: 'H' 'e' 'l' 'l' 'o' '\0'
    // '\0' - null character. marks the end of the string

    // ways to declare strings

    // array initialization
    char str2[] = "Hello";

    // explicit characters
    char str3[] = {'H', 'e', 'l', 'l', 'o', '\0'};

    // pointer to string literal
    char *str4 = "Hello";

    // str[]    ->  Modifiable
    // str*     ->  Usually poits to read-only memory

    str2[0] = 'A'; // Allowed
    // str4[0] = 'A'; // Program may crash

    str3[4] = '\0'; // Terminates the string at position 4

    // printing
    printf("%s\n", str2);

    // input
    char str[10];
    fgets(str, sizeof(str), stdin); // reads first 9 characters and stores in str

    printf("string: %s\n", str);

    return 0;
}