# Process creation in linux

## System Calls
- fork(): Clones the parent to create a child.
- execve(): Loads the new program into the clone.
- wait(): Parent waits for the child to finish to prevent 'zombies'.
- exit(): The process tells the kernel it is finished.

## Example: running ls command in shell

1. Parent preparation (fork):
When ls is typed in the shell, the shell is currently running a process.
- To run ls, the shell calls the fork() system call. 
- The kernel creates a near-exact duplicate of the parent process.
- A new task_struct is allocated in the kernel memory. It copies the parent's
file descriptors, user IDs, and resource limits.
- The child gets a copy of the parent's address space. Linux uses Copy-On-Write for this.
- Copy-On-Write(COW): Instead of copying the RAM immediately, both the parent and child
processes point to the same physical memory pages. The kernel creates a unique copy for
the process only if one of them tries to write data.

2. Identity creation (execve):
Currently there are two identical processes or shells.
- The child process then calls execve().
- The kernel loads the binary executables (ls program) into the child's memory space.
- The child's old memory (cloned shell code) is discarded. Its stack, heap and data
segments are initialised fresh for the new program.
- PID remains the same, and the copied file descriptors are preserved.

3. Kernel's 'Birth' routine (do_fork):
The function do_fork() deep inside the kernel performs the heavy-lifting.
- copy_process(): This checks if the user has permissions to create a new process (resource limits).
- The kernel looks at its PID bitmap and assigns the next available unique integer to the child.
- The child is set to TASK_RUNNING, and placed into the Run Queue. Its not running yet.

4. Scheduling:
- The child process is now a runnable entity. It sits in the Red-Black Tree of the
Completely Fair Scheduler (CFS).
- Eventually the scheduler decides its the child's turn. It performs a context switch:
	- The CPU registers are loaded with the child's instruction pointer.
	- The CPU jumps to the entry point of the new code.
	- The process is now alive and executing instructions.

5. Death and 'Zombie Process' phase:
- When the process finishes, it calls exit(). The kernel releases most of its memory and resources.
- The process becomes a Zombie. It keeps its entry in the process table so that the parent can
see how it died (the exit code).
- The parent process calls wait(). Once the parent acknowledges the death, the kernel finally
deletes the task_struct and the PID is freed for reuse.

