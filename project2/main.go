package main

import (
	"os"
)

func main() {
	file, _ := os.Open("test.txt")
	file.Close()
}
