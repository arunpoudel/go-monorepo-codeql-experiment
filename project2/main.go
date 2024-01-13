package main

import (
	"os"
)

func main() {
	file, err := os.Open("test.txt")
	file.Close()
	if err != nil {
		panic(err)
	}
}
