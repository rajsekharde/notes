# ip

Primary tool for managing network interfaces, IP addresses, and routing in modern Linux systems

## Commands

1. ip -br addr
- Shows network summary of host machine
- Best daily tool
- Columns: Interface name, Status, IP Address
- Example output: wlp2s0    UP    10.175.49.21/24

2. ip addr
- Shows all network interfaces and their addresses
- Each block of output shows one interface
- Interface Names:
    - wlp2s0 - WiFi
    - enp1s0 - Ethernet
    - lo - Loopback
- State Flags:
    - UP - interface enabled
    - LOWER_UP - physically connected
    - DOWN/NO-CARRIER - not connected
- IP Address: inet 10.175.49.21/24
    - Actual System IP
    - inet - IPv4 address
    - inet6 - IPv6 address
    - /24 subnet size

My IP when using WiFi:
- wlp2s0: <UP, LOWER_UP, ...>
-  inet 10.175.49.21/24

3. ip route
- Shows Internet path, IP address of router
- Structure: default via 10.175.49.41 dev wlp2s0
    - default: all unknown traffic (internet)
    - via 10.175.49.41 : my router
    - dev wlp2s0 : which interface is used (WiFi)
