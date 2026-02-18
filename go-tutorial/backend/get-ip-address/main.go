package main

import (
	"fmt"
	"net"
	"os"
)

func main() {
	host, _ := os.Hostname()
	fmt.Println("Hostname:", host)

	addrs, _ := net.InterfaceAddrs()
	for _, addr := range addrs {
		if ipnet, ok := addr.(*net.IPNet); ok && !ipnet.IP.IsLoopback() && ipnet.IP.To4() != nil {
			fmt.Println("LAN IP:", ipnet.IP.String())
		}
	}
}
