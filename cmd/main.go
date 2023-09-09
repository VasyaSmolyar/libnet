package main

import (
	"flag"
	"libnet/internal/app"
)

var (
	configPath = flag.String("conf", "./configs/app.json", "path to config file")
)

func main() {
	flag.Parse()

	server, err := app.Init(*configPath)
	if err == nil {
		server.Start()
	}
}
