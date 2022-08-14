package main

import (
	"fmt"
	"net"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) != 3 {
		fmt.Printf("Usage: %s <host> <port>\n", os.Args[0])
		os.Exit(1)
	} else {
		var host string = os.Args[1]
		var port, _ = strconv.Atoi(os.Args[2])
		if port < 1 || port > 65535 {
			fmt.Printf("Invalid port number: %d\n", port)
			os.Exit(1)
		}
		if net.ParseIP(host) == nil {
			fmt.Printf("Invalid host: %s\n", host)
			os.Exit(1)
		}
		StartServer(port, host)
	}
}
