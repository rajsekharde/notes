# Memory management in C

## Memory layout of a process
- Heap Memory: Manually managed using malloc & free
- Stack Memory: Automatically allocated & deallocated by CPU
- Global/Static variables:
    - Un-initialised
    - Initialised
- Text/Code: Stores the code and instructions

## Stack Memory

Managed by CPU and follows a LIFO pattern. When a function is called, all of its local variables are pushed to the stack. When the function returns, that memory is popped and immediately available for reuse.

The stack memory for a variable exists as long as it's in scope. It is faster but it's size is limited (usually ~8MB)

## Heap Memory

Managed manually by the user. Allocated using malloc and de-allocated using free. Any function with the pointer to a heap variable can access its memory.

Heap memory is released only when free() is called or the entire process terminates. It is slower as it requires searching for free memory, but has a large size (limited by RAM).

## Global Variables

Shared by all members of a program. Not owned by anyone. Allocated when program starts running and released when program ends.