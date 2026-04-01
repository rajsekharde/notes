# Modern CPU Architecture

## Basic CPU Components
- ALU (Arithmetic & Logic Unit): math & logic ops
- Registers: Fastes storage inside CPU
- Control Unit: Orchestrates execution
- Cache (L1, L2, L3): Small, fast memory
- Execution Units: Multiple ALUs, FPUs, vector units

## Memory Hierarchy
```bash
Registers (Fastest) - 1 cycle
    |
L1 Cache - ~1 cycle
    |
L2 Cache - ~10 cycles
    |
L3 Cache - ~40 cycles
    |
   RAM - ~200 cycles
    |
Disk (Slowest)
```
Analogy: CPU - Workspace
- Registers: Hands (actively holding things)
- L1 Cache: Desk (very close, very fast)
- L2/L3 Cache: Shelves nearby
- RAM: Storage room far away

CPU can only operate on data in registers. Everything else is just 'places to fetch from'

## Caching Hierarchy

### L1 Cache
- Located inside each core
- Speed: 1 cycle
- Size: 64KB - 128KB

### L2 Cache
- Located inside each core
- Speed: ~10 cycles
- Size: 1MB - 2MB

### L3 Cache
- Shared by all cores
- Speed: ~40 cycles
- Size: 16MB - 128MB
