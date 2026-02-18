package main

import (
	"fmt"
	"log"
	"net"
	// "os"

	"github.com/skip2/go-qrcode"
)

func main() {
	// Get LAN IP
	var lanIP string
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		log.Fatal(err)
	}
	for _, addr := range addrs {
		if ipnet, ok := addr.(*net.IPNet); ok && !ipnet.IP.IsLoopback() && ipnet.IP.To4() != nil {
			lanIP = ipnet.IP.String()
			break
		}
	}

	if lanIP == "" {
		log.Fatal("No LAN IP found")
	}

	// Generate URL
	url := fmt.Sprintf("http://%s:8000", lanIP)
	fmt.Println("Open this URL on your phone:", url)

	// Generate QR code PNG
	err = qrcode.WriteFile(url, qrcode.Medium, 256, "server_qr.png")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("QR code saved as server_qr.png")
}
