# ss

The ss command (socket statistics) is used for displaying detailed information about network sockets. Shows which connections and ports are open on my machine

## Commands

1. ss -tuln

Output:
```bash
Netid      State       Recv-Q      Send-Q            Local Address:Port             Peer Address:Port      Process      
udp        UNCONN      0           0                       0.0.0.0:5353                  0.0.0.0:*                      
udp        UNCONN      0           0                 192.168.122.1:53                    0.0.0.0:*                      
udp        UNCONN      0           0                             *:53054                       *:*                      
udp        UNCONN      0           0                          [::]:5353                     [::]:*                      
tcp        LISTEN      0           4096                 127.0.0.54:53                    0.0.0.0:*                      
tcp        LISTEN      0           10                      0.0.0.0:8080                  0.0.0.0:*                      
tcp        LISTEN      0           4096              127.0.0.53%lo:53                    0.0.0.0:*                      
tcp        LISTEN      0           32                192.168.122.1:53                    0.0.0.0:*                      
tcp        LISTEN      0           4096                  127.0.0.1:631                   0.0.0.0:*                      
tcp        LISTEN      0           4096                      [::1]:631                      [::]:*             
```

Columns:
- Netid : tcp or udp
- State:
    - LISTEN - waiting for connections (server)
    - ESTAB - active connection
    - UNCONN - ready to receive packets (UDP)
- Local Address:Port : 0.0.0.0:22
    - Port 22 is open (SSH)
    - 0.0.0.0 - Accessible from anywhere
- Peer Address:Port
    - Who is connected (for active connections)
    - * : Anyone

2. ss -tulpn
- -p flag shows the processes using these sockets

3. ss -tuln | grep 8080 or ss -tulpn | grep 8080
- Filter by port