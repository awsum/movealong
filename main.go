package main

import "fmt"

func fn(i int) error {
	if i > 1 {
		return fmt.Errorf("%d", i)
	}
	return nil
}

func main() {}
