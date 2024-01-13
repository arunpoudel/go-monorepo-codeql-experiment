package main

import (
	"net/http"

	"golang.org/x/net/http2"
)

func main() {
	client := http.DefaultClient
	_ = http2.ConfigureTransport(client.Transport.(*http.Transport))
}
