# File Descriptor

It's an ID number that the OS gives to a process when it opens a file or I/O resource.

When a program opens a file, the OS opens the file internally, assigns an integer (file descriptor) to represent the open file. The program uses that number to interact with the file.

Default file descriptor: Every Unix process starts with three default file descriptors:
FD  Name    Purpose
0   stdin   Standard Input
1   stdout  Standard Output
2   stderr  Standard Error

A file descriptor is not only for files, it also represents: network sockets, pipes, devices, FIFOs


## Process file descriptor table

The OS maintains a table of open files for each running process.
Each process has a table like this:
FD  Pointer
0   stdin
1   stdout
2   stderr
3   file.txt
4   socket


