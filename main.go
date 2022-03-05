package main

import (
	"flag"

	"nyiyui.ca/qr/server"
)

func main() {
	var host string
	flag.StringVar(&host, "host", ":8080", "host to bind to")
	flag.Parse()

	server.New().Run(host)
}
