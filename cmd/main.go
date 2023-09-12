package main

import (
	"flag"
	"libnet/internal/app"

	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
)

var (
	configPath = flag.String("conf", "./app/configs/app.json", "path to config file")
)

func main() {
	flag.Parse()

	server, err := app.Init(*configPath)
	if err != nil {
		log.Fatal(errors.WithStack(err))
	}

	server.Start()
}
