package main

import (
	"github.com/containerd/containerd"
)

func main() {
	ok := containerd.CheckRuntime("runc", "crio")
	if !ok {
		panic("containerd runtime not found")
	}
}
