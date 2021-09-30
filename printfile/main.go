package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	f := os.Args[1]
	res, err := os.Open(f)
	if err != nil {
		fmt.Println("Error reading file", f)
		os.Exit(1)
	}
	io.Copy(os.Stdout, res)
}
