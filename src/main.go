package main

import (
	"log"
	"os"

	"github.com/tienducitt/go-restful/src/app/migrate"
	"github.com/tienducitt/go-restful/src/app/server"
	"github.com/urfave/cli"
)

func main() {
	clientApp := cli.NewApp()
	clientApp.Name = "My App"
	clientApp.Version = "0.0.1"
	clientApp.Action = server.StartHttpServer
	clientApp.Commands = []cli.Command{
		{
			Name:        "start",
			Description: "Start HTTP Server",
			Action:      server.StartHttpServer,
		},
		{
			Name:        "migrate",
			Description: "migrate db",
			Action:      migrate.Migrate,
		},
	}

	if err := clientApp.Run(os.Args); err != nil {
		log.Println(err.Error())
	}
}
