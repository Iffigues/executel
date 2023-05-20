package main

import (
	"executel-server/src/conf"
	"executel-server/src/tcp"
	"fmt"
	"os"
)

func main() {

	c := conf.NewConf("data.yaml")
	fmt.Println(c.GetAddr())
	conn, err := tcp.NewTcp(c)
	if err != nil {
		fmt.Println("Error connecting:", err)
		os.Exit(1)
	}
	defer conn.Close()

	fmt.Println("Connected to server:")
}
