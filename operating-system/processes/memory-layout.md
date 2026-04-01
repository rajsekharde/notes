# Memory layout of processes

## Overview
```bash
-------------------     High Address
    Stack           Grows Downward
-------------------
    Memory-Mapped   (Shared libraries, mmap)
        Region
-------------------
    Heap            Grows Upward
-------------------
    BSS-Segment     (Uninitialised Global/Static variables)
-------------------
    Data-Segment    (Initialised Global/Static variables)
-------------------
    Text-Segment    (Program Code)
-------------------     Low Address
```

## Text Segment
- Contains the compiled machine code of the program
- Usually Read-Only to prevent accidental modification
- Can be shared among multiple processes running the same program
- Instructions for main live here

## Data Segment
- Stores global and static variables initialised with values
- Exists for the lifetime of the program

## BSS Segment (Block Started by Symbol)
- Contains unintialised global and static variables
- Variables initialised to zero by default
- Does not occupy space in the executable (Just size info)

## Heap
- Used for dynamic memory allocation
- Managed at runtime by the program
- Grows upward (Toward higher addresses)
- Slower than stack allocation
- Can fragment over time

## Memory-Mapped Region
- Stores:
    - Shared libraries
    - Memory-Mapped files (mmap)
    - Anonymous mappings
- Used for:
    - loading libraries like libc
    - Inter-Process Communication

## Stack
- Stores:
    - Function call frames
    - Local variables
    - Return addresses
- Grows downward
- Very fast allocation/deallocation (FIFO)
- Limited size: Stack overflow possible

## Kernel Space (Not Directly Accessible)
- Each process also interacts with Kernel memory:
    - System Calls
    - Process control structures
- User processes cannot directly access it

## Permissions
Stack       Read + Write
Heap        Read + Write
Data        Read + Write
Text        Read + Executable