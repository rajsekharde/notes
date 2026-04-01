### Task Execution by CPU

Fetch -> Decode -> Execute -> Repeat

1. Instruction fetch:
- CPU uses the instruction pointer / program counter. Fetches next instruction from memory.
- First checks caches: L1 > L2 > L3. Loads from memory if not found in cache.

2. Instruction decode:
- Machine instruction are in binary.
- CPU decodes it into:
    - Operation (ADD, LOAD, JUMP ,etc)
    - Operands (registers, memory addresses)

3. Scheduled, reordered

## Example
```bash
for(int i=0; i<1000; i++)
{
    sum += arr[i];
}
```

1. Load arr[i]
- CPU first checks for it in L1 Cache
- If not found in L1, it checks L2 -> L3-> RAM

2. Bring into register
- R1 = arr[i]
- R2 = sum

3. Add in ALU
- R2 = R2 + R1

4. Store Back
- sum = R2

Instead of storing single variables in cache, cache lines are stored. Cache line = 64 bytes. If arr[0] is accessed, the CPU loads arr[0] to arr[15] in cache.