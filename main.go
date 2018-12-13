package main

import (
	"fmt"

	"gitlab.com/tokend/keypair"
)

func findPR() error {
	keypair.Random()
	return nil
}

func main() {
	if err := findPR(); err != nil {
		fmt.Println(err)
	}
}
