package main

import (
	"strings"

	"github.com/containerd/containerd"
	progressbar "github.com/schollz/progressbar/v3"
	"github.com/sirupsen/logrus"
)

func main() {
	bar := progressbar.Default(100)
	values := strings.Split("the answer to life, the universe and everything", ",")
	searchName := "the answer to life"
	// BAD: index could be equal to length
	for i := 0; i <= len(values); i++ {
		bar.Add(1)
		// When i = length, this access will be out of bounds
		if values[i] == searchName {
			logrus.Debugf("Found at position %d", i)
		}
	}
	ok := containerd.CheckRuntime("runc", "crio")
	if !ok {
		panic("containerd runtime not found")
	}
}
