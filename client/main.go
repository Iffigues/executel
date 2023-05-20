package main

import (
	"os"
	"fmt"
)

func main() {
	fmt.Println(os.Args)
	e := os.Environ()
	for _, val := range e {
		fmt.Println(val)
	}
}
