package main

import (
	"fmt"
	"net"
	"os"
)

func main() {
	serverAddr := "localhost:8080" // Change it to the desired server address

	conn, err := net.Dial("tcp", serverAddr)
	if err != nil {
		fmt.Println("Error connecting:", err)
		os.Exit(1)
	}
	defer conn.Close()

	fmt.Println("Connected to server:", serverAddr)
}
