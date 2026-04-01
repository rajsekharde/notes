# perf

Profiling and performance analysis tool built into the linux kernel.
It tells how programs use CPU, memory and other system resources.

## Uses
- Measure CPU usage and hotspots
- Track hardware events like cache misses, branch mispredictions
- Analyse soft events like context switches, page faults
- Profile applications to find performance bottlenecks

## Profiling a C program

Compile with debug info
```bash
gcc program.c -g -o program
```

High-level summary
```bash
perf stat ./program
```
Shows:
- CPU cycles
- Instructions executed
- Cache misses
- Execution time

Find where time is spent
```bash

# Record execution
perf record ./program

# View report
perf report
```
Tells which functions are slow


See call stacks
```bash
perf record -g ./program
perf report
```
Shows who called the slow function, full call hierarchy

