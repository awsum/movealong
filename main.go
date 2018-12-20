package main

import (
	"fmt"
	"os"
)

func findPR() error {
	return nil
}

func main() {
	fmt.Println(os.Args)
	if err := findPR(); err != nil {
		fmt.Println(err)
	}
}
