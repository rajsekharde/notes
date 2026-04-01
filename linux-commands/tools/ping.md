# ping

The ping command is used to test the reachability of a host on an IP network and to measure the round-trip time for messages sent to that destination. It sends small packets and waits for a reply.

## Basic Usage

ping google.com

OR

ping 8.8.8.8

Stop with: Ctrl + C

Output:
```bash
PING google.com (2404:6800:4007:815::200e) 56 data bytes
64 bytes from maa05s14-in-x0e.1e100.net (2404:6800:4007:815::200e): icmp_seq=1 ttl=117 time=203 ms
64 bytes from maa05s14-in-x0e.1e100.net (2404:6800:4007:815::200e): icmp_seq=2 ttl=117 time=228 ms
64 bytes from maa05s14-in-x0e.1e100.net (2404:6800:4007:815::200e): icmp_seq=3 ttl=117 time=249 ms
64 bytes from maa05s14-in-x0e.1e100.net (2404:6800:4007:815::200e): icmp_seq=4 ttl=117 time=272 ms
^C
--- google.com ping statistics ---
5 packets transmitted, 4 received, 20% packet loss, time 4005ms
rtt min/avg/max/mdev = 203.273/237.918/271.651/25.327 ms
```
- 64 bytes : Response  size
- icmp_seq=1 : Packet number
- time=203 ms : Latency

Limit number of pings:
- ping -c 3 google.com

## Diagnosis

Ping router: ping 10.175.49.41
- tests WiFi connection, local network
- No reply means WiFi/Network issue

Ping internet (No DNS): ping 8.8.8.8
- Tests internet connectivity
- No reply means no internet

Ping domain: ping google.com
- Tests DNS + Internet
- No reply but 8.8.8.8 works means DNS problem