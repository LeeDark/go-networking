package main

import (
	"fmt"
	"net"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Fprintf(os.Stderr, "Usage: %s hostname\n", os.Args[0])
		os.Exit(1)
	}
	name := os.Args[1]

	addr, err := net.ResolveIPAddr("ip6", name)
	if err != nil {
		fmt.Println("Resolution error", err.Error())
		os.Exit(2)
	}

	fmt.Println("Resolved address is", addr.String())
	os.Exit(0)
}
