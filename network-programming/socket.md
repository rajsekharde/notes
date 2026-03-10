# Socket

A way to speak to other programs using standars Unix file descriptors. It's a communication endpoint used by programs to send and receive data over a network, using protocols like TCP, UDP.
In Unix/linux, sockets are treated like files. When a program creates a socket, the OS returns a file descriptor.
A network socket is uniquely identified by: IP Address + Port Number

## Two types of sockets- Stream sockets & Datagram sockets


