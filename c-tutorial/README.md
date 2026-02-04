# C Basics

## Building and running a C file
```bash
# Check GCC compiler version
gcc --version

# Write the code in a file- hello.c
#include <stdio.h>

int main() {
    printf("Hello World");
    return 0;
}

#Compile the file:
gcc hello.c -o hello_program
# Omitting the -o flag creates a default executable 'a.out'

#Run the executable
./hello_program
```

## Building and running a C++ file
```bash
# Check g++ compiler version
g++ --version

# Write the code in a file- hello.cpp
#include <iostream>
using namespace std;

int main() {
    cout << "Hello World!\n";
    return 0;
}

#Compile the file:
g++ hello.cpp -o hello_program_cpp
# Omitting the -o flag creates a default executable 'a.out'

#Run the executable
./hello_program_cpp
```

