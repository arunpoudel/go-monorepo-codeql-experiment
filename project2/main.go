package main

import (
	"fmt"
	"strings"

	"github.com/containerd/containerd"
)

func main() {
	values := strings.Split("the answer to life, the universe and everything", ",")
	searchName := "the answer to life"
	// BAD: index could be equal to length
	for i := 0; i <= len(values); i++ {
		// When i = length, this access will be out of bounds
		if values[i] == searchName {
			fmt.Println("Found at position", i)
		}
	}
	ok := containerd.CheckRuntime("runc", "crio")
	if !ok {
		panic("containerd runtime not found")
	}
}
