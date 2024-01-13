package main

import (
	"fmt"
	"os"
)

func main() {
	ptr, err := os.Open("test.go")
	// BAD: ptr is dereferenced before either it or `err` has been checked.
	fmt.Printf("Opened %v\n", *ptr)
	if err != nil {
		fmt.Printf("Bad input: %s\n", err)
	}
}
